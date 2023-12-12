package reactive

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"

	"github.com/nsf/jsondiff"

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

	lg := options.logger
	if lg != nil {
		lg = options.logger.WithGroup("value")
	}

	return &Controller[T]{
		ControllerOptions: options,
		tracker:           tracker,
		reactiveMessages: newPathTrie(util.NewMessage[T]().ProtoReflect().Descriptor(), func() *reactiveValue {
			return newReactiveValue(lg)
		}),
	}
}

func (s *Controller[T]) traceLog(msg string, attrs ...any) {
	if s.logger == nil {
		return
	}
	s.logger.Log(context.Background(), traceLogLevel, msg, attrs...)
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
	s.traceLog("controller starting", "revision", rev)
	go func() {
		for {
			cfg, ok := <-w
			if !ok {
				s.traceLog("controller watch channel closed")
				return
			}
			s.traceLog("controller received watch event")
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

	var currentRev int64
	var currentVal protoreflect.Value
	var diffMask *fieldmaskpb.FieldMask
	switch cfg.EventType {
	case storage.WatchEventDelete:
		if s.logger != nil {
			s.logger.With(
				"key", cfg.Previous.Key(),
				"prevRevision", cfg.Previous.Revision(),
			).Info("configuration deleted")
		}
		currentRev = cfg.Previous.Revision()
		currentVal = protoreflect.Value{}
		diffMask = fieldmask.Diff(cfg.Previous.Value().ProtoReflect(), cfg.Previous.Value().ProtoReflect().Type().Zero())
	case storage.WatchEventPut:
		currentRev = cfg.Current.Revision()
		currentVal = protoreflect.ValueOf(cfg.Current.Value().ProtoReflect())
		s.traceLog("configuration updated", "revision", currentRev)

		// efficiently compute a list of paths (or prefixes) that have changed
		var prevValue T
		if cfg.Previous != nil {
			s.traceLog("previous config existed", "revision", cfg.Previous.Revision())
			prevValue = cfg.Previous.Value()
		}
		diffMask = fieldmask.Diff(prevValue.ProtoReflect(), cfg.Current.Value().ProtoReflect())
		s.traceLog("diff mask", "paths", diffMask.GetPaths())

		if s.logger != nil {
			opts := jsondiff.DefaultConsoleOptions()
			opts.SkipMatches = true
			diff, _ := driverutil.RenderJsonDiff(prevValue, cfg.Current.Value(), opts)
			stat := driverutil.DiffStat(diff, opts)
			switch s.diffMode {
			case DiffStat:
				s.logger.Info("configuration updated", "revision", currentRev, "diff", stat)
			case DiffFull:
				s.logger.Info("configuration updated", "revision", currentRev, "diff", stat)
				s.logger.Info("⤷ diff:\n" + diff)
			}
		}
	}
	s.currentRevMu.Lock()
	s.currentRev = currentRev
	s.traceLog("storing current revision", "revision", currentRev)
	s.currentRevMu.Unlock()

	pathsToNotify := make(map[string]struct{})
	for _, pathStr := range diffMask.GetPaths() {
		pathsToNotify[pathStr] = struct{}{}
	}
	s.traceLog("paths to notify", "paths", pathsToNotify)
	s.reactiveMessages.WalkValues(func(node *pathTrieNode[*reactiveValue], value protoreflect.Value) {
		var shouldNotify bool
		if len(node.Path) == 1 {
			// root node, always notify
			shouldNotify = true
		} else if _, ok := pathsToNotify[node.Path[1:].String()[1:]]; ok {
			shouldNotify = true
		}
		s.traceLog("dispatching update", "currentRev", currentRev, "path", node.Path, "value", value, "shouldNotify", shouldNotify)
		node.value.Update(currentRev, value, group, shouldNotify)
	}, currentVal)
}

func (s *Controller[T]) usePathTrie(fn func(p *pathTrie[*reactiveValue])) {
	s.reactiveMessagesMu.Lock()
	defer s.reactiveMessagesMu.Unlock()
	fn(s.reactiveMessages)
}

func (s *Controller[T]) Reactive(path protopath.Path) Value {
	if len(path) < 2 || path[0].Kind() != protopath.RootStep {
		panic(fmt.Sprintf("invalid reactive message path: %s", path))
	}
	if v := s.reactiveMessages.Find(path); v != nil {
		return v.value
	}
	panic(fmt.Sprintf("bug: reactive message not found: %s", path))
}

type pathTrie[V any] struct {
	root  *pathTrieNode[V]
	index map[string]*pathTrieNode[V]
}

type pathTrieNode[V any] struct {
	protopath.Path
	parent *pathTrieNode[V]
	nodes  map[protoreflect.Name]*pathTrieNode[V]
	value  V
}

func newPathTrie[V any](desc protoreflect.MessageDescriptor, newV func() V) *pathTrie[V] {
	t := &pathTrie[V]{
		root: &pathTrieNode[V]{
			Path: protopath.Path{protopath.Root(desc)},
		},
	}
	buildNode[V](t.root, desc, newV)
	t.index = make(map[string]*pathTrieNode[V])
	t.Walk(func(node *pathTrieNode[V]) {
		if len(node.Path) == 1 {
			return
		}
		t.index[node.Path[1:].String()[1:]] = node
	})
	return t
}

func buildNode[V any](node *pathTrieNode[V], desc protoreflect.MessageDescriptor, newV func() V) {
	node.nodes = make(map[protoreflect.Name]*pathTrieNode[V])
	node.value = newV()
	for i := 0; i < desc.Fields().Len(); i++ {
		field := desc.Fields().Get(i)
		newNode := &pathTrieNode[V]{
			parent: node,
			Path:   append(append(protopath.Path{}, node.Path...), protopath.FieldAccess(field)),
			value:  newV(),
		}
		if field.Kind() == protoreflect.MessageKind && !field.IsMap() && !field.IsList() {
			buildNode(newNode, field.Message(), newV)
		}
		node.nodes[field.Name()] = newNode
	}
}

func (t *pathTrie[V]) Find(path protopath.Path) *pathTrieNode[V] {
	return t.index[path[1:].String()[1:]]
}

func (t *pathTrie[V]) FindString(pathStr string) *pathTrieNode[V] {
	return t.index[pathStr]
}

// Walk performs a depth-first post-order traversal of the trie, calling fn
// for each node. The root node is visited last.
func (t *pathTrie[V]) Walk(fn func(*pathTrieNode[V])) {
	var walk func(*pathTrieNode[V])
	walk = func(node *pathTrieNode[V]) {
		for _, child := range node.nodes {
			walk(child)
		}
		fn(node)
	}
	walk(t.root)
}

func (t *pathTrie[V]) WalkValues(fn func(node *pathTrieNode[V], value protoreflect.Value), rootValue protoreflect.Value) {
	var walk func(*pathTrieNode[V], protoreflect.Value)
	walk = func(node *pathTrieNode[V], value protoreflect.Value) {
		for _, child := range node.nodes {
			if !value.IsValid() {
				walk(child, protoreflect.Value{})
			} else {
				walk(child, value.Message().Get(child.Path.Index(-1).FieldDescriptor()))
			}
		}
		fn(node, value)
	}
	walk(t.root, rootValue)
}
