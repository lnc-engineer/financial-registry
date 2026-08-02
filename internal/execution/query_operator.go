package execution

import "strings"

type QueryOperator string

const (
	OperatorEquals             QueryOperator = "equals"
	OperatorNotEquals          QueryOperator = "not_equals"
	OperatorContains           QueryOperator = "contains"
	OperatorStartsWith         QueryOperator = "starts_with"
	OperatorGreaterThan        QueryOperator = "greater_than"
	OperatorGreaterThanOrEqual QueryOperator = "greater_than_or_equal"
	OperatorLessThan           QueryOperator = "less_than"
	OperatorLessThanOrEqual    QueryOperator = "less_than_or_equal"
)

func MatchValue(value, expected string, operator QueryOperator) bool {
	switch operator {
	case OperatorEquals:
		return value == expected

	case OperatorNotEquals:
		return value != expected

	case OperatorContains:
		return strings.Contains(value, expected)

	case OperatorStartsWith:
		return strings.HasPrefix(value, expected)

	default:
		return false
	}
}
