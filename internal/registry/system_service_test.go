package registry

import (
	"errors"
	"testing"
)

var errRepositoryFailure = errors.New("repository failure")

type fakeSystemRepository struct {
	registerCalled      bool
	registerSystem      System
	registerErr         error
	getCalled           bool
	getID               string
	getSystem           System
	getExists           bool
	listCalled          bool
	listSystems         []System
	listByStatusCalled  bool
	listByStatus        SystemStatus
	listByStatusSystems []System
}

func (f *fakeSystemRepository) Register(system System) error {
	f.registerCalled = true
	f.registerSystem = system
	return f.registerErr
}

func (f *fakeSystemRepository) Get(id string) (System, bool) {
	f.getCalled = true
	f.getID = id
	return f.getSystem, f.getExists
}

func (f *fakeSystemRepository) List() []System {
	f.listCalled = true
	return f.listSystems
}

func (f *fakeSystemRepository) ListByStatus(status SystemStatus) []System {
	f.listByStatusCalled = true
	f.listByStatus = status
	return f.listByStatusSystems
}

var _ SystemRepository = (*fakeSystemRepository)(nil)

func TestNewSystemServiceCreatesService(t *testing.T) {
	repository := &fakeSystemRepository{}

	service := NewSystemService(repository)

	if service == nil {
		t.Fatal("expected service to be created")
	}

	if service.repository != repository {
		t.Fatal("expected service to use provided repository")
	}
}

func TestSystemServiceRegisterDelegatesToRepository(t *testing.T) {
	repository := &fakeSystemRepository{}

	service := NewSystemService(repository)

	system := System{
		ID:          "system-001",
		Name:        "Payments Gateway",
		Description: "External payments processing system",
		Status:      SystemStatusActive,
	}

	if err := service.Register(system); err != nil {
		t.Fatalf("expected registration to succeed, got %v", err)
	}

	if !repository.registerCalled {
		t.Fatal("expected repository Register to be called")
	}

	if repository.registerSystem != system {
		t.Fatalf("expected system %+v, got %+v", system, repository.registerSystem)
	}
}

func TestSystemServiceRegisterValidatesBeforeRepository(t *testing.T) {
	repository := &fakeSystemRepository{}

	service := NewSystemService(repository)

	system := System{
		ID:     "",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	err := service.Register(system)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repository.registerCalled {
		t.Fatal("expected repository Register not to be called")
	}
}

func TestSystemServiceRegisterRejectsMissingName(t *testing.T) {
	repository := &fakeSystemRepository{}

	service := NewSystemService(repository)

	system := System{
		ID:     "system-001",
		Status: SystemStatusActive,
	}

	err := service.Register(system)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != "system name is required" {
		t.Fatalf("expected name validation error, got %v", err)
	}

	if repository.registerCalled {
		t.Fatal("expected repository Register not to be called")
	}
}

func TestSystemServiceRegisterRejectsInvalidStatus(t *testing.T) {
	repository := &fakeSystemRepository{}

	service := NewSystemService(repository)

	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatus("unknown"),
	}

	err := service.Register(system)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if err.Error() != `invalid system status: "unknown"` {
		t.Fatalf("expected status validation error, got %v", err)
	}

	if repository.registerCalled {
		t.Fatal("expected repository Register not to be called")
	}
}

func TestSystemServiceRegisterReturnsRepositoryError(t *testing.T) {
	repository := &fakeSystemRepository{
		registerErr: errRepositoryFailure,
	}

	service := NewSystemService(repository)

	system := System{
		ID:     "system-001",
		Name:   "Payments Gateway",
		Status: SystemStatusActive,
	}

	err := service.Register(system)

	if err != errRepositoryFailure {
		t.Fatalf("expected repository error %v, got %v", errRepositoryFailure, err)
	}
}

func TestSystemServiceGetDelegatesToRepository(t *testing.T) {
	system := System{
		ID:          "system-001",
		Name:        "Payments Gateway",
		Description: "External payments processing system",
		Status:      SystemStatusActive,
	}

	repository := &fakeSystemRepository{
		getSystem: system,
		getExists: true,
	}

	service := NewSystemService(repository)

	got, exists := service.Get("system-001")

	if !repository.getCalled {
		t.Fatal("expected repository Get to be called")
	}

	if repository.getID != "system-001" {
		t.Fatalf("expected ID %q, got %q", "system-001", repository.getID)
	}

	if !exists {
		t.Fatal("expected system to exist")
	}

	if got != system {
		t.Fatalf("expected system %+v, got %+v", system, got)
	}
}

func TestSystemServiceGetReturnsRepositoryResult(t *testing.T) {
	repository := &fakeSystemRepository{
		getExists: false,
	}

	service := NewSystemService(repository)

	got, exists := service.Get("system-999")

	if exists {
		t.Fatal("expected system not to exist")
	}

	if got != (System{}) {
		t.Fatalf("expected zero-value system, got %+v", got)
	}
}

func TestSystemServiceListDelegatesToRepository(t *testing.T) {
	systems := []System{
		{
			ID:     "system-001",
			Name:   "Payments Gateway",
			Status: SystemStatusActive,
		},
		{
			ID:     "system-002",
			Name:   "Market Data Platform",
			Status: SystemStatusInactive,
		},
	}

	repository := &fakeSystemRepository{
		listSystems: systems,
	}

	service := NewSystemService(repository)

	got := service.List()

	if !repository.listCalled {
		t.Fatal("expected repository List to be called")
	}

	if len(got) != len(systems) {
		t.Fatalf("expected %d systems, got %d", len(systems), len(got))
	}

	for i := range systems {
		if got[i] != systems[i] {
			t.Fatalf("expected system %+v, got %+v", systems[i], got[i])
		}
	}
}

func TestSystemServiceListByStatusDelegatesToRepository(t *testing.T) {
	systems := []System{
		{
			ID:     "system-001",
			Name:   "Payments Gateway",
			Status: SystemStatusActive,
		},
	}

	repository := &fakeSystemRepository{
		listByStatusSystems: systems,
	}

	service := NewSystemService(repository)

	got := service.ListByStatus(SystemStatusActive)

	if !repository.listByStatusCalled {
		t.Fatal("expected repository ListByStatus to be called")
	}

	if repository.listByStatus != SystemStatusActive {
		t.Fatalf(
			"expected status %q, got %q",
			SystemStatusActive,
			repository.listByStatus,
		)
	}

	if len(got) != len(systems) {
		t.Fatalf("expected %d systems, got %d", len(systems), len(got))
	}

	if got[0] != systems[0] {
		t.Fatalf("expected system %+v, got %+v", systems[0], got[0])
	}
}
