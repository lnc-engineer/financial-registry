package execution

import "time"

// ExecutionEvent is the canonical event type for the system
type ExecutionEvent struct {
	Type         string            `json:"type"`
	RequestID    string            `json:"request_id"`
	TraceID      string            `json:"trace_id"`
	SpanID       string            `json:"span_id"`
	ParentSpanID string            `json:"parent_span_id"`
	Timestamp    time.Time         `json:"timestamp"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}
