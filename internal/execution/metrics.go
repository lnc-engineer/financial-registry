package execution

import (
	"fmt"
	"sync/atomic"
)

var TotalRequests uint64
var SuccessfulRequests uint64
var FailedRequests uint64
var LastRequestDuration uint64

func RecordRequest() {
	atomic.AddUint64(&TotalRequests, 1)
}

func RecordSuccess(_ ExecutionContext) {
	atomic.AddUint64(&SuccessfulRequests, 1)
}

func RecordFailure(_ ExecutionContext) {
	atomic.AddUint64(&FailedRequests, 1)
}

func RecordDuration(duration uint64) {
	atomic.StoreUint64(&LastRequestDuration, duration)
}

func PrintMetrics() {
	fmt.Println("TOTAL:", TotalRequests)
	fmt.Println("SUCCESS:", SuccessfulRequests)
	fmt.Println("FAILURE:", FailedRequests)
	fmt.Println("LAST DURATION:", LastRequestDuration)
}

type MetricsSnapshot struct {
	TotalRequests         uint64
	Successes             uint64
	Failures              uint64
	LastRequestDuration   uint64
}

func Snapshot() MetricsSnapshot {
	return MetricsSnapshot{
		TotalRequests:       TotalRequests,
		Successes:           SuccessfulRequests,
		Failures:            FailedRequests,
		LastRequestDuration: LastRequestDuration,
	}
}

func ResetMetrics() {
	atomic.StoreUint64(&TotalRequests, 0)
	atomic.StoreUint64(&SuccessfulRequests, 0)
	atomic.StoreUint64(&FailedRequests, 0)
	atomic.StoreUint64(&LastRequestDuration, 0)
}