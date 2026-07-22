package execution

import "testing"

func TestTraceQueryBuilderBuild(t *testing.T) {
	filter := NewTraceQueryBuilder().
		WithTraceID("trace-123").
		WithStatus("success").
		Build()

	if filter.TraceID != "trace-123" {
		t.Fatalf("expected TraceID %q, got %q", "trace-123", filter.TraceID)
	}

	if filter.Status != "success" {
		t.Fatalf("expected Status %q, got %q", "success", filter.Status)
	}
}
