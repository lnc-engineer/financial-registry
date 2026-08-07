package execution

import "testing"

func TestApplyUnionCombinesResults(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
		{
			TraceID: "trace-002",
			Status:  "failed",
		},
		{
			TraceID: "trace-003",
			Status:  "success",
		},
	}

	query := UnionQuery{
		Queries: []CompositeFilter{
			{
				Status: "success",
			},
			{
				TraceID: "trace-002",
			},
		},
	}

	results := ApplyUnion(query, spans)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
}

func TestApplyUnionRemovesDuplicates(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
		{
			TraceID: "trace-002",
			Status:  "success",
		},
	}

	query := UnionQuery{
		Queries: []CompositeFilter{
			{
				Status: "success",
			},
			{
				TraceID: "trace-001",
			},
		},
	}

	results := ApplyUnion(query, spans)

	if len(results) != 2 {
		t.Fatalf("expected 2 unique results, got %d", len(results))
	}
}

func TestApplyUnionEmptyQueries(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
	}

	query := UnionQuery{
		Queries: []CompositeFilter{},
	}

	results := ApplyUnion(query, spans)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyUnionSingleQuery(t *testing.T) {
	spans := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
		{
			TraceID: "trace-002",
			Status:  "failed",
		},
	}

	query := UnionQuery{
		Queries: []CompositeFilter{
			{
				Status: "success",
			},
		},
	}

	results := ApplyUnion(query, spans)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].TraceID != "trace-001" {
		t.Fatalf("expected trace-001, got %s", results[0].TraceID)
	}
}
