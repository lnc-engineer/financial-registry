package execution

import "time"

// ContextKey defines safe context key usage
type ContextKey string

const ExecutionContextKey ContextKey = "execution_context"

// ExecutionContext represents a single unit of execution trace
type ExecutionContext struct {
	RequestID string
	StartTime time.Time
	Metadata   map[string]string
}