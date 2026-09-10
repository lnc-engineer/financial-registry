package registry

import "testing"

func TestSystemStatusValues(t *testing.T) {
	if SystemStatusActive != "active" {
		t.Fatalf("expected active status, got %q", SystemStatusActive)
	}

	if SystemStatusInactive != "inactive" {
		t.Fatalf("expected inactive status, got %q", SystemStatusInactive)
	}
}

func TestSystemStoresRegistryDetails(t *testing.T) {
	system := System{
		ID:          "system-001",
		Name:        "Payments Gateway",
		Description: "External payments processing system",
		Status:      SystemStatusActive,
	}

	if system.ID != "system-001" {
		t.Fatalf("expected ID system-001, got %q", system.ID)
	}

	if system.Name != "Payments Gateway" {
		t.Fatalf("expected name Payments Gateway, got %q", system.Name)
	}

	if system.Description != "External payments processing system" {
		t.Fatalf("unexpected description: %q", system.Description)
	}

	if system.Status != SystemStatusActive {
		t.Fatalf("expected active status, got %q", system.Status)
	}
}
