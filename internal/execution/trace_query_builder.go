package execution

type TraceQueryBuilder struct {
	filter CompositeFilter
}

func NewTraceQueryBuilder() *TraceQueryBuilder {
	return &TraceQueryBuilder{}
}

func (b *TraceQueryBuilder) WithTraceID(traceID string) *TraceQueryBuilder {
	b.filter.TraceID = traceID
	return b
}

func (b *TraceQueryBuilder) WithStatus(status string) *TraceQueryBuilder {
	b.filter.Status = status
	return b
}

func (b *TraceQueryBuilder) Build() CompositeFilter {
	return b.filter
}
