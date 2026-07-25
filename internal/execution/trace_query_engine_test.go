package execution

import "testing"

func TestMatchesQueryByTraceID(t *testing.T) {
	ctx := ExecutionContext{
		TraceID: "trace-123",
	}

	conditions := []QueryCondition{
		{
			Field:    "trace_id",
			Operator: OperatorEquals,
			Value:    "trace-123",
		},
	}

	if !MatchesQuery(ctx, conditions) {
		t.Fatal("expected query to match")
	}
}

func TestMatchesQueryByStatus(t *testing.T) {
	ctx := ExecutionContext{
		Status: "success",
	}

	conditions := []QueryCondition{
		{
			Field:    "status",
			Operator: OperatorEquals,
			Value:    "success",
		},
	}

	if !MatchesQuery(ctx, conditions) {
		t.Fatal("expected query to match")
	}
}

func TestMatchesQueryByAttribute(t *testing.T) {
	ctx := ExecutionContext{
		Attributes: map[string]string{
			"service": "payments",
		},
	}

	conditions := []QueryCondition{
		{
			Field:    "service",
			Operator: OperatorEquals,
			Value:    "payments",
		},
	}

	if !MatchesQuery(ctx, conditions) {
		t.Fatal("expected attribute query to match")
	}
}

func TestMatchesQueryFails(t *testing.T) {
	ctx := ExecutionContext{
		Status: "failed",
	}

	conditions := []QueryCondition{
		{
			Field:    "status",
			Operator: OperatorEquals,
			Value:    "success",
		},
	}

	if MatchesQuery(ctx, conditions) {
		t.Fatal("expected query not to match")
	}
}

func TestMatchesMultipleConditions(t *testing.T) {
	ctx := ExecutionContext{
		TraceID: "trace-123",
		Status:  "success",
	}

	conditions := []QueryCondition{
		{
			Field:    "trace_id",
			Operator: OperatorEquals,
			Value:    "trace-123",
		},
		{
			Field:    "status",
			Operator: OperatorEquals,
			Value:    "success",
		},
	}

	if !MatchesQuery(ctx, conditions) {
		t.Fatal("expected all conditions to match")
	}
}
