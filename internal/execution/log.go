package execution

import (
	"encoding/json"
	"fmt"
)

func LogEvent(ctx ExecutionContext, event string) {
	e := NewEvent(ctx, event)

	RecordEvent(e)

	b, _ := json.Marshal(e)
	fmt.Println(string(b))
}
