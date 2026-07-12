package execution

import (
	"testing"
	"time"
)

func TestCalculateTraceStats(t *testing.T) {
	now := time.Now()

	spans := []ExecutionContext{
		{
			SpanID:       "1",
			ParentSpanID: "",
			Status:       "success",
			StartTime:    now,
			EndTime:      now.Add(10 * time.Millisecond),
		},
		{
			SpanID:       "2",
			ParentSpanID: "1",
			Status:       "failure",
			StartTime:    now,
			EndTime:      now.Add(20 * time.Millisecond),
		},
		{
			SpanID:       "3",
			ParentSpanID: "1",
			Status:       "",
			StartTime:    now,
			EndTime:      now.Add(30 * time.Millisecond),
		},
	}

	stats := CalculateTraceStats(spans)

	if stats.TotalSpans != 3 {
		t.Fatalf("expected 3 spans, got %d", stats.TotalSpans)
	}

	if stats.RootSpans != 1 {
		t.Fatalf("expected 1 root span, got %d", stats.RootSpans)
	}

	if stats.ChildSpans != 2 {
		t.Fatalf("expected 2 child spans, got %d", stats.ChildSpans)
	}

	if stats.SuccessSpans != 1 {
		t.Fatalf("expected 1 success span, got %d", stats.SuccessSpans)
	}

	if stats.FailureSpans != 1 {
		t.Fatalf("expected 1 failure span, got %d", stats.FailureSpans)
	}

	if stats.UnknownSpans != 1 {
		t.Fatalf("expected 1 unknown span, got %d", stats.UnknownSpans)
	}

	if stats.LongestDuration != 30*time.Millisecond {
		t.Fatalf("expected longest duration of 30ms, got %v", stats.LongestDuration)
	}

	if stats.ShortestDuration != 10*time.Millisecond {
		t.Fatalf("expected shortest duration of 10ms, got %v", stats.ShortestDuration)
	}

	if stats.AverageDuration != 20*time.Millisecond {
		t.Fatalf("expected average duration of 20ms, got %v", stats.AverageDuration)
	}
}
