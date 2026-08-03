package execution

import "testing"

func TestApplyDistinctRemovesDuplicateValues(t *testing.T) {
	records := []ExecutionContext{
		{
			TraceID: "trace-1",
			Status:  "success",
		},
		{
			TraceID: "trace-2",
			Status:  "success",
		},
		{
			TraceID: "trace-3",
			Status:  "failure",
		},
	}

	result := ApplyDistinct(records, []string{"status"})

	if len(result) != 2 {
		t.Fatalf("expected 2 distinct records, got %d", len(result))
	}
}

func TestApplyDistinctSupportsMultipleFields(t *testing.T) {
	records := []ExecutionContext{
		{
			Status:   "success",
			SpanName: "parse",
		},
		{
			Status:   "success",
			SpanName: "parse",
		},
		{
			Status:   "success",
			SpanName: "validate",
		},
	}

	result := ApplyDistinct(records, []string{"status", "span_name"})

	if len(result) != 2 {
		t.Fatalf("expected 2 distinct records, got %d", len(result))
	}
}

func TestApplyDistinctWithEmptyProjectionReturnsOriginalRecords(t *testing.T) {
	records := []ExecutionContext{
		{
			Status: "success",
		},
		{
			Status: "success",
		},
	}

	result := ApplyDistinct(records, []string{})

	if len(result) != len(records) {
		t.Fatalf("expected original record count %d, got %d", len(records), len(result))
	}
}

func TestApplyDistinctWithUniqueRecords(t *testing.T) {
	records := []ExecutionContext{
		{
			Status: "success",
		},
		{
			Status: "failure",
		},
	}

	result := ApplyDistinct(records, []string{"status"})

	if len(result) != 2 {
		t.Fatalf("expected 2 records, got %d", len(result))
	}
}
