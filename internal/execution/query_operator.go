package execution

import "strings"

type QueryOperator string

const (
	OperatorEquals     QueryOperator = "equals"
	OperatorContains   QueryOperator = "contains"
	OperatorStartsWith QueryOperator = "starts_with"
)

func MatchValue(value, expected string, operator QueryOperator) bool {
	switch operator {
	case OperatorEquals:
		return value == expected
	case OperatorContains:
		return strings.Contains(value, expected)
	case OperatorStartsWith:
		return strings.HasPrefix(value, expected)
	default:
		return false
	}
}
