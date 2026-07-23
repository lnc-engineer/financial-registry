package execution

import "testing"

func TestMatchValueEquals(t *testing.T) {
	if !MatchValue("success", "success", OperatorEquals) {
		t.Fatal("expected values to match")
	}
}

func TestMatchValueContains(t *testing.T) {
	if !MatchValue("processing complete", "process", OperatorContains) {
		t.Fatal("expected substring match")
	}
}

func TestMatchValueStartsWith(t *testing.T) {
	if !MatchValue("trace-123", "trace", OperatorStartsWith) {
		t.Fatal("expected prefix match")
	}
}

func TestMatchValueNoMatch(t *testing.T) {
	if MatchValue("failure", "success", OperatorEquals) {
		t.Fatal("expected values not to match")
	}
}
