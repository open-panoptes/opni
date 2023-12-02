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
	value         protoreflect.Value
	watchChannels map[string]chan protoreflect.Value
	watchFuncs    map[string]func(int64, protoreflect.Value)
	group         <-chan struct{}
	newGroup      chan struct{}
}

func newReactiveValue() *reactiveValue {
	return &reactiveValue{
		watchChannels: make(map[string]chan protoreflect.Value),
		watchFuncs:    make(map[string]func(int64, protoreflect.Value)),
		newGroup:      make(chan struct{}),
	}
}

func (r *reactiveValue) Update(rev int64, v protoreflect.Value, group <-chan struct{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.rev = rev
	r.value = v
	if r.group != nil {
		select {
		case r.newGroup <- struct{}{}:
		default:
		}
	}
	r.group = group

	for _, w := range r.watchChannels {
		w <- v
	}

	for _, f := range r.watchFuncs {
		f(rev, v)
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

	if r.rev != 0 {
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
	r.watchFuncWithRev(ctx, func(_ int64, value protoreflect.Value) {
		onChanged(value)
	})
}

func (r *reactiveValue) watchFuncWithRev(ctx context.Context, onChanged func(int64, protoreflect.Value)) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.rev != 0 {
		defer onChanged(r.rev, r.value)
	}

	key := uuid.NewString()
	r.watchFuncs[key] = onChanged

	context.AfterFunc(ctx, func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		delete(r.watchFuncs, key)
	})
}

func (r *reactiveValue) wait() {
	for {
		r.mu.Lock()
		group := r.group
		if group == nil {
			return
		}
		r.mu.Unlock()
		select {
		case <-group:
			return
		case <-r.newGroup:
			continue
		}
	}
}
