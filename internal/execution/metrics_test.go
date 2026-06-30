package execution

import "testing"

func TestResetMetrics(t *testing.T) {
	RecordRequest()
	RecordSuccess(ExecutionContext{})
	RecordFailure(ExecutionContext{})
	RecordDuration(123)

	ResetMetrics()

	s := Snapshot()

	if s.TotalRequests != 0 {
		t.Errorf("expected TotalRequests = 0, got %d", s.TotalRequests)
	}
	if s.Successes != 0 {
		t.Errorf("expected Successes = 0, got %d", s.Successes)
	}
	if s.Failures != 0 {
		t.Errorf("expected Failures = 0, got %d", s.Failures)
	}
	if s.LastRequestDuration != 0 {
		t.Errorf("expected LastRequestDuration = 0, got %d", s.LastRequestDuration)
	}
}