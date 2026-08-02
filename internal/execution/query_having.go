package execution

type HavingCondition struct {
	Field    string
	Operator QueryOperator
	Value    int
}

func MatchesHaving(
	aggregation GroupAggregation,
	condition HavingCondition,
) bool {

	var actual int

	switch condition.Field {
	case "count":
		actual = aggregation.Count
	default:
		return false
	}

	switch condition.Operator {

	case OperatorEquals:
		return actual == condition.Value

	case OperatorGreaterThan:
		return actual > condition.Value

	case OperatorGreaterThanOrEqual:
		return actual >= condition.Value

	case OperatorLessThan:
		return actual < condition.Value

	case OperatorLessThanOrEqual:
		return actual <= condition.Value

	default:
		return false
	}

}

func ApplyHaving(
	aggregations []GroupAggregation,
	condition HavingCondition,
) []GroupAggregation {

	filtered := make([]GroupAggregation, 0)

	for _, aggregation := range aggregations {
		if MatchesHaving(aggregation, condition) {
			filtered = append(filtered, aggregation)
		}
	}

	return filtered
}
