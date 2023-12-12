package reactive

import (
	"context"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type reactiveValue struct {
	mu            sync.RWMutex
	logger        *slog.Logger
	rev           int64
	pendingInit   bool
	value         protoreflect.Value
	watchChannels map[string]chan protoreflect.Value
	watchFuncs    map[string]func(int64, protoreflect.Value, <-chan struct{})
}

func newReactiveValue(lg *slog.Logger) *reactiveValue {
	return &reactiveValue{
		logger:        lg,
		pendingInit:   true,
		watchChannels: make(map[string]chan protoreflect.Value),
		watchFuncs:    make(map[string]func(int64, protoreflect.Value, <-chan struct{})),
	}
}

func (s *reactiveValue) traceLog(msg string, attrs ...any) {
	if s.logger == nil {
		return
	}
	s.logger.Log(context.Background(), traceLogLevel, msg, attrs...)
}

func (r *reactiveValue) Update(rev int64, v protoreflect.Value, group <-chan struct{}, notify bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.traceLog("updating reactive value", "revision", rev, "group", group, "notify", notify)

	r.rev = rev
	r.value = v

	if r.pendingInit {
		r.pendingInit = false
		r.traceLog("pending init complete")
	} else if !notify {
		r.traceLog("skipping notify")
		return
	}

	r.traceLog("notifying watch channels", "count", len(r.watchChannels))
	for _, w := range r.watchChannels {
		w <- v
	}

	r.traceLog("notifying watch funcs", "count", len(r.watchFuncs))
	for _, f := range r.watchFuncs {
		f(rev, v, group)
	}

	r.traceLog("update complete")
}

func (r *reactiveValue) Value() protoreflect.Value {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.value
}

func (r *reactiveValue) Watch(ctx context.Context) <-chan protoreflect.Value {
	r.mu.Lock()
	defer r.mu.Unlock()
	ch := make(chan protoreflect.Value, 4)

	r.traceLog("starting watch", "pendingInit", r.pendingInit)

	if !r.pendingInit {
		r.traceLog("not pending init; writing current value", "revision", r.rev, "value", r.value)
		ch <- r.value
	}

	key := uuid.NewString()
	r.watchChannels[key] = ch
	r.traceLog("adding watch channel", "key", key)
	context.AfterFunc(ctx, func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		r.traceLog("removing watch channel", "key", key)
		delete(r.watchChannels, key)
		close(ch)
	})
	return ch
}

func (r *reactiveValue) WatchFunc(ctx context.Context, onChanged func(protoreflect.Value)) {
	r.watchFuncInternal(ctx, func(_ int64, value protoreflect.Value, _ <-chan struct{}) {
		onChanged(value)
	})
}

func (r *reactiveValue) watchFuncInternal(ctx context.Context, onChanged func(int64, protoreflect.Value, <-chan struct{})) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.traceLog("starting watch func", "pendingInit", r.pendingInit)

	if !r.pendingInit {
		r.traceLog("not pending init; calling watch func", "revision", r.rev, "value", r.value)
		onChanged(r.rev, r.value, nil)
	}

	key := uuid.NewString()
	r.watchFuncs[key] = onChanged
	r.traceLog("adding watch func", "key", key)
	context.AfterFunc(ctx, func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		r.traceLog("removing watch func", "key", key)
		delete(r.watchFuncs, key)
	})
}
