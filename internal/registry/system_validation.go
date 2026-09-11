package registry

import "fmt"

// Validate checks whether the system contains valid domain values.
func (s System) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("system ID is required")
	}

	if s.Name == "" {
		return fmt.Errorf("system name is required")
	}

	switch s.Status {
	case SystemStatusActive, SystemStatusInactive:
		return nil
	default:
		return fmt.Errorf("invalid system status: %q", s.Status)
	}
}
