package execution

import "testing"

func TestFindByAttribute(t *testing.T) {
	spans := []ExecutionContext{
		{
			SpanID: "span-1",
			Attributes: map[string]string{
				"service": "payments",
			},
		},
		{
			SpanID: "span-2",
			Attributes: map[string]string{
				"service": "registry",
			},
		},
		{
			SpanID: "span-3",
			Attributes: map[string]string{
				"service": "payments",
			},
		},
	}

	t.Run("finds matching spans", func(t *testing.T) {
		result := FindByAttribute(spans, "service", "payments")

		if len(result) != 2 {
			t.Fatalf("expected 2 spans, got %d", len(result))
		}
	})

	t.Run("returns no matches", func(t *testing.T) {
		result := FindByAttribute(spans, "service", "unknown")

		if len(result) != 0 {
			t.Fatalf("expected 0 spans, got %d", len(result))
		}
	})

	t.Run("handles missing attribute", func(t *testing.T) {
		result := FindByAttribute(spans, "missing", "value")

		if len(result) != 0 {
			t.Fatalf("expected 0 spans, got %d", len(result))
		}
	})
}
