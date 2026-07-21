package execution

import "testing"

func TestCompositeFilterAppliesMultipleCriteria(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-1",
			Status:  "success",
			Lifecycle: []LifecycleEvent{
				{
					Name: "started",
				},
			},
		},
		{
			TraceID: "trace-1",
			Status:  "failure",
			Lifecycle: []LifecycleEvent{
				{
					Name: "failed",
				},
			},
		},
		{
			TraceID: "trace-2",
			Status:  "success",
			Lifecycle: []LifecycleEvent{
				{
					Name: "started",
				},
			},
		},
	}

	filter := CompositeFilter{
		TraceID:   "trace-1",
		Status:    "success",
		Lifecycle: "started",
	}

	result := filter.Apply(spans)

	if len(result) != 1 {
		t.Fatalf("expected 1 matching span, got %d", len(result))
	}

	if result[0].TraceID != "trace-1" {
		t.Fatalf("expected trace-1, got %s", result[0].TraceID)
	}
}
