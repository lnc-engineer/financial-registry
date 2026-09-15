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

func TestRegistryListReturnsEmptyCollectionForEmptyRegistry(t *testing.T) {
	registry := NewRegistry()

	systems := registry.List()

	if systems == nil {
		t.Fatal("expected List to return an empty collection, got nil")
	}

	if len(systems) != 0 {
		t.Fatalf("expected empty collection, got %d systems", len(systems))
	}
}

func TestRegistryListReturnsRegisteredSystems(t *testing.T) {
	registry := NewRegistry()

	first := System{
		ID:          "system-001",
		Name:        "Payments Gateway",
		Description: "External payments processing system",
		Status:      SystemStatusActive,
	}

	second := System{
		ID:          "system-002",
		Name:        "Market Data Platform",
		Description: "External market data system",
		Status:      SystemStatusInactive,
	}

	if err := registry.Register(first); err != nil {
		t.Fatalf("expected first registration to succeed, got %v", err)
	}

	if err := registry.Register(second); err != nil {
		t.Fatalf("expected second registration to succeed, got %v", err)
	}

	systems := registry.List()

	if len(systems) != 2 {
		t.Fatalf("expected 2 systems, got %d", len(systems))
	}

	found := make(map[string]System)
	for _, system := range systems {
		found[system.ID] = system
	}

	if found[first.ID] != first {
		t.Fatalf("expected first system %+v, got %+v", first, found[first.ID])
	}

	if found[second.ID] != second {
		t.Fatalf("expected second system %+v, got %+v", second, found[second.ID])
	}
}

func TestRegistryListReturnsSystemValuesWithoutMutatingRegistry(t *testing.T) {
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

	systems := registry.List()

	if len(systems) != 1 {
		t.Fatalf("expected 1 system, got %d", len(systems))
	}

	systems[0].Name = "Modified Gateway"

	registered, exists := registry.Get("system-001")
	if !exists {
		t.Fatal("expected registered system to be found")
	}

	if registered.Name != "Payments Gateway" {
		t.Fatalf("expected registry to retain original name, got %q", registered.Name)
	}
}

func TestRegistryListReturnsIndependentCollection(t *testing.T) {
	registry := NewRegistry()

	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	if err := registry.Register(system); err != nil {
		t.Fatalf("expected registration to succeed, got %v", err)
	}

	systems := registry.List()

	if len(systems) != 1 {
		t.Fatalf("expected 1 system, got %d", len(systems))
	}

	systems[0] = System{
		ID:     "system-999",
		Name:   "Modified System",
		Status: SystemStatusInactive,
	}

	if len(registry.List()) != 1 {
		t.Fatalf("expected registry to still contain 1 system, got %d", len(registry.List()))
	}

	registered, exists := registry.Get("system-001")
	if !exists {
		t.Fatal("expected original system to remain registered")
	}

	if registered != system {
		t.Fatalf("expected original system %+v, got %+v", system, registered)
	}
}

func TestRegistryListByStatusReturnsMatchingSystems(t *testing.T) {
	registry := NewRegistry()

	active := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	inactive := System{
		ID:     "system-002",
		Name:   "Market Data Platform",
		Status: SystemStatusInactive,
	}

	secondActive := System{
		ID:     "system-003",
		Name:   "Risk Analytics Platform",
		Status: SystemStatusActive,
	}

	if err := registry.Register(active); err != nil {
		t.Fatalf("expected active system registration to succeed, got %v", err)
	}

	if err := registry.Register(inactive); err != nil {
		t.Fatalf("expected inactive system registration to succeed, got %v", err)
	}

	if err := registry.Register(secondActive); err != nil {
		t.Fatalf("expected second active system registration to succeed, got %v", err)
	}

	systems := registry.ListByStatus(SystemStatusActive)

	if len(systems) != 2 {
		t.Fatalf("expected 2 active systems, got %d", len(systems))
	}

	found := make(map[string]System)
	for _, system := range systems {
		found[system.ID] = system
	}

	if found[active.ID] != active {
		t.Fatalf("expected active system %+v, got %+v", active, found[active.ID])
	}

	if found[secondActive.ID] != secondActive {
		t.Fatalf("expected second active system %+v, got %+v", secondActive, found[secondActive.ID])
	}

	if _, exists := found[inactive.ID]; exists {
		t.Fatalf("did not expect inactive system %q in active results", inactive.ID)
	}
}

func TestRegistryListByStatusReturnsEmptyCollectionWhenNoSystemsMatch(t *testing.T) {
	registry := NewRegistry()

	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	if err := registry.Register(system); err != nil {
		t.Fatalf("expected registration to succeed, got %v", err)
	}

	systems := registry.ListByStatus(SystemStatusInactive)

	if systems == nil {
		t.Fatal("expected ListByStatus to return an empty collection, got nil")
	}

	if len(systems) != 0 {
		t.Fatalf("expected no inactive systems, got %d", len(systems))
	}
}

func TestRegistryListByStatusReturnsIndependentCollection(t *testing.T) {
	registry := NewRegistry()

	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	if err := registry.Register(system); err != nil {
		t.Fatalf("expected registration to succeed, got %v", err)
	}

	systems := registry.ListByStatus(SystemStatusActive)

	if len(systems) != 1 {
		t.Fatalf("expected 1 active system, got %d", len(systems))
	}

	systems[0] = System{
		ID:     "system-999",
		Name:   "Modified System",
		Status: SystemStatusInactive,
	}

	registered, exists := registry.Get(system.ID)
	if !exists {
		t.Fatal("expected original system to remain registered")
	}

	if registered != system {
		t.Fatalf("expected original system %+v, got %+v", system, registered)
	}
}
