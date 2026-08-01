package execution

import "testing"

func TestAggregateGroups(t *testing.T) {
	groups := []QueryGroup{
		{
			Key: "success",
			Results: []ExecutionContext{
				{},
				{},
			},
		},
		{
			Key: "failure",
			Results: []ExecutionContext{
				{},
			},
		},
	}

	aggregations := AggregateGroups(groups)

	if len(aggregations) != 2 {
		t.Fatalf("expected 2 aggregations, got %d", len(aggregations))
	}

	counts := make(map[string]int)

	for _, aggregation := range aggregations {
		counts[aggregation.Key] = aggregation.Count
	}

	if counts["success"] != 2 {
		t.Errorf("expected success count of 2, got %d", counts["success"])
	}

	if counts["failure"] != 1 {
		t.Errorf("expected failure count of 1, got %d", counts["failure"])
	}
}

func TestAggregateGroupsEmpty(t *testing.T) {
	aggregations := AggregateGroups(nil)

	if len(aggregations) != 0 {
		t.Fatalf("expected no aggregations, got %d", len(aggregations))
	}
}
