package registry

// SystemStatus represents the lifecycle state of a registered system.
type SystemStatus string

const (
	SystemStatusActive   SystemStatus = "active"
	SystemStatusInactive SystemStatus = "inactive"
)

// System represents a financial system registered with the platform.
type System struct {
	ID          string
	Name        string
	Description string
	Status      SystemStatus
}
