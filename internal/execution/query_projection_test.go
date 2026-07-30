package execution

import "testing"

func TestApplyProjectionSingleField(t *testing.T) {
	ctx := ExecutionContext{
		TraceID: "trace-001",
	}

	projection := QueryProjection{
		Fields: []string{"trace_id"},
	}

	result := ApplyProjection(ctx, projection)

	if result["trace_id"] != "trace-001" {
		t.Fatalf("expected trace_id to be trace-001")
	}
}

func TestApplyProjectionMultipleFields(t *testing.T) {
	ctx := ExecutionContext{
		TraceID: "trace-001",
		SpanID:  "span-001",
		Status:  "success",
	}

	projection := QueryProjection{
		Fields: []string{
			"trace_id",
			"span_id",
			"status",
		},
	}

	result := ApplyProjection(ctx, projection)

	if len(result) != 3 {
		t.Fatalf("expected 3 projected fields, got %d", len(result))
	}

	if result["trace_id"] != "trace-001" {
		t.Fatal("unexpected trace_id")
	}

	if result["span_id"] != "span-001" {
		t.Fatal("unexpected span_id")
	}

	if result["status"] != "success" {
		t.Fatal("unexpected status")
	}
}

func TestApplyProjectionAttributes(t *testing.T) {
	ctx := ExecutionContext{
		Attributes: map[string]string{
			"service": "payments",
		},
	}

	projection := QueryProjection{
		Fields: []string{"service"},
	}

	result := ApplyProjection(ctx, projection)

	if result["service"] != "payments" {
		t.Fatal("expected attribute to be projected")
	}
}

func TestApplyProjectionIgnoresUnknownFields(t *testing.T) {
	ctx := ExecutionContext{}

	projection := QueryProjection{
		Fields: []string{"does_not_exist"},
	}

	result := ApplyProjection(ctx, projection)

	if len(result) != 0 {
		t.Fatalf("expected empty projection, got %d fields", len(result))
	}
}

func TestApplyProjectionEmptyProjection(t *testing.T) {
	ctx := ExecutionContext{}

	result := ApplyProjection(ctx, QueryProjection{})

	if len(result) != 0 {
		t.Fatalf("expected empty result")
	}
}

func TestApplyProjectionToResults(t *testing.T) {
	results := []ExecutionContext{
		{
			TraceID: "trace-001",
			Status:  "success",
		},
		{
			TraceID: "trace-002",
			Status:  "failure",
		},
	}

	projection := QueryProjection{
		Fields: []string{
			"trace_id",
			"status",
		},
	}

	projected := ApplyProjectionToResults(results, projection)

	if len(projected) != 2 {
		t.Fatalf("expected 2 projected results, got %d", len(projected))
	}

	if projected[0]["trace_id"] != "trace-001" {
		t.Fatal("unexpected first trace_id")
	}

	if projected[1]["status"] != "failure" {
		t.Fatal("unexpected second status")
	}
}
