package reactive

import (
	"fmt"
	"io"

	"go.etcd.io/etcd/pkg/adt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (s *Controller[T]) DebugDumpReactiveMessagesInfo(out io.Writer) {
	s.reactiveMessagesMu.Lock()
	defer s.reactiveMessagesMu.Unlock()

	s.reactiveMessages.Visit(adt.NewStringAffineInterval("\x00", ""), func(iv *adt.IntervalValue) bool {
		rm := iv.Val.(*reactiveValue)
		key := iv.Ivl.Begin.(adt.StringAffineComparable)
		numWatchers := len(rm.watchChannels)

		rm.watchFuncs.Range(func(_ string, _ *func(int64, protoreflect.Value)) bool {
			numWatchers++
			return true
		})
		fmt.Fprintf(out, "message %s: {watchers: %d; rev: %d}\n", key, numWatchers, rm.rev)

		return true
	})
}

type ReactiveValue = reactiveValue
