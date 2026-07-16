package execution

import "testing"

func sampleSpans() []ExecutionContext {
	return []ExecutionContext{
		{
			TraceID:      "trace-1",
			SpanID:       "span-1",
			Status:       "success",
			ParentSpanID: "",
		},
		{
			TraceID:      "trace-1",
			SpanID:       "span-2",
			Status:       "failure",
			ParentSpanID: "span-1",
		},
		{
			TraceID:      "trace-2",
			SpanID:       "span-3",
			Status:       "success",
			ParentSpanID: "",
		},
	}
}

func TestFilterByTraceID(t *testing.T) {
	result := FilterByTraceID(sampleSpans(), "trace-1")

	if len(result) != 2 {
		t.Fatalf("expected 2 spans, got %d", len(result))
	}
}

func TestFilterByStatus(t *testing.T) {
	result := FilterByStatus(sampleSpans(), "success")

	if len(result) != 2 {
		t.Fatalf("expected 2 spans, got %d", len(result))
	}
}

func TestFilterRootSpans(t *testing.T) {
	result := FilterRootSpans(sampleSpans())

	if len(result) != 2 {
		t.Fatalf("expected 2 root spans, got %d", len(result))
	}
}
