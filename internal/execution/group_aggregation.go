package execution

type GroupAggregation struct {
	Key   string
	Count int
}

func AggregateGroups(groups []QueryGroup) []GroupAggregation {
	aggregations := make([]GroupAggregation, 0, len(groups))

	for _, group := range groups {
		aggregations = append(aggregations, GroupAggregation{
			Key:   group.Key,
			Count: len(group.Results),
		})
	}

	return aggregations
}
