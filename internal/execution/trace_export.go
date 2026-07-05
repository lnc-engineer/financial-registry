package execution

import (
	"encoding/json"
	"fmt"
)

func PrintTraceJSON(roots []*TraceNode) {

	var exports []TraceExport

	for _, r := range roots {
		exports = append(exports, r.Export())
	}

	data, err := json.MarshalIndent(exports, "", "  ")
	if err != nil {
		fmt.Println("JSON export error:", err)
		return
	}

	fmt.Println("\nTRACE JSON")
	fmt.Println(string(data))
}

type TraceExport struct {
	SpanName     string        `json:"span"`
	TraceID      string        `json:"trace_id"`
	SpanID       string        `json:"span_id"`
	ParentSpanID string        `json:"parent_span_id,omitempty"`
	Status       string        `json:"status"`
	Duration     string        `json:"duration"`
	Children     []TraceExport `json:"children,omitempty"`
}

func (t *TraceNode) Export() TraceExport {

	export := TraceExport{
		SpanName:     t.Context.SpanName,
		TraceID:      t.Context.TraceID,
		SpanID:       t.Context.SpanID,
		ParentSpanID: t.Context.ParentSpanID,
		Status:       deriveStatus(t.Context),
		Duration:     t.Context.Duration().String(),
	}

	for _, child := range t.Children {
		export.Children = append(export.Children, child.Export())
	}

	return export
}

func deriveStatus(ctx ExecutionContext) string {
	if ctx.Status != "" {
		return ctx.Status
	}

	return "unknown"
}
