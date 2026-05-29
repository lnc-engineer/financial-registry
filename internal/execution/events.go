package execution

import "fmt"

func LogEvent(
	ctx ExecutionContext,
	stage string,
) {
	fmt.Printf(
		"[EXECUTION_EVENT] request_id=%s stage=%s\n",
		ctx.RequestID,
		stage,
	)
}