package registry

// SystemRepository defines storage operations for registered systems.
type SystemRepository interface {
	Register(system System) error
	Get(id string) (System, bool)
	List() []System
	ListByStatus(status SystemStatus) []System
}

// Ensure Registry implements SystemRepository.
var _ SystemRepository = (*Registry)(nil)
