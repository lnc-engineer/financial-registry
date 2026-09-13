package registry

import "testing"

func TestNewRegistryCreatesEmptyRegistry(t *testing.T) {
	registry := NewRegistry()

	if registry == nil {
		t.Fatal("expected registry to be created")
	}

	if registry.systems == nil {
		t.Fatal("expected registry systems map to be initialized")
	}

	if len(registry.systems) != 0 {
		t.Fatalf("expected empty registry, got %d systems", len(registry.systems))
	}
}

func TestRegistryRegisterStoresValidSystem(t *testing.T) {
	registry := NewRegistry()

	system := System{
		ID:          "system-001",
		Name:        "Payments Gateway",
		Description: "External payments processing system",
		Status:      SystemStatusActive,
	}

	if err := registry.Register(system); err != nil {
		t.Fatalf("expected registration to succeed, got %v", err)
	}

	registered, exists := registry.systems["system-001"]
	if !exists {
		t.Fatal("expected system to be registered")
	}

	if registered != system {
		t.Fatalf("expected registered system %+v, got %+v", system, registered)
	}
}

func TestRegistryRegisterRejectsDuplicateID(t *testing.T) {
	registry := NewRegistry()

	first := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	second := System{
		ID:     "system-001",
		Name:   "Market Data Platform",
		Status: SystemStatusActive,
	}

	if err := registry.Register(first); err != nil {
		t.Fatalf("expected first registration to succeed, got %v", err)
	}

	if err := registry.Register(second); err == nil {
		t.Fatal("expected duplicate ID registration to fail")
	}

	if len(registry.systems) != 1 {
		t.Fatalf("expected registry to contain 1 system, got %d", len(registry.systems))
	}

	registered := registry.systems["system-001"]

	if registered.Name != "Payments Gateway" {
		t.Fatalf("expected original system to remain registered, got %q", registered.Name)
	}
}

func TestRegistryRegisterRejectsInvalidSystem(t *testing.T) {
	registry := NewRegistry()

	system := System{
		ID:     "",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	if err := registry.Register(system); err == nil {
		t.Fatal("expected invalid system registration to fail")
	}

	if len(registry.systems) != 0 {
		t.Fatalf("expected registry to remain empty, got %d systems", len(registry.systems))
	}
}

func TestRegistryGetReturnsRegisteredSystem(t *testing.T) {
	registry := NewRegistry()

	system := System{
		ID:          "system-001",
		Name:        "Payments Gateway",
		Description: "External payments processing system",
		Status:      SystemStatusActive,
	}

	if err := registry.Register(system); err != nil {
		t.Fatalf("expected registration to succeed, got %v", err)
	}

	registered, exists := registry.Get("system-001")
	if !exists {
		t.Fatal("expected registered system to be found")
	}

	if registered != system {
		t.Fatalf("expected system %+v, got %+v", system, registered)
	}
}

func TestRegistryGetReturnsFalseForUnknownID(t *testing.T) {
	registry := NewRegistry()

	system, exists := registry.Get("system-999")

	if exists {
		t.Fatal("expected unknown system ID to return false")
	}

	if system != (System{}) {
		t.Fatalf("expected zero-value system, got %+v", system)
	}
}
