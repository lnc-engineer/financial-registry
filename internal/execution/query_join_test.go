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

func TestApplyJoinSupportsMultipleRightMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
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
				"account_id": "account-001",
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

func TestApplyJoinMatchesMultipleConditions(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
				"status":     "active",
			},
		},
	}

	results := ApplyJoin(
		left,
		right,
		JoinCondition{
			Conditions: []JoinPredicate{
				{
					LeftField:  "account_id",
					RightField: "account_id",
				},
				{
					LeftField:  "currency",
					RightField: "currency",
				},
			},
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["right_status"] != "active" {
		t.Fatalf("expected active status, got %s", results[0].Attributes["right_status"])
	}
}

func TestApplyJoinRejectsWhenOneConditionFails(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "USD",
			},
		},
	}

	results := ApplyJoin(
		left,
		right,
		JoinCondition{
			Conditions: []JoinPredicate{
				{
					LeftField:  "account_id",
					RightField: "account_id",
				},
				{
					LeftField:  "currency",
					RightField: "currency",
				},
			},
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].TraceID != "left-001" {
		t.Fatalf("expected unmatched left context, got %s", results[0].TraceID)
	}

	if _, exists := results[0].Attributes["right_currency"]; exists {
		t.Fatal("expected no right attributes when one join condition fails")
	}
}

func TestApplyJoinMultipleConditionsSupportsMultipleMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "USD",
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
				"status":     "active",
			},
		},
		{
			TraceID: "right-003",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
				"status":     "pending",
			},
		},
	}

	results := ApplyJoin(
		left,
		right,
		JoinCondition{
			Conditions: []JoinPredicate{
				{
					LeftField:  "account_id",
					RightField: "account_id",
				},
				{
					LeftField:  "currency",
					RightField: "currency",
				},
			},
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["right_status"] != "active" {
		t.Fatalf("expected active, got %s", results[0].Attributes["right_status"])
	}

	if results[1].Attributes["right_status"] != "pending" {
		t.Fatalf("expected pending, got %s", results[1].Attributes["right_status"])
	}
}

func TestApplyJoinPreservesLegacySingleConditionBehavior(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"status":     "active",
			},
		},
	}

	results := ApplyJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["right_status"] != "active" {
		t.Fatalf("expected active, got %s", results[0].Attributes["right_status"])
	}
}

func TestApplyRightJoinMatchesRecords(t *testing.T) {
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

	results := ApplyRightJoin(
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

	if results[0].Attributes["left_name"] != "transaction" {
		t.Fatalf("expected joined name attribute")
	}
}

func TestApplyRightJoinPreservesUnmatchedRightResults(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
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

	results := ApplyRightJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "right-001" {
		t.Fatalf("expected right-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[1].TraceID)
	}

	if results[0].Attributes["left_account_id"] != "account-001" {
		t.Fatal("expected matched right result to contain left attributes")
	}

	if _, exists := results[1].Attributes["left_account_id"]; exists {
		t.Fatal("expected unmatched right result to have no left attributes")
	}
}

func TestApplyRightJoinPreservesAllRightResultsWhenLeftIsEmpty(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
		},
		{
			TraceID: "right-002",
		},
	}

	condition := JoinCondition{
		LeftField:  "trace_id",
		RightField: "trace_id",
	}

	results := ApplyRightJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "right-001" {
		t.Fatalf("expected right-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[1].TraceID)
	}
}

func TestApplyRightJoinReturnsEmptyWhenRightIsEmpty(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
		},
	}

	right := []ExecutionContext{}

	condition := JoinCondition{
		LeftField:  "trace_id",
		RightField: "trace_id",
	}

	results := ApplyRightJoin(left, right, condition)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyRightJoinSupportsMultipleLeftMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"source":     "transaction",
			},
		},
		{
			TraceID: "left-002",
			Attributes: map[string]string{
				"account_id": "account-001",
				"source":     "payment",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "account_id",
		RightField: "account_id",
	}

	results := ApplyRightJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["left_source"] != "transaction" {
		t.Fatalf("expected transaction, got %s", results[0].Attributes["left_source"])
	}

	if results[1].Attributes["left_source"] != "payment" {
		t.Fatalf("expected payment, got %s", results[1].Attributes["left_source"])
	}
}

func TestApplyFullOuterJoinMatchesRecords(t *testing.T) {
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

	results := ApplyFullOuterJoin(
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

func TestApplyFullOuterJoinPreservesUnmatchedLeftResults(t *testing.T) {
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

	results := ApplyFullOuterJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

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

func TestApplyFullOuterJoinPreservesUnmatchedRightResults(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
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

	results := ApplyFullOuterJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "left-001" {
		t.Fatalf("expected left-001, got %s", results[0].TraceID)
	}

	if results[0].Attributes["right_currency"] != "GBP" {
		t.Fatalf("expected GBP, got %s", results[0].Attributes["right_currency"])
	}

	if results[1].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[1].TraceID)
	}

	if _, exists := results[1].Attributes["left_account_id"]; exists {
		t.Fatal("expected unmatched right result to have no left attributes")
	}
}

func TestApplyFullOuterJoinPreservesBothUnmatchedSides(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	results := ApplyFullOuterJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "left-001" {
		t.Fatalf("expected left-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "right-001" {
		t.Fatalf("expected right-001, got %s", results[1].TraceID)
	}
}

func TestApplyFullOuterJoinPreservesAllResultsWhenLeftIsEmpty(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
		},
		{
			TraceID: "right-002",
		},
	}

	results := ApplyFullOuterJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "trace_id",
			RightField: "trace_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "right-001" {
		t.Fatalf("expected right-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[1].TraceID)
	}
}

func TestApplyFullOuterJoinPreservesAllResultsWhenRightIsEmpty(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
		},
		{
			TraceID: "left-002",
		},
	}

	right := []ExecutionContext{}

	results := ApplyFullOuterJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "trace_id",
			RightField: "trace_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "left-001" {
		t.Fatalf("expected left-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "left-002" {
		t.Fatalf("expected left-002, got %s", results[1].TraceID)
	}
}

func TestApplyFullOuterJoinSupportsMultipleMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "left-002",
			Attributes: map[string]string{
				"account_id": "account-001",
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
				"account_id": "account-001",
				"currency":   "USD",
			},
		},
	}

	results := ApplyFullOuterJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}
}

func TestApplyCrossJoinProducesCartesianProduct(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "left-002",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"currency": "GBP",
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"currency": "USD",
			},
		},
	}

	results := ApplyCrossJoin(left, right)

	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}
}

func TestApplyCrossJoinPreservesLeftAttributes(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"currency": "GBP",
			},
		},
	}

	results := ApplyCrossJoin(left, right)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["account_id"] != "account-001" {
		t.Fatalf("expected account-001, got %s", results[0].Attributes["account_id"])
	}
}

func TestApplyCrossJoinPrefixesRightAttributes(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"currency": "GBP",
			},
		},
	}

	results := ApplyCrossJoin(left, right)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["right_currency"] != "GBP" {
		t.Fatalf("expected GBP, got %s", results[0].Attributes["right_currency"])
	}
}

func TestApplyCrossJoinPreservesAllCombinations(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
		},
		{
			TraceID: "left-002",
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
		},
		{
			TraceID: "right-002",
		},
		{
			TraceID: "right-003",
		},
	}

	results := ApplyCrossJoin(left, right)

	if len(results) != 6 {
		t.Fatalf("expected 6 results, got %d", len(results))
	}
}

func TestApplyCrossJoinReturnsEmptyWhenLeftIsEmpty(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
		},
		{
			TraceID: "right-002",
		},
	}

	results := ApplyCrossJoin(left, right)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyCrossJoinReturnsEmptyWhenRightIsEmpty(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
		},
		{
			TraceID: "left-002",
		},
	}

	right := []ExecutionContext{}

	results := ApplyCrossJoin(left, right)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplySelfJoinMatchesRecords(t *testing.T) {
	contexts := []ExecutionContext{
		{
			TraceID: "trace-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"role":       "transaction",
			},
		},
		{
			TraceID: "trace-002",
			Attributes: map[string]string{
				"account_id": "account-002",
				"role":       "payment",
			},
		},
	}

	results := ApplySelfJoin(
		contexts,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["role"] != "transaction" {
		t.Fatalf("expected transaction role, got %s", results[0].Attributes["role"])
	}

	if results[0].Attributes["right_role"] != "transaction" {
		t.Fatalf("expected right_transaction role, got %s", results[0].Attributes["right_role"])
	}
}

func TestApplySelfJoinSupportsMultipleMatches(t *testing.T) {
	contexts := []ExecutionContext{
		{
			TraceID: "trace-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"role":       "transaction",
			},
		},
		{
			TraceID: "trace-002",
			Attributes: map[string]string{
				"account_id": "account-001",
				"role":       "payment",
			},
		},
	}

	results := ApplySelfJoin(
		contexts,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}
}

func TestApplySelfJoinReturnsEmptyWhenInputIsEmpty(t *testing.T) {
	contexts := []ExecutionContext{}

	results := ApplySelfJoin(
		contexts,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyNaturalJoinMatchesSharedFields(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"role":       "transaction",
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

	results := ApplyNaturalJoin(left, right)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["role"] != "transaction" {
		t.Fatalf("expected transaction role, got %s", results[0].Attributes["role"])
	}

	if results[0].Attributes["right_currency"] != "GBP" {
		t.Fatalf("expected GBP currency, got %s", results[0].Attributes["right_currency"])
	}
}

func TestApplyNaturalJoinRejectsMismatchedSharedFields(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	results := ApplyNaturalJoin(left, right)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyNaturalJoinRequiresAllSharedFieldsToMatch(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "USD",
			},
		},
	}

	results := ApplyNaturalJoin(left, right)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyNaturalJoinSupportsMultipleMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "USD",
			},
		},
	}

	results := ApplyNaturalJoin(left, right)

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

func TestApplyNaturalJoinWithNoCommonFieldsProducesCartesianProduct(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"currency": "GBP",
			},
		},
		{
			Attributes: map[string]string{
				"currency": "USD",
			},
		},
	}

	results := ApplyNaturalJoin(left, right)

	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}
}

func TestApplyNaturalJoinReturnsEmptyWhenInputIsEmpty(t *testing.T) {
	left := []ExecutionContext{}
	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	results := ApplyNaturalJoin(left, right)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyLeftAntiJoinExcludesMatchedRecords(t *testing.T) {
	left := []ExecutionContext{
		{TraceID: "left-001", Attributes: map[string]string{"account_id": "account-001"}},
		{TraceID: "left-002", Attributes: map[string]string{"account_id": "account-002"}},
	}
	right := []ExecutionContext{
		{TraceID: "right-001", Attributes: map[string]string{"account_id": "account-001"}},
	}

	results := ApplyLeftAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].TraceID != "left-002" {
		t.Fatalf("expected left-002, got %s", results[0].TraceID)
	}
}

func TestApplyLeftAntiJoinReturnsAllLeftWhenRightIsEmpty(t *testing.T) {
	left := []ExecutionContext{
		{TraceID: "left-001"},
		{TraceID: "left-002"},
	}
	right := []ExecutionContext{}

	results := ApplyLeftAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "trace_id",
			RightField: "trace_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}

func TestApplyLeftAntiJoinReturnsEmptyWhenLeftIsEmpty(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	results := ApplyLeftAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyRightAntiJoinExcludesMatchedRecords(t *testing.T) {
	left := []ExecutionContext{
		{TraceID: "left-001", Attributes: map[string]string{"account_id": "account-001"}},
	}
	right := []ExecutionContext{
		{TraceID: "right-001", Attributes: map[string]string{"account_id": "account-001"}},
		{TraceID: "right-002", Attributes: map[string]string{"account_id": "account-002"}},
	}

	results := ApplyRightAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[0].TraceID)
	}
}

func TestApplyRightAntiJoinReturnsAllRightWhenLeftIsEmpty(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{TraceID: "right-001"},
		{TraceID: "right-002"},
	}

	results := ApplyRightAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "trace_id",
			RightField: "trace_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "right-001" {
		t.Fatalf("expected right-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[1].TraceID)
	}
}

func TestApplyRightAntiJoinReturnsEmptyWhenRightIsEmpty(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{}

	results := ApplyRightAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyLeftAntiJoinReturnsAllLeftWhenThereAreNoMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "left-002",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-003",
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"account_id": "account-004",
			},
		},
	}

	results := ApplyLeftAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "left-001" {
		t.Fatalf("expected left-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "left-002" {
		t.Fatalf("expected left-002, got %s", results[1].TraceID)
	}
}

func TestApplyRightAntiJoinReturnsAllRightWhenThereAreNoMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "left-002",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-003",
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"account_id": "account-004",
			},
		},
	}

	results := ApplyRightAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].TraceID != "right-001" {
		t.Fatalf("expected right-001, got %s", results[0].TraceID)
	}

	if results[1].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[1].TraceID)
	}
}

func TestApplyLeftAntiJoinPreservesAttributes(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
				"status":     "active",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	results := ApplyLeftAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["account_id"] != "account-001" {
		t.Fatalf("expected account-001, got %s", results[0].Attributes["account_id"])
	}

	if results[0].Attributes["currency"] != "GBP" {
		t.Fatalf("expected GBP, got %s", results[0].Attributes["currency"])
	}

	if results[0].Attributes["status"] != "active" {
		t.Fatalf("expected active, got %s", results[0].Attributes["status"])
	}

	if _, exists := results[0].Attributes["right_account_id"]; exists {
		t.Fatal("expected no right attributes in anti join result")
	}
}

func TestApplyRightAntiJoinPreservesAttributes(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "USD",
				"status":     "pending",
			},
		},
	}

	results := ApplyRightAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["account_id"] != "account-002" {
		t.Fatalf("expected account-002, got %s", results[0].Attributes["account_id"])
	}

	if results[0].Attributes["currency"] != "USD" {
		t.Fatalf("expected USD, got %s", results[0].Attributes["currency"])
	}

	if results[0].Attributes["status"] != "pending" {
		t.Fatalf("expected pending, got %s", results[0].Attributes["status"])
	}

	if _, exists := results[0].Attributes["left_account_id"]; exists {
		t.Fatal("expected no left attributes in anti join result")
	}
}

func TestApplyLeftAntiJoinExcludesRecordsWithMultipleMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "left-002",
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
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	results := ApplyLeftAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].TraceID != "left-002" {
		t.Fatalf("expected left-002, got %s", results[0].TraceID)
	}
}

func TestApplyRightAntiJoinExcludesRecordsWithMultipleMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "left-002",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"account_id": "account-002",
			},
		},
	}

	results := ApplyRightAntiJoin(
		left,
		right,
		JoinCondition{
			LeftField:  "account_id",
			RightField: "account_id",
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[0].TraceID)
	}
}

func TestApplySemiJoin(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "1", "name": "Alice"}},
		{Attributes: map[string]string{"id": "2", "name": "Bob"}},
		{Attributes: map[string]string{"id": "3", "name": "Charlie"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "1", "status": "active"}},
		{Attributes: map[string]string{"id": "3", "status": "active"}},
	}

	condition := JoinCondition{
		LeftField:  "id",
		RightField: "id",
	}

	results := ApplySemiJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["id"] != "1" {
		t.Fatalf("expected first result id 1, got %q", results[0].Attributes["id"])
	}

	if results[1].Attributes["id"] != "3" {
		t.Fatalf("expected second result id 3, got %q", results[1].Attributes["id"])
	}

	if _, exists := results[0].Attributes["right_status"]; exists {
		t.Fatal("semi join should not include right-side attributes")
	}
}

func TestApplySemiJoinMultipleRightMatches(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "1", "name": "Alice"}},
		{Attributes: map[string]string{"id": "2", "name": "Bob"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "1", "status": "active"}},
		{Attributes: map[string]string{"id": "1", "status": "pending"}},
		{Attributes: map[string]string{"id": "2", "status": "active"}},
	}

	condition := JoinCondition{
		LeftField:  "id",
		RightField: "id",
	}

	results := ApplySemiJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["id"] != "1" {
		t.Fatalf("expected first result id 1, got %q", results[0].Attributes["id"])
	}

	if results[1].Attributes["id"] != "2" {
		t.Fatalf("expected second result id 2, got %q", results[1].Attributes["id"])
	}
}

func TestApplySemiJoinNoMatches(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "1"}},
		{Attributes: map[string]string{"id": "2"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "3"}},
		{Attributes: map[string]string{"id": "4"}},
	}

	condition := JoinCondition{
		LeftField:  "id",
		RightField: "id",
	}

	results := ApplySemiJoin(left, right, condition)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplySemiJoinEmptyRight(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "1"}},
		{Attributes: map[string]string{"id": "2"}},
	}

	var right []ExecutionContext

	condition := JoinCondition{
		LeftField:  "id",
		RightField: "id",
	}

	results := ApplySemiJoin(left, right, condition)

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplySemiJoinPreservesLeftAttributes(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":         "1",
				"name":       "Alice",
				"department": "finance",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":     "1",
				"status": "active",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "id",
		RightField: "id",
	}

	results := ApplySemiJoin(left, right, condition)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "1" {
		t.Fatalf("expected id 1, got %q", results[0].Attributes["id"])
	}

	if results[0].Attributes["name"] != "Alice" {
		t.Fatalf("expected name Alice, got %q", results[0].Attributes["name"])
	}

	if results[0].Attributes["department"] != "finance" {
		t.Fatalf("expected department finance, got %q", results[0].Attributes["department"])
	}

	if _, exists := results[0].Attributes["right_status"]; exists {
		t.Fatal("semi join should not include right-side attributes")
	}
}

func TestApplyRightJoinMatchesMultipleConditions(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
				"status":     "active",
			},
		},
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "USD",
				"status":     "inactive",
			},
		},
	}

	results := ApplyRightJoin(
		left,
		right,
		JoinCondition{
			Conditions: []JoinPredicate{
				{
					LeftField:  "account_id",
					RightField: "account_id",
				},
				{
					LeftField:  "currency",
					RightField: "currency",
				},
			},
		},
	)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["left_currency"] != "GBP" {
		t.Fatalf("expected GBP, got %s", results[0].Attributes["left_currency"])
	}

	if results[0].Attributes["status"] != "active" {
		t.Fatalf("expected active, got %s", results[0].Attributes["status"])
	}

	if _, exists := results[1].Attributes["left_currency"]; exists {
		t.Fatal("expected unmatched right result to have no left attributes")
	}
}

func TestApplyFullOuterJoinMatchesMultipleConditions(t *testing.T) {
	left := []ExecutionContext{
		{
			TraceID: "left-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
		{
			TraceID: "left-002",
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "USD",
			},
		},
	}

	right := []ExecutionContext{
		{
			TraceID: "right-001",
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
				"status":     "active",
			},
		},
		{
			TraceID: "right-002",
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "EUR",
				"status":     "pending",
			},
		},
	}

	results := ApplyFullOuterJoin(
		left,
		right,
		JoinCondition{
			Conditions: []JoinPredicate{
				{
					LeftField:  "account_id",
					RightField: "account_id",
				},
				{
					LeftField:  "currency",
					RightField: "currency",
				},
			},
		},
	)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	if results[0].TraceID != "left-001" {
		t.Fatalf("expected left-001, got %s", results[0].TraceID)
	}

	if results[0].Attributes["right_status"] != "active" {
		t.Fatalf("expected active, got %s", results[0].Attributes["right_status"])
	}

	if results[1].TraceID != "left-002" {
		t.Fatalf("expected left-002, got %s", results[1].TraceID)
	}

	if results[2].TraceID != "right-002" {
		t.Fatalf("expected right-002, got %s", results[2].TraceID)
	}
}

func TestApplyLeftAntiJoinMatchesMultipleConditions(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
		{
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "USD",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
		{
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "EUR",
			},
		},
	}

	results := ApplyLeftAntiJoin(
		left,
		right,
		JoinCondition{
			Conditions: []JoinPredicate{
				{
					LeftField:  "account_id",
					RightField: "account_id",
				},
				{
					LeftField:  "currency",
					RightField: "currency",
				},
			},
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["account_id"] != "account-002" {
		t.Fatalf(
			"expected account-002, got %s",
			results[0].Attributes["account_id"],
		)
	}
}

func TestApplyRightAntiJoinMatchesMultipleConditions(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
		{
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "USD",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"account_id": "account-001",
				"currency":   "GBP",
			},
		},
		{
			Attributes: map[string]string{
				"account_id": "account-002",
				"currency":   "EUR",
			},
		},
	}

	results := ApplyRightAntiJoin(
		left,
		right,
		JoinCondition{
			Conditions: []JoinPredicate{
				{
					LeftField:  "account_id",
					RightField: "account_id",
				},
				{
					LeftField:  "currency",
					RightField: "currency",
				},
			},
		},
	)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["account_id"] != "account-002" {
		t.Fatalf(
			"expected account-002, got %s",
			results[0].Attributes["account_id"],
		)
	}
}

func TestApplyJoinOneToManyMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "L1",
				"code": "A",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "R1",
				"code": "A",
			},
		},
		{
			Attributes: map[string]string{
				"id":   "R2",
				"code": "A",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "code",
		RightField: "code",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 joined results, got %d", len(results))
	}
}

func TestApplyJoinManyToOneMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "L1",
				"code": "A",
			},
		},
		{
			Attributes: map[string]string{
				"id":   "L2",
				"code": "A",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "R1",
				"code": "A",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "code",
		RightField: "code",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 joined results, got %d", len(results))
	}
}

func TestApplyJoinManyToManyMatches(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "L1",
				"code": "A",
			},
		},
		{
			Attributes: map[string]string{
				"id":   "L2",
				"code": "A",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "R1",
				"code": "A",
			},
		},
		{
			Attributes: map[string]string{
				"id":   "R2",
				"code": "A",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "code",
		RightField: "code",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 4 {
		t.Fatalf("expected 4 joined results, got %d", len(results))
	}
}

func TestApplyJoinPreservesUnmatchedLeftRows(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "L1",
				"code": "A",
			},
		},
		{
			Attributes: map[string]string{
				"id":   "L2",
				"code": "B",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "R1",
				"code": "A",
			},
		},
		{
			Attributes: map[string]string{
				"id":   "R2",
				"code": "C",
			},
		},
	}

	condition := JoinCondition{
		LeftField:  "code",
		RightField: "code",
	}

	results := ApplyJoin(left, right, condition)

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["id"] != "L1" {
		t.Fatalf("expected first result to have left id L1, got %q", results[0].Attributes["id"])
	}

	if results[0].Attributes["right_id"] != "R1" {
		t.Fatalf("expected first result to have right id R1, got %q", results[0].Attributes["right_id"])
	}

	if results[1].Attributes["id"] != "L2" {
		t.Fatalf("expected second result to preserve unmatched left id L2, got %q", results[1].Attributes["id"])
	}
}
