package execution

type MetricsSnapshot struct {
	TotalRequests         uint64 `json:"total_requests"`
	Successes             uint64 `json:"successes"`
	Failures              uint64 `json:"failures"`
	LastRequestDurationMs uint64 `json:"last_request_duration_ms"`
}
