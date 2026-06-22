package execution

import (
	"fmt"
	"time"
)

func NewChildSpan(parent ExecutionContext, spanName string) ExecutionContext {
	return ExecutionContext{
		RequestID:    parent.RequestID,
		TraceID:      parent.TraceID,
		SpanID:       fmt.Sprintf("span-%d", time.Now().UnixNano()),
		ParentSpanID: parent.SpanID,
		SpanName: 	  spanName,
		StartTime:    time.Now(),
		Metadata:     make(map[string]string),
	}
}

func LogSpan(label string, ec ExecutionContext) {
	fmt.Printf(
		"[SPAN:%s] name=%s trace_id=%s span_id=%s parent_span_id=%s request_id=%s\n",
		label,
		ec.SpanName,
		ec.TraceID,
		ec.SpanID,
		ec.ParentSpanID,
		ec.RequestID,
	)
}
