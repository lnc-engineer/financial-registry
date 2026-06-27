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

func FinishSpan(ec ExecutionContext) ExecutionContext {
	ec.EndTime = time.Now()
	return ec
}

func (ec ExecutionContext) Duration() time.Duration {
	if ec.EndTime.IsZero() {
		return time.Since(ec.StartTime)
	}
	return ec.EndTime.Sub(ec.StartTime)
}

func StartSpan(ec ExecutionContext) ExecutionContext {
	ec.StartTime = time.Now()
	return ec
}
