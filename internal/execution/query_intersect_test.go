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
