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
	results := ApplyLeftAntiJoin(left, right, JoinCondition{LeftField: "account_id", RightField: "account_id"})
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
	results := ApplyLeftAntiJoin(left, right, JoinCondition{LeftField: "trace_id", RightField: "trace_id"})
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}
