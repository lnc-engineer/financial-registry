package execution

import "testing"

func TestFilterByLifecycle(t *testing.T) {
	spans := []ExecutionContext{
		{
			Lifecycle: []LifecycleEvent{
				{Name: "started"},
			},
		},
		{
			Lifecycle: []LifecycleEvent{
				{Name: "completed"},
			},
		},
		{
			Lifecycle: []LifecycleEvent{
				{Name: "started"},
			},
		},
		{
			Lifecycle: []LifecycleEvent{
				{Name: "failed"},
			},
		},
	}

	if got := len(FilterByLifecycle(spans, "started")); got != 2 {
		t.Fatalf("expected 2 started spans, got %d", got)
	}

	if got := len(FilterByLifecycle(spans, "completed")); got != 1 {
		t.Fatalf("expected 1 completed span, got %d", got)
	}

	if got := len(FilterByLifecycle(spans, "failed")); got != 1 {
		t.Fatalf("expected 1 failed span, got %d", got)
	}

	if got := len(FilterByLifecycle(spans, "unknown")); got != 0 {
		t.Fatalf("expected 0 unknown spans, got %d", got)
	}
}
