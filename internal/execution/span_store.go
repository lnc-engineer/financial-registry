package execution

import "sync"

var (
    spansMu sync.Mutex
    spans   []ExecutionContext
)

func RecordSpan(span ExecutionContext) {
    spansMu.Lock()
    defer spansMu.Unlock()

    spans = append(spans, span)
}

func GetSpans() []ExecutionContext {
    spansMu.Lock()
    defer spansMu.Unlock()

    out := make([]ExecutionContext, len(spans))
    copy(out, spans)
    return out
}

func ResetSpans() {
    spansMu.Lock()
    defer spansMu.Unlock()

    spans = nil
}