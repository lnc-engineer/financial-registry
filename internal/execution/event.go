package execution

import "time"

// ExecutionEvent is the canonical event type for the system
type ExecutionEvent struct {
	Type      string    `json:"type"`
	RequestID string    `json:"request_id"`
	Timestamp time.Time `json:"timestamp"`
}
