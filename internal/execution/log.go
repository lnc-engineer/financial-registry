package execution

import (
	"encoding/json"
	"fmt"
	"time"
)

func NewEvent(ctx ExecutionContext, eventType string) ExecutionEvent {
	metadataCopy := make(map[string]string)

	for k, v := range ctx.Metadata {
		metadataCopy[k] = v
	}

	return ExecutionEvent{
		Type:      eventType,
		RequestID: ctx.RequestID,
		TraceID:   ctx.TraceID,
		SpanID:    ctx.SpanID,
		Timestamp: time.Now().UTC(),
		Metadata:  metadataCopy,
	}
}

func LogEvent(ctx ExecutionContext, event string) {
	e := NewEvent(ctx, event)

	RecordEvent(e)

	b, _ := json.Marshal(e)
	fmt.Println(string(b))
}
