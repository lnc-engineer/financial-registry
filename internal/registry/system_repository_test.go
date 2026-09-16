package registry

import "testing"

func TestRegistryImplementsSystemRepository(t *testing.T) {
	var repository SystemRepository = NewRegistry()

	if repository == nil {
		t.Fatal("expected repository to be initialized")
	}
}

func TestSystemRepositorySupportsRegistryOperations(t *testing.T) {
	var repository SystemRepository = NewRegistry()

	system := System{
		ID:          "system-001",
		Name:        "Payments Gateway",
		Description: "External payments processing system",
		Status:      SystemStatusActive,
	}

	if err := repository.Register(system); err != nil {
		t.Fatalf("expected registration to succeed, got %v", err)
	}

	registered, exists := repository.Get(system.ID)
	if !exists {
		t.Fatal("expected registered system to be found")
	}

	if registered != system {
		t.Fatalf("expected system %+v, got %+v", system, registered)
	}

	systems := repository.List()

	if len(systems) != 1 {
		t.Fatalf("expected 1 system, got %d", len(systems))
	}

	activeSystems := repository.ListByStatus(SystemStatusActive)

	if len(activeSystems) != 1 {
		t.Fatalf("expected 1 active system, got %d", len(activeSystems))
	}
}
