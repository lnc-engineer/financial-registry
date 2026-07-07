package execution

import "time"

type TraceSummary struct {
	TraceID       string
	TotalSpans    int
	SuccessCount  int
	FailureCount  int
	MaxDepth      int
	TotalDuration time.Duration
}

func BuildTraceSummary(roots []*TraceNode) TraceSummary {
	summary := TraceSummary{}

	if len(roots) == 0 {
		return summary
	}

	buildSummary(roots[0], 1, &summary)

	for _, root := range roots {
		if summary.TraceID == "" {
			summary.TraceID = root.Context.TraceID
		}

		if root.Context.Duration() > summary.TotalDuration {
			summary.TotalDuration = root.Context.Duration()
		}
	}

	return summary
}

func buildSummary(node *TraceNode, depth int, summary *TraceSummary) {
	if node == nil {
		return
	}

	summary.TotalSpans++

	if depth > summary.MaxDepth {
		summary.MaxDepth = depth
	}

	switch node.Context.Status {
	case "SUCCESS":
		summary.SuccessCount++
	case "FAILURE":
		summary.FailureCount++
	}

	for _, child := range node.Children {
		buildSummary(child, depth+1, summary)
	}
}
