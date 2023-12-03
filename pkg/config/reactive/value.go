package reactive

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type reactiveValue struct {
	mu            sync.RWMutex
	rev           int64
	pendingInit   bool
	value         protoreflect.Value
	watchChannels map[string]chan protoreflect.Value
	watchFuncs    map[string]func(int64, protoreflect.Value, <-chan struct{})
}

func newReactiveValue() *reactiveValue {
	return &reactiveValue{
		pendingInit:   true,
		watchChannels: make(map[string]chan protoreflect.Value),
		watchFuncs:    make(map[string]func(int64, protoreflect.Value, <-chan struct{})),
	}
}

func (r *reactiveValue) Update(rev int64, v protoreflect.Value, group <-chan struct{}, notify bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.rev = rev
	r.value = v

	if r.pendingInit {
		r.pendingInit = false
	} else if !notify {
		return
	}

	for _, w := range r.watchChannels {
		w <- v
	}

	for _, f := range r.watchFuncs {
		f(rev, v, group)
	}
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

	if r.rev != 0 && r.pendingInit {
		r.pendingInit = false
		ch <- r.value
	}

	key := uuid.NewString()
	r.watchChannels[key] = ch
	context.AfterFunc(ctx, func() {
		r.mu.Lock()
		defer r.mu.Unlock()
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

	if r.rev != 0 && r.pendingInit {
		r.pendingInit = false
		onChanged(r.rev, r.value, nil)
	}

	key := uuid.NewString()
	r.watchFuncs[key] = onChanged

	context.AfterFunc(ctx, func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		delete(r.watchFuncs, key)
	})
}
