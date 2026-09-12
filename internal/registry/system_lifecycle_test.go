package registry

import "testing"

func TestSystemActivate(t *testing.T) {
	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusInactive,
	}

	err := system.Activate()
	if err != nil {
		t.Fatalf("expected activation to succeed, got %v", err)
	}

	if system.Status != SystemStatusActive {
		t.Fatalf("expected active status, got %q", system.Status)
	}
}

func TestSystemActivateWhenAlreadyActive(t *testing.T) {
	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	err := system.Activate()
	if err == nil {
		t.Fatal("expected activation of active system to fail")
	}

	if system.Status != SystemStatusActive {
		t.Fatalf("expected status to remain active, got %q", system.Status)
	}
}

func TestSystemDeactivate(t *testing.T) {
	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	err := system.Deactivate()
	if err != nil {
		t.Fatalf("expected deactivation to succeed, got %v", err)
	}

	if system.Status != SystemStatusInactive {
		t.Fatalf("expected inactive status, got %q", system.Status)
	}
}

func TestSystemDeactivateWhenAlreadyInactive(t *testing.T) {
	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusInactive,
	}

	err := system.Deactivate()
	if err == nil {
		t.Fatal("expected deactivation of inactive system to fail")
	}

	if system.Status != SystemStatusInactive {
		t.Fatalf("expected status to remain inactive, got %q", system.Status)
	}
}
