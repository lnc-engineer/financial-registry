package execution

import "sync"

var (
	eventMu sync.RWMutex
	events  []ExecutionEvent
)

func RecordEvent(event ExecutionEvent) {
	eventMu.Lock()
	defer eventMu.Unlock()

	events = append(events, event)
}

func Events() []ExecutionEvent {
	eventMu.RLock()
	defer eventMu.RUnlock()

	result := make([]ExecutionEvent, len(events))
	copy(result, events)

	return result
}
