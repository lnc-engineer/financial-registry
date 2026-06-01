package execution

import "sync/atomic"

var TotalRequests uint64
var SuccessfulRequests uint64
var FailedRequests uint64

func RecordRequest() {
	atomic.AddUint64(&TotalRequests, 1)
}

func RecordSuccess() {
	atomic.AddUint64(&SuccessfulRequests, 1)
}

func RecordFailure() {
	atomic.AddUint64(&FailedRequests, 1)
}
