package execution

import "testing"

func TestAggregateByStatus(t *testing.T) {
	records := []ExecutionContext{
		{
			Status: "success",
		},
		{
			Status: "success",
		},
		{
			Status: "failure",
		},
	}

	aggregation := QueryAggregation{
		Type:  AggregationCount,
		Field: "status",
	}

	result := Aggregate(records, aggregation)

	if result["success"] != 2 {
		t.Fatalf("expected success count 2, got %d", result["success"])
	}

	if result["failure"] != 1 {
		t.Fatalf("expected failure count 1, got %d", result["failure"])
	}
}

func TestAggregateBySpanName(t *testing.T) {
	records := []ExecutionContext{
		{
			SpanName: "ProcessRecord",
		},
		{
			SpanName: "ProcessRecord",
		},
		{
			SpanName: "ValidateInput",
		},
	}

	aggregation := QueryAggregation{
		Type:  AggregationCount,
		Field: "span_name",
	}

	result := Aggregate(records, aggregation)

	if result["ProcessRecord"] != 2 {
		t.Fatalf("expected ProcessRecord count 2, got %d", result["ProcessRecord"])
	}

	if result["ValidateInput"] != 1 {
		t.Fatalf("expected ValidateInput count 1, got %d", result["ValidateInput"])
	}
}

func TestAggregateByAttribute(t *testing.T) {
	records := []ExecutionContext{
		{
			Attributes: map[string]string{
				"service": "payments",
			},
		},
		{
			Attributes: map[string]string{
				"service": "payments",
			},
		},
		{
			Attributes: map[string]string{
				"service": "registry",
			},
		},
	}

	aggregation := QueryAggregation{
		Type:  AggregationCount,
		Field: "service",
	}

	result := Aggregate(records, aggregation)

	if result["payments"] != 2 {
		t.Fatalf("expected payments count 2, got %d", result["payments"])
	}

	if result["registry"] != 1 {
		t.Fatalf("expected registry count 1, got %d", result["registry"])
	}
}

func TestAggregateEmptyRecords(t *testing.T) {
	records := []ExecutionContext{}

	aggregation := QueryAggregation{
		Type:  AggregationCount,
		Field: "status",
	}

	result := Aggregate(records, aggregation)

	if len(result) != 0 {
		t.Fatalf("expected empty result, got %v", result)
	}
}

func TestAggregateUnknownField(t *testing.T) {
	records := []ExecutionContext{
		{
			Status: "success",
		},
		{
			Status: "failure",
		},
	}

	aggregation := QueryAggregation{
		Type:  AggregationCount,
		Field: "unknown_field",
	}

	result := Aggregate(records, aggregation)

	if result[""] != 2 {
		t.Fatalf("expected unknown field count 2, got %d", result[""])
	}
}
