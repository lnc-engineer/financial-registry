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
