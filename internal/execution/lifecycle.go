package execution

import "time"

type LifecycleEvent struct {
	Name      string
	Timestamp time.Time
}

func NewLifecycleEvent(name string) LifecycleEvent {
	return LifecycleEvent{
		Name:      name,
		Timestamp: time.Now(),
	}
}

func (ctx *ExecutionContext) AddLifecycleEvent(name string) {
	ctx.Lifecycle = append(
		ctx.Lifecycle,
		NewLifecycleEvent(name),
	)
}
