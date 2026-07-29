package execution

import "testing"

func samplePaginationResults(count int) []ExecutionContext {
	results := make([]ExecutionContext, count)

	for i := range results {
		results[i] = ExecutionContext{
			TraceID: string(rune('A' + i)),
		}
	}

	return results
}

func TestApplyPaginationFirstPage(t *testing.T) {
	results := samplePaginationResults(5)

	page := ApplyPagination(results, QueryPagination{
		Limit: 2,
	})

	if len(page) != 2 {
		t.Fatalf("expected 2 results, got %d", len(page))
	}

	if page[0].TraceID != "A" || page[1].TraceID != "B" {
		t.Fatalf("unexpected first page")
	}
}

func TestApplyPaginationOffset(t *testing.T) {
	results := samplePaginationResults(5)

	page := ApplyPagination(results, QueryPagination{
		Limit:  2,
		Offset: 2,
	})

	if len(page) != 2 {
		t.Fatalf("expected 2 results, got %d", len(page))
	}

	if page[0].TraceID != "C" || page[1].TraceID != "D" {
		t.Fatalf("unexpected offset page")
	}
}

func TestApplyPaginationOffsetBeyondResults(t *testing.T) {
	results := samplePaginationResults(3)

	page := ApplyPagination(results, QueryPagination{
		Limit:  2,
		Offset: 10,
	})

	if len(page) != 0 {
		t.Fatalf("expected empty result")
	}
}

func TestApplyPaginationUnlimited(t *testing.T) {
	results := samplePaginationResults(4)

	page := ApplyPagination(results, QueryPagination{})

	if len(page) != 4 {
		t.Fatalf("expected all results")
	}
}

func TestApplyPaginationNegativeOffset(t *testing.T) {
	results := samplePaginationResults(4)

	page := ApplyPagination(results, QueryPagination{
		Limit:  2,
		Offset: -5,
	})

	if len(page) != 2 {
		t.Fatalf("expected 2 results")
	}

	if page[0].TraceID != "A" {
		t.Fatalf("unexpected first result")
	}
}
