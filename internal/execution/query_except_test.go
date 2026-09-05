package execution

import "testing"

func TestApplyExceptReturnsLeftOnlyRecords(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "C"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "D"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected A, got %q", results[0].Attributes["id"])
	}
}

func TestApplyExceptExcludesMatchingRecords(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected A, got %q", results[0].Attributes["id"])
	}
}

func TestApplyExceptExcludesRightOnlyRecordsFromOutput(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyExceptRemovesDuplicates(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 distinct results, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected first result A, got %q", results[0].Attributes["id"])
	}

	if results[1].Attributes["id"] != "B" {
		t.Fatalf("expected second result B, got %q", results[1].Attributes["id"])
	}
}

func TestApplyExceptWithEmptyLeft(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyExceptWithEmptyRight(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{}

	results := ApplyExcept(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected first result A, got %q", results[0].Attributes["id"])
	}

	if results[1].Attributes["id"] != "B" {
		t.Fatalf("expected second result B, got %q", results[1].Attributes["id"])
	}
}

func TestApplyExceptWithNoMatches(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}

func TestApplyExceptPreservesLeftContext(t *testing.T) {
	left := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":   "A",
				"name": "Alice",
			},
		},
	}

	right := []ExecutionContext{
		{
			Attributes: map[string]string{
				"id":    "B",
				"other": "value",
			},
		},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected id A, got %q", results[0].Attributes["id"])
	}

	if results[0].Attributes["name"] != "Alice" {
		t.Fatalf("expected name Alice, got %q", results[0].Attributes["name"])
	}

	if _, exists := results[0].Attributes["other"]; exists {
		t.Fatal("expected right-only attribute to be excluded")
	}
}

func TestApplyExceptDeduplicatesRepeatedLeftValues(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A", "source": "first"}},
		{Attributes: map[string]string{"id": "A", "source": "second"}},
		{Attributes: map[string]string{"id": "B", "source": "third"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["source"] != "first" {
		t.Fatalf(
			"expected first matching left context to be preserved, got %s",
			results[0].Attributes["source"],
		)
	}
}

func TestApplyExceptPreservesLeftOrder(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	expected := []string{"C", "B"}

	for i, want := range expected {
		if results[i].Attributes["id"] != want {
			t.Fatalf(
				"expected result %d to be %s, got %s",
				i,
				want,
				results[i].Attributes["id"],
			)
		}
	}
}

func TestApplyExceptRightOrderDoesNotAffectResults(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "C"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected A, got %q", results[0].Attributes["id"])
	}
}

func TestApplyExceptDeduplicatesRepeatedValuesOnBothSides(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A", "source": "first"}},
		{Attributes: map[string]string{"id": "A", "source": "second"}},
		{Attributes: map[string]string{"id": "B", "source": "third"}},
		{Attributes: map[string]string{"id": "B", "source": "fourth"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "C"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["source"] != "first" {
		t.Fatalf("expected first A context, got %s", results[0].Attributes["source"])
	}

	if results[1].Attributes["source"] != "third" {
		t.Fatalf("expected first B context, got %s", results[1].Attributes["source"])
	}
}

func TestApplyExceptWithMissingFieldOnLeft(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"name": "Alice"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	results := ApplyExcept(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["name"] != "Alice" {
		t.Fatalf("expected missing-id left context to be preserved")
	}

	if results[1].Attributes["id"] != "B" {
		t.Fatalf("expected B, got %q", results[1].Attributes["id"])
	}
}

func TestApplyExceptPreservesLeftResultOrder(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-003"}},
		{Attributes: map[string]string{"id": "txn-001"}},
		{Attributes: map[string]string{"id": "txn-004"}},
		{Attributes: map[string]string{"id": "txn-002"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-004"}},
	}

	result := ApplyExcept(left, right, "id")

	if len(result) != 3 {
		t.Fatalf("expected 3 results, got %d", len(result))
	}

	expected := []string{"txn-003", "txn-001", "txn-002"}

	for i, ctx := range result {
		if got := ctx.Attributes["id"]; got != expected[i] {
			t.Fatalf("result[%d]: expected %q, got %q", i, expected[i], got)
		}
	}
}

func TestApplyExceptRemovesAllMatchingRows(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-001"}},
		{Attributes: map[string]string{"id": "txn-002"}},
		{Attributes: map[string]string{"id": "txn-003"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-001"}},
		{Attributes: map[string]string{"id": "txn-002"}},
	}

	result := ApplyExcept(left, right, "id")

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	if got := result[0].Attributes["id"]; got != "txn-003" {
		t.Fatalf("expected txn-003, got %q", got)
	}
}

func TestApplyExceptIdenticalResultsReturnEmpty(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-001"}},
		{Attributes: map[string]string{"id": "txn-002"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-001"}},
		{Attributes: map[string]string{"id": "txn-002"}},
	}

	result := ApplyExcept(left, right, "id")

	if len(result) != 0 {
		t.Fatalf("expected empty result, got %d results", len(result))
	}
}

func TestApplyExceptDisjointResultsPreserveLeftOrder(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-003"}},
		{Attributes: map[string]string{"id": "txn-001"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-004"}},
		{Attributes: map[string]string{"id": "txn-002"}},
	}

	result := ApplyExcept(left, right, "id")

	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}

	expected := []string{"txn-003", "txn-001"}

	for i, ctx := range result {
		if got := ctx.Attributes["id"]; got != expected[i] {
			t.Fatalf("result[%d]: expected %q, got %q", i, expected[i], got)
		}
	}
}

func TestApplyExceptEmptyLeftReturnsEmpty(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-001"}},
	}

	result := ApplyExcept(left, right, "id")

	if len(result) != 0 {
		t.Fatalf("expected empty result, got %d results", len(result))
	}
}

func TestApplyExceptEmptyRightPreservesLeft(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-003"}},
		{Attributes: map[string]string{"id": "txn-001"}},
	}

	right := []ExecutionContext{}

	result := ApplyExcept(left, right, "id")

	if len(result) != len(left) {
		t.Fatalf("expected %d results, got %d", len(left), len(result))
	}

	for i := range left {
		if got := result[i].Attributes["id"]; got != left[i].Attributes["id"] {
			t.Fatalf(
				"result[%d]: expected %q, got %q",
				i,
				left[i].Attributes["id"],
				got,
			)
		}
	}
}

func TestApplyExceptChainedOperations(t *testing.T) {
	first := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-003"}},
		{Attributes: map[string]string{"id": "txn-001"}},
		{Attributes: map[string]string{"id": "txn-004"}},
		{Attributes: map[string]string{"id": "txn-002"}},
	}

	second := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-001"}},
	}

	third := []ExecutionContext{
		{Attributes: map[string]string{"id": "txn-004"}},
	}

	result := ApplyExcept(ApplyExcept(first, second, "id"), third, "id")

	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}

	expected := []string{"txn-003", "txn-002"}

	for i, ctx := range result {
		if got := ctx.Attributes["id"]; got != expected[i] {
			t.Fatalf("result[%d]: expected %q, got %q", i, expected[i], got)
		}
	}
}
