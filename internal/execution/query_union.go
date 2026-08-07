package execution

type UnionQuery struct {
	Queries []CompositeFilter
}

func ApplyUnion(query UnionQuery, spans []ExecutionContext) []ExecutionContext {
	var results []ExecutionContext

	for _, filter := range query.Queries {
		filtered := filter.Apply(spans)
		results = append(results, filtered...)
	}

	return ApplyDistinct(results, []string{"trace_id"})
}
