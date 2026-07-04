package execution

import (
	"encoding/json"
	"fmt"
	"time"
)

func NewEvent(ctx ExecutionContext, eventType string) ExecutionEvent {
	attributesCopy := make(map[string]string)

	for k, v := range ctx.Attributes {
		attributesCopy[k] = v
	}

	return ExecutionEvent{
		Type:         eventType,
		RequestID:    ctx.RequestID,
		TraceID:      ctx.TraceID,
		SpanID:       ctx.SpanID,
		ParentSpanID: ctx.ParentSpanID,
		Timestamp:    time.Now().UTC(),
		Attributes:   attributesCopy,
	}
}

func LogEvent(ctx ExecutionContext, event string) {
	e := NewEvent(ctx, event)

	RecordEvent(e)

	b, _ := json.Marshal(e)
	fmt.Println(string(b))
}
