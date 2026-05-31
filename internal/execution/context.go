package execution

import (
	"context"
	"time"
)

// ContextKey defines safe context key usage
type ContextKey string

const ExecutionContextKey ContextKey = "execution_context"

// ExecutionContext represents a single unit of execution trace
type ExecutionContext struct {
	RequestID string
	StartTime time.Time
	Metadata  map[string]string
}

// FromContext safely extracts execution context from request context
func FromContext(ctx context.Context) (ExecutionContext, bool) {
	ec, ok := ctx.Value(ExecutionContextKey).(ExecutionContext)
	return ec, ok
}
