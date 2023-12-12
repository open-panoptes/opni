package reactive

import (
	"fmt"
	"io"
)

func (s *Controller[T]) DebugDumpReactiveMessagesInfo(out io.Writer) {
	s.reactiveMessagesMu.Lock()
	defer s.reactiveMessagesMu.Unlock()

	s.reactiveMessages.Walk(func(node *pathTrieNode[*reactiveValue]) {
		node.value.mu.RLock()
		defer node.value.mu.RUnlock()
		numWatchers := len(node.value.watchChannels) + len(node.value.watchFuncs)

		fmt.Fprintf(out, "message %s: {watchers: %d; rev: %d}\n", node.Path.String(), numWatchers, node.value.rev)
	})
}

var NewReactiveValue = newReactiveValue
