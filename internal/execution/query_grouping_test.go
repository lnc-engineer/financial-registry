package execution

import "testing"

func TestGroupByFieldStatus(t *testing.T) {
	results := []ExecutionContext{
		{Status: "success"},
		{Status: "success"},
		{Status: "failure"},
	}

	groups := GroupByField(results, "status")

	if len(groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(groups))
	}

	counts := make(map[string]int)
	for _, group := range groups {
		counts[group.Key] = len(group.Results)
	}

	if counts["success"] != 2 {
		t.Errorf("expected 2 success results, got %d", counts["success"])
	}

	if counts["failure"] != 1 {
		t.Errorf("expected 1 failure result, got %d", counts["failure"])
	}
}

func TestGroupByFieldEmpty(t *testing.T) {
	groups := GroupByField(nil, "status")

	if len(groups) != 0 {
		t.Fatalf("expected no groups, got %d", len(groups))
	}
}

func TestGroupByFieldUnknownField(t *testing.T) {
	results := []ExecutionContext{
		{},
		{},
	}

	groups := GroupByField(results, "unknown")

	if len(groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(groups))
	}

	if groups[0].Key != "" {
		t.Errorf("expected empty key, got %q", groups[0].Key)
	}
}
