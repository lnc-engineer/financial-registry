package execution

import "testing"

func TestMatchesHavingCountEquals(t *testing.T) {

	aggregation := GroupAggregation{
		Key:   "success",
		Count: 5,
	}

	condition := HavingCondition{
		Field:    "count",
		Operator: OperatorEquals,
		Value:    5,
	}

	if !MatchesHaving(aggregation, condition) {
		t.Fatal("expected aggregation to match having condition")
	}
}

func TestMatchesHavingCountGreaterThan(t *testing.T) {

	aggregation := GroupAggregation{
		Key:   "success",
		Count: 10,
	}

	condition := HavingCondition{
		Field:    "count",
		Operator: OperatorGreaterThan,
		Value:    5,
	}

	if !MatchesHaving(aggregation, condition) {
		t.Fatal("expected aggregation count to be greater than value")
	}
}

func TestApplyHavingFiltersAggregations(t *testing.T) {

	aggregations := []GroupAggregation{
		{
			Key:   "success",
			Count: 10,
		},
		{
			Key:   "failure",
			Count: 2,
		},
	}

	condition := HavingCondition{
		Field:    "count",
		Operator: OperatorGreaterThan,
		Value:    5,
	}

	filtered := ApplyHaving(aggregations, condition)

	if len(filtered) != 1 {
		t.Fatalf("expected 1 aggregation, got %d", len(filtered))
	}

	if filtered[0].Key != "success" {
		t.Fatalf("expected success aggregation, got %s", filtered[0].Key)
	}
}

func TestMatchesHavingUnknownField(t *testing.T) {

	aggregation := GroupAggregation{
		Key:   "success",
		Count: 5,
	}

	condition := HavingCondition{
		Field:    "unknown",
		Operator: OperatorEquals,
		Value:    5,
	}

	if MatchesHaving(aggregation, condition) {
		t.Fatal("expected unknown field to return false")
	}
}

func TestApplyHavingEmptyAggregations(t *testing.T) {

	filtered := ApplyHaving(
		nil,
		HavingCondition{
			Field:    "count",
			Operator: OperatorGreaterThan,
			Value:    5,
		},
	)

	if len(filtered) != 0 {
		t.Fatalf("expected empty result, got %d", len(filtered))
	}
}
