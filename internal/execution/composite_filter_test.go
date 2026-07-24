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

func TestCompositeFilterSupportsOperators(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
		{
			TraceID: "trace-002",
			Status:  "failure",
		},
		{
			TraceID: "payment-001",
			Status:  "successful",
		},
	}

	filter := CompositeFilter{
		TraceID:         "trace-",
		TraceIDOperator: OperatorStartsWith,
	}

	result := filter.Apply(spans)

	if len(result) != 2 {
		t.Fatalf("expected 2 matching spans, got %d", len(result))
	}
}

func TestCompositeFilterSupportsStatusOperators(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
		{
			TraceID: "trace-002",
			Status:  "failure",
		},
		{
			TraceID: "trace-003",
			Status:  "successful",
		},
	}

	filter := CompositeFilter{
		Status:         "succ",
		StatusOperator: OperatorStartsWith,
	}

	result := filter.Apply(spans)

	if len(result) != 2 {
		t.Fatalf("expected 2 matching spans, got %d", len(result))
	}
}

func TestCompositeFilterSupportsLifecycleOperators(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Lifecycle: []LifecycleEvent{
				{Name: "started"},
			},
		},
		{
			TraceID: "trace-002",
			Lifecycle: []LifecycleEvent{
				{Name: "failed"},
			},
		},
		{
			TraceID: "trace-003",
			Lifecycle: []LifecycleEvent{
				{Name: "starting"},
			},
		},
	}

	filter := CompositeFilter{
		Lifecycle:         "start",
		LifecycleOperator: OperatorStartsWith,
	}

	result := filter.Apply(spans)

	if len(result) != 2 {
		t.Fatalf("expected 2 matching spans, got %d", len(result))
	}
}
