package execution

import (
	"fmt"
	"sync/atomic"
)

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

func PrintMetrics() {
	fmt.Println("TOTAL:", atomic.LoadUint64(&TotalRequests))
	fmt.Println("SUCCESS:", atomic.LoadUint64(&SuccessfulRequests))
	fmt.Println("FAILURE:", atomic.LoadUint64(&FailedRequests))
}
