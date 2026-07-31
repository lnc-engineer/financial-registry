package execution

type AggregationType string

const (
	AggregationCount AggregationType = "count"
)

type QueryAggregation struct {
	Type  AggregationType
	Field string
}

func Aggregate(
	records []ExecutionContext,
	aggregation QueryAggregation,
) map[string]int {

	result := make(map[string]int)

	for _, record := range records {
		value := resolveField(record, aggregation.Field)
		result[value]++

	}

	return result
}
