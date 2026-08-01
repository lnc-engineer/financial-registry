package execution

type QueryGroup struct {
	Key     string
	Results []ExecutionContext
}

func GroupByField(results []ExecutionContext, field string) []QueryGroup {
	groups := make(map[string][]ExecutionContext)

	for _, result := range results {
		key := resolveField(result, field)
		groups[key] = append(groups[key], result)
	}

	queryGroups := make([]QueryGroup, 0, len(groups))

	for key, results := range groups {
		queryGroups = append(queryGroups, QueryGroup{
			Key:     key,
			Results: results,
		})
	}

	return queryGroups
}
