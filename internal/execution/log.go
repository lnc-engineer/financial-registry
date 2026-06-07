package execution

import (
	"encoding/json"
	"fmt"
	"time"
)

func NewEvent(ctx ExecutionContext, eventType string) ExecutionEvent {
	return ExecutionEvent{
		Type:      eventType,
		RequestID: ctx.RequestID,
		Timestamp: time.Now().UTC(),
	}
}

func LogEvent(ctx ExecutionContext, event string) {
	e := NewEvent(ctx, event)

	RecordEvent(e)

	b, _ := json.Marshal(e)
	fmt.Println(string(b))
}
