package execution

import "sort"

type SortField string

const (
	SortByStartTime SortField = "start_time"
	SortByEndTime   SortField = "end_time"
	SortBySpanName  SortField = "span_name"
)

type SortOrder string

const (
	SortAscending  SortOrder = "asc"
	SortDescending SortOrder = "desc"
)

type QuerySort struct {
	Field SortField
	Order SortOrder
}

func SortExecutionContexts(contexts []ExecutionContext, querySort QuerySort) {
	sort.Slice(contexts, func(i, j int) bool {
		var less bool

		switch querySort.Field {
		case SortByStartTime:
			less = contexts[i].StartTime.Before(contexts[j].StartTime)

		case SortByEndTime:
			less = contexts[i].EndTime.Before(contexts[j].EndTime)

		case SortBySpanName:
			less = contexts[i].SpanName < contexts[j].SpanName

		default:
			less = true
		}

		if querySort.Order == SortDescending {
			return !less
		}

		return less
	})
}
