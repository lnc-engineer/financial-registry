package execution

type QueryPagination struct {
	Limit  int
	Offset int
}

func ApplyPagination(results []ExecutionContext, pagination QueryPagination) []ExecutionContext {
	if len(results) == 0 {
		return results
	}

	limit := pagination.Limit
	offset := pagination.Offset

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 {
		limit = len(results)
	}

	if offset >= len(results) {
		return []ExecutionContext{}
	}

	end := offset + limit
	if end > len(results) {
		end = len(results)
	}

	return results[offset:end]
}
