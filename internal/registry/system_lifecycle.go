package registry

import "fmt"

// Activate changes an inactive system to the active state.
func (s *System) Activate() error {
	if s.Status == SystemStatusActive {
		return fmt.Errorf("system is already active")
	}

	s.Status = SystemStatusActive
	return nil
}

// Deactivate changes an active system to the inactive state.
func (s *System) Deactivate() error {
	if s.Status == SystemStatusInactive {
		return fmt.Errorf("system is already inactive")
	}

	s.Status = SystemStatusInactive
	return nil
}
