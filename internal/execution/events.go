package execution

import (
	"time"
)

const ServiceName = "financial-registry"

// ExecutionEvent represents a single structured execution log event.
type ExecutionEvent struct {
	RequestID string                 `json:"request_id"`
	Event     string                 `json:"event"`
	Timestamp time.Time              `json:"timestamp"`
	Service   string                 `json:"service"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// NewEvent creates a new execution event with consistent defaults.
func NewEvent(ctx ExecutionContext, event string) ExecutionEvent {
	return ExecutionEvent{
		RequestID: ctx.RequestID,
		Event:     event,
		Timestamp: time.Now(),
		Service:   ServiceName,
		Metadata:  make(map[string]interface{}),
	}
}
