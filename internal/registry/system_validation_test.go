package registry

import "testing"

func TestSystemValidateAcceptsValidSystem(t *testing.T) {
	system := System{
		ID:          "system-001",
		Name:        "Payments System",
		Description: "Core payments platform",
		Status:      SystemStatusActive,
	}

	if err := system.Validate(); err != nil {
		t.Fatalf("expected valid system, got error: %v", err)
	}
}

func TestSystemValidateRequiresID(t *testing.T) {
	system := System{
		Name:   "Payments System",
		Status: SystemStatusActive,
	}

	if err := system.Validate(); err == nil {
		t.Fatal("expected validation error for missing ID")
	}
}

func TestSystemValidateRequiresName(t *testing.T) {
	system := System{
		ID:     "system-001",
		Status: SystemStatusActive,
	}

	if err := system.Validate(); err == nil {
		t.Fatal("expected validation error for missing name")
	}
}

func TestSystemValidateRejectsInvalidStatus(t *testing.T) {
	system := System{
		ID:     "system-001",
		Name:   "Payments System",
		Status: SystemStatus("unknown"),
	}

	if err := system.Validate(); err == nil {
		t.Fatal("expected validation error for invalid status")
	}
}

func TestSystemValidateAllowsOptionalDescription(t *testing.T) {
	system := System{
		ID:     "system-001",
		Name:   "Payments System",
		Status: SystemStatusInactive,
	}

	if err := system.Validate(); err != nil {
		t.Fatalf("expected valid system without description, got error: %v", err)
	}
}
