package reactive

import (
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (s *Controller[T]) DebugDumpReactiveMessagesInfo(out io.Writer) {
	s.reactiveMessagesMu.Lock()
	defer s.reactiveMessagesMu.Unlock()

	s.reactiveMessages.Walk(func(node *pathTrieNode[*reactiveValue]) {
		numWatchers := len(node.value.watchChannels)

		node.value.watchFuncs.Range(func(_ string, _ *func(int64, protoreflect.Value)) bool {
			numWatchers++
			return true
		})
		fmt.Fprintf(out, "message %s: {watchers: %d; rev: %d}\n", node.Path.String(), numWatchers, node.value.rev)
	})
}

var NewReactiveValue = newReactiveValue
