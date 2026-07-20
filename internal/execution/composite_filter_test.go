package execution

import "testing"

func TestCompositeFilterAppliesMultipleFilters(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
		{
			TraceID: "trace-001",
			Status:  "failure",
		},
		{
			TraceID: "trace-002",
			Status:  "success",
		},
	}

	filter := CompositeFilter{
		TraceID: "trace-001",
		Status:  "success",
	}

	result := filter.Apply(spans)

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	if result[0].Status != "success" {
		t.Fatal("expected success status")
	}
}
