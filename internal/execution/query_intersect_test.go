package execution

import "testing"

func TestApplyIntersectReturnsCommonRecords(t *testing.T) {
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

	results := ApplyIntersect(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["id"] != "B" {
		t.Fatalf("expected first result B, got %q", results[0].Attributes["id"])
	}

	if results[1].Attributes["id"] != "C" {
		t.Fatalf("expected second result C, got %q", results[1].Attributes["id"])
	}
}

func TestApplyIntersectExcludesLeftOnlyRecords(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "B" {
		t.Fatalf("expected B, got %q", results[0].Attributes["id"])
	}
}

func TestApplyIntersectExcludesRightOnlyRecords(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected A, got %q", results[0].Attributes["id"])
	}
}

func TestApplyIntersectRemovesDuplicates(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

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

func TestApplyIntersectWithEmptyLeft(t *testing.T) {
	left := []ExecutionContext{}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyIntersectWithEmptyRight(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	right := []ExecutionContext{}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyIntersectWithNoMatches(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}

func TestApplyIntersectPreservesLeftContext(t *testing.T) {
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
				"id":    "A",
				"other": "value",
			},
		},
	}

	results := ApplyIntersect(left, right, "id")

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

func TestApplyIntersectDeduplicatesRepeatedLeftValues(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A", "source": "first"}},
		{Attributes: map[string]string{"id": "A", "source": "second"}},
		{Attributes: map[string]string{"id": "B", "source": "third"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected first result id A, got %s", results[0].Attributes["id"])
	}

	if results[0].Attributes["source"] != "first" {
		t.Fatalf("expected first matching left context to be preserved, got %s", results[0].Attributes["source"])
	}
}

func TestApplyIntersectDeduplicatesRepeatedRightValues(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A", "source": "first"}},
		{Attributes: map[string]string{"id": "A", "source": "second"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Attributes["id"] != "A" {
		t.Fatalf("expected result id A, got %s", results[0].Attributes["id"])
	}
}

func TestApplyIntersectDeduplicatesRepeatedValuesOnBothSides(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A", "source": "first"}},
		{Attributes: map[string]string{"id": "A", "source": "second"}},
		{Attributes: map[string]string{"id": "B", "source": "third"}},
		{Attributes: map[string]string{"id": "B", "source": "fourth"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

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

func TestApplyIntersectPreservesLeftOrder(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "A"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	expected := []string{"C", "A", "B"}

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

func TestApplyIntersectRightOrderDoesNotAffectResults(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "C"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	expected := []string{"A", "B", "C"}

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

func TestApplyIntersectPreservesFirstMatchingLeftOrderWithDuplicates(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "B", "source": "first-b"}},
		{Attributes: map[string]string{"id": "A", "source": "first-a"}},
		{Attributes: map[string]string{"id": "B", "source": "second-b"}},
		{Attributes: map[string]string{"id": "C", "source": "first-c"}},
		{Attributes: map[string]string{"id": "A", "source": "second-a"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	expectedIDs := []string{"B", "A", "C"}
	expectedSources := []string{"first-b", "first-a", "first-c"}

	for i := range expectedIDs {
		if results[i].Attributes["id"] != expectedIDs[i] {
			t.Fatalf(
				"expected result %d id %s, got %s",
				i,
				expectedIDs[i],
				results[i].Attributes["id"],
			)
		}

		if results[i].Attributes["source"] != expectedSources[i] {
			t.Fatalf(
				"expected result %d source %s, got %s",
				i,
				expectedSources[i],
				results[i].Attributes["source"],
			)
		}
	}
}

func TestApplyIntersectPreservesOrderWhenUnmatchedRowsAreInterleaved(t *testing.T) {
	left := []ExecutionContext{
		{Attributes: map[string]string{"id": "X"}},
		{Attributes: map[string]string{"id": "C"}},
		{Attributes: map[string]string{"id": "Y"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "Z"}},
		{Attributes: map[string]string{"id": "B"}},
	}

	right := []ExecutionContext{
		{Attributes: map[string]string{"id": "B"}},
		{Attributes: map[string]string{"id": "A"}},
		{Attributes: map[string]string{"id": "C"}},
	}

	results := ApplyIntersect(left, right, "id")

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	expected := []string{"C", "A", "B"}

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
