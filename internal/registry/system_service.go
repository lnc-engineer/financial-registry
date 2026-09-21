package registry

// SystemService provides application-level operations for registered systems.
type SystemService struct {
	repository SystemRepository
}

// NewSystemService creates a system service using the provided repository.
func NewSystemService(repository SystemRepository) *SystemService {
	return &SystemService{
		repository: repository,
	}
}

// Register registers a system through the configured repository.
func (s *SystemService) Register(system System) error {
	if err := system.Validate(); err != nil {
		return err
	}

	return s.repository.Register(system)
}

// Get retrieves a registered system by ID through the configured repository.
func (s *SystemService) Get(id string) (System, bool) {
	return s.repository.Get(id)
}

// List returns all registered systems through the configured repository.
func (s *SystemService) List() []System {
	return s.repository.List()
}

// ListByStatus returns registered systems matching the requested status.
func (s *SystemService) ListByStatus(status SystemStatus) []System {
	return s.repository.ListByStatus(status)
}
