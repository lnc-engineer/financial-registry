package execution

import "time"

type TraceStats struct {
	TotalSpans       int
	RootSpans        int
	ChildSpans       int
	SuccessSpans     int
	FailureSpans     int
	UnknownSpans     int
	AverageDuration  time.Duration
	LongestDuration  time.Duration
	ShortestDuration time.Duration
}

func CalculateTraceStats(spans []ExecutionContext) TraceStats {
	stats := TraceStats{}

	if len(spans) == 0 {
		return stats
	}

	var totalDuration time.Duration

	for _, span := range spans {
		stats.TotalSpans++

		if span.ParentSpanID == "" {
			stats.RootSpans++
		} else {
			stats.ChildSpans++
		}

		switch span.Status {
		case "success":
			stats.SuccessSpans++
		case "failure":
			stats.FailureSpans++
		default:
			stats.UnknownSpans++
		}

		duration := span.EndTime.Sub(span.StartTime)
		totalDuration += duration

		if stats.ShortestDuration == 0 || duration < stats.ShortestDuration {
			stats.ShortestDuration = duration
		}

		if duration > stats.LongestDuration {
			stats.LongestDuration = duration
		}
	}

	stats.AverageDuration = totalDuration / time.Duration(stats.TotalSpans)

	return stats
}
