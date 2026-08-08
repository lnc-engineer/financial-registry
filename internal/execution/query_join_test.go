package execution

import "testing"

func TestApplyJoinMatchesRecords(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "trace-001",
			Attributes: map[string]string{
				"name": "transaction",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "trace-001",
			Attributes: map[string]string{
				"service": "payments",
			},
		},
	}

	results := ApplyJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "trace_id",
			RightField: "trace_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["right_service"] != "payments" {
		t.Fatalf("expected joined service attribute")
	}
}

func TestApplyJoinPreservesMatchedResults(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "trace-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "trace-002",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "USD",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "account_id",
		RightField: "account_id",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["right_currency"] != "GBP" {
		t.Fatalf("expected GBP, got %s", results[0].Attributes["right_currency"])
	}

	if results[1].Attributes["right_currency"] != "USD" {
		t.Fatalf("expected USD, got %s", results[1].Attributes["right_currency"])
	}
}

func TestApplyJoinPreservesUnmatchedLeftResults(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "trace-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "trace-002",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "account_id",
		RightField: "account_id",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "trace-001" {
		t.Fatalf("expected trace-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "trace-002" {
		t.Fatalf("expected trace-002, got %s", results[1].TraceID)
	}

	if _, exists := results[1].Attributes["right_currency"]; exists {
		t.Fatal("expected unmatched left result to have no right attributes")
	}
}

func TestApplyJoinPreservesAllLeftResultsWhenRightIsEmpty(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "trace-001",
		},
		{
			TraceID: "trace-002",
		},
	}

	right := []ExecutionContext{}

	condition := JoinCondition{
		LeftField:  "trace_id",
		RightField: "trace_id",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "trace-001" {
		t.Fatalf("expected trace-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "trace-002" {
		t.Fatalf("expected trace-002, got %s", results[1].TraceID)
	}
}

func TestApplyJoinReturnsEmptyWhenLeftIsEmpty(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
		},
	}

	condition := JoinCondition{
		LeftField:  "trace_id",
		RightField: "trace_id",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}
