package reactive

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/nsf/jsondiff"
	"go.etcd.io/etcd/pkg/adt"

	"github.com/rancher/opni/pkg/plugins/driverutil"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/util"
	"github.com/rancher/opni/pkg/util/fieldmask"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type DiffMode int

const (
	DiffStat DiffMode = iota
	DiffFull
)

type ControllerOptions struct {
	logger   *slog.Logger
	diffMode DiffMode
}

type ControllerOption func(*ControllerOptions)

func (o *ControllerOptions) apply(opts ...ControllerOption) {
	for _, op := range opts {
		op(o)
	}
}

func WithLogger(eventLogger *slog.Logger) ControllerOption {
	return func(o *ControllerOptions) {
		o.logger = eventLogger
	}
}

func WithDiffMode(mode DiffMode) ControllerOption {
	return func(o *ControllerOptions) {
		o.diffMode = mode
	}
}

type Controller[T driverutil.ConfigType[T]] struct {
	ControllerOptions
	tracker *driverutil.DefaultingConfigTracker[T]

	runOnce    atomic.Bool
	runContext context.Context

	reactiveMessagesMu sync.Mutex
	reactiveMessages   *pathTrie[*reactiveValue]

	currentRevMu sync.Mutex
	currentRev   int64
}

func NewController[T driverutil.ConfigType[T]](tracker *driverutil.DefaultingConfigTracker[T], opts ...ControllerOption) *Controller[T] {
	options := ControllerOptions{}
	options.apply(opts...)

	return &Controller[T]{
		ControllerOptions: options,
		tracker:           tracker,
		reactiveMessages:  newPathTrie(util.NewMessage[T]().ProtoReflect().Descriptor(), newReactiveValue),
	}
}

func (s *Controller[T]) Start(ctx context.Context) error {
	if !s.runOnce.CompareAndSwap(false, true) {
		panic("bug: Run called twice")
	}
	s.runContext = ctx

	var rev int64
	_, err := s.tracker.ActiveStore().Get(ctx, storage.WithRevisionOut(&rev))
	if err != nil {
		if !storage.IsNotFound(err) {
			return err
		}
	}
	w, err := s.tracker.ActiveStore().Watch(ctx, storage.WithRevision(rev))
	if err != nil {
		return err
	}
	if rev != 0 {
		// The first event must be handled before this function returns, if
		// there is an existing configuration. Otherwise, a logic race will occur
		// between the goroutine below and calls to Reactive() after this
		// function returns. New reactive values have late-join initialization
		// logic; if the first event is not consumed, the late-join logic and
		// the goroutine below (when it is scheduled) would cause duplicate
		// updates to be sent to newly created reactive values.
		firstEvent, ok := <-w
		if !ok {
			return fmt.Errorf("watch channel closed unexpectedly")
		}

		// At this point there will most likely not be any reactive values, but
		// this function sets s.currentRev and also logs the first event.
		s.handleWatchEvent(firstEvent)
	}
	go func() {
		for {
			cfg, ok := <-w
			if !ok {
				return
			}
			s.handleWatchEvent(cfg)
		}
	}()
	return nil
}

func (s *Controller[T]) handleWatchEvent(cfg storage.WatchEvent[storage.KeyRevision[T]]) {
	s.reactiveMessagesMu.Lock()
	defer s.reactiveMessagesMu.Unlock()

	group := make(chan struct{})
	defer func() {
		close(group)
	}()

	switch cfg.EventType {
	case storage.WatchEventDelete:
		if s.logger != nil {
			s.logger.With(
				"key", cfg.Previous.Key(),
				"prevRevision", cfg.Previous.Revision(),
			).Info("configuration deleted")
		}
		all := s.reactiveMessages.Stab(adt.NewStringAffineInterval("\x00", ""))
		for _, iv := range all {
			rm := iv.Val.(*reactiveValue)
			rm.Update(cfg.Previous.Revision(), protoreflect.Value{}, group)
		}
	case storage.WatchEventPut:
		s.currentRevMu.Lock()
		s.currentRev = cfg.Current.Revision()
		s.currentRevMu.Unlock()

		// efficiently compute a list of paths (or prefixes) that have changed
		var prevValue T
		if cfg.Previous != nil {
			prevValue = cfg.Previous.Value()
		}
		diffMask := fieldmask.Diff(prevValue, cfg.Current.Value())

		if s.logger != nil {
			opts := jsondiff.DefaultConsoleOptions()
			opts.SkipMatches = true
			diff, _ := driverutil.RenderJsonDiff(prevValue, cfg.Current.Value(), opts)
			stat := driverutil.DiffStat(diff, opts)
			switch s.diffMode {
			case DiffStat:
				s.logger.Info("configuration updated", "revision", cfg.Current.Revision(), "diff", stat)
			case DiffFull:
				s.logger.Info("configuration updated", "revision", cfg.Current.Revision(), "diff", stat)
				s.logger.Info("⤷ diff:\n" + diff)
			}
		}
		for _, path := range diffMask.Paths {
			ivs := s.reactiveMessages.Stab(adt.NewStringAffineInterval(path+".", prefixRangeEnd(path+".")))
			iv1 := s.reactiveMessages.Find(adt.NewStringAffinePoint(path))
			for _, iv := range append(ivs, iv1) {
				// get the new value of the current message
				key := iv.Ivl.Begin.(adt.StringAffineComparable)
				value := protoreflect.ValueOf(cfg.Current.Value().ProtoReflect())
				for _, part := range strings.Split(string(key), ".") {
					value = value.Message().Get(value.Message().Descriptor().Fields().ByName(protoreflect.Name(part)))
				}
				// update the reactive messages
				rm := iv.Val.(*reactiveValue)
				rm.Update(cfg.Current.Revision(), value, group)
			}
		}
	}
}

func (s *Controller[T]) Reactive(path protopath.Path) Value {
	if len(path) < 2 || path[0].Kind() != protopath.RootStep {
		panic(fmt.Sprintf("invalid reactive message path: %s", path))
	}
	return s.reactiveMessages.Find(path).value
}

type pathTrie[V any] struct {
	root *pathTrieNode[V]
}

type pathTrieNode[V any] struct {
	protopath.Step
	parent *pathTrieNode[V]
	nodes  map[protoreflect.Name]*pathTrieNode[V]
	value  V
}

func newPathTrie[V any](desc protoreflect.MessageDescriptor, newV func() V) *pathTrie[V] {
	t := &pathTrie[V]{
		root: &pathTrieNode[V]{
			Step: protopath.Root(desc),
		},
	}
	buildNode[V](t.root, desc, newV)
	return t
}

func buildNode[V any](node *pathTrieNode[V], desc protoreflect.MessageDescriptor, newV func() V) {
	node.nodes = make(map[protoreflect.Name]*pathTrieNode[V])
	node.value = newV()
	for i := 0; i < desc.Fields().Len(); i++ {
		field := desc.Fields().Get(i)
		node := &pathTrieNode[V]{
			parent: node,
			Step:   protopath.FieldAccess(field),
		}
		if field.Kind() == protoreflect.MessageKind && !field.IsMap() && !field.IsList() {
			buildNode(node, field.Message(), newV)
		}
		node.nodes[field.Name()] = node
	}
}

func (t *pathTrie[V]) Find(path protopath.Path) *pathTrieNode[V] {
	node := t.root
	for _, step := range path[1:] {
		node = node.nodes[step.FieldDescriptor().Name()]
		if node == nil {
			return nil
		}
	}
	return node
}

// Walk performs a depth-first post-order traversal of the trie, calling fn
// for each node. The root node is visited last.
func (t *pathTrie[V]) Range(fn func(*pathTrieNode[V])) {
	var walk func(*pathTrieNode[V])
	walk = func(node *pathTrieNode[V]) {
		for _, child := range node.nodes {
			walk(child)
		}
		fn(node)
	}
	walk(t.root)
}

// Given a field mask and a root value, visitFn is called for each unique path
// step in the trie that intersects with a path in the field mask.
//
// Example:
// Given the following message definition and object:
//
//		message Root {
//		  string r1 = 1;
//		  message A {
//		    string a1 = 1;
//		    message B {
//		      string b1 = 1;
//		    }
//		    B b = 2;
//		    message C {
//		      string c1 = 1;
//		    }
//		    C c = 3;
//		  }
//		  A a = 2;
//		}
//
//	 root := &Root{
//	   R1: "r1value",
//	   A: &Root_A{
//	     A1: "a1value",
//	     B: &Root_A_B{
//	       B1: "b1value",
//	     },
//	     C: &Root_A_C{
//	       C1: "c1value",
//	     },
//	   },
//	 }
//
// The list below shows the paths that would be visited for each field mask:
//
//  1. mask: ["a.b.b1"]
//     path        | value
//     ------------|------
//     root.A.B.B1 | "b1value"
//     root.A.B    | &Root_A_B{...}
//     root.A      | &Root_A{...}
//     root        | &Root{...}
//
//  2. mask: ["a.b", "a.c.c1"]
//     path        | value
//     ------------|------
//     root.A.B    | &Root_A_B{...}
//     root.A.C.C1 | "c1value"
//     root.A.C    | &Root_A_C{...}
//     root.A      | &Root_A{...}
//     root        | &Root{...}
//
//  3. mask: ["a.b", "a.c"]
//     path        | value
//     ------------|------
//     root.A.B    | &Root_A_B{...}
//     root.A.C    | &Root_A_C{...}
//     root.A      | &Root_A{...}
//     root        | &Root{...}
//
//  4. mask: ["a.a1", "a.b.b1", "a.c.c1", "r1"]
//     path        | value
//     ------------|------
//     root.A.B.B1 | "b1value"
//     root.A.B    | &Root_A_B{...}
//     root.A.C.C1 | "c1value"
//     root.A.C    | &Root_A_C{...}
//     root.A.A1   | "a1value"
//     root.A      | &Root_A{...}
//     root.R1     | "r1value"
//     root        | &Root{...}
func (t *pathTrie[V]) Stab(m *fieldmaskpb.FieldMask, rootValue protoreflect.Message, visitFn func(V, protoreflect.Value)) {

}
