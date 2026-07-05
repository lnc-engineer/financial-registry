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
	RequestID    string
	TraceID      string
	SpanID       string
	ParentSpanID string
	SpanName     string
	StartTime    time.Time
	EndTime      time.Time
	Status       string
	Lifecycle    []LifecycleEvent
	Attributes   map[string]string
}

// FromContext safely extracts execution context from request context
func FromContext(ctx context.Context) (ExecutionContext, bool) {
	ec, ok := ctx.Value(ExecutionContextKey).(ExecutionContext)
	return ec, ok
}

func (ec ExecutionContext) WithAttribute(key, value string) ExecutionContext {
	if ec.Attributes == nil {
		ec.Attributes = make(map[string]string)
	}

	ec.Attributes[key] = value
	return ec
}
