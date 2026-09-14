package registry

import "fmt"

// Registry stores financial systems by their unique ID.
type Registry struct {
	systems map[string]System
}

// NewRegistry creates an empty system registry.
func NewRegistry() *Registry {
	return &Registry{
		systems: make(map[string]System),
	}
}

// Register adds a valid system to the registry.
func (r *Registry) Register(system System) error {
	if err := system.Validate(); err != nil {
		return err
	}

	if _, exists := r.systems[system.ID]; exists {
		return fmt.Errorf("system with ID %q is already registered", system.ID)
	}

	r.systems[system.ID] = system
	return nil
}

// Get returns a registered system by ID.
func (r *Registry) Get(id string) (System, bool) {
	system, exists := r.systems[id]
	return system, exists
}

// List returns all systems currently registered.
func (r *Registry) List() []System {
	systems := make([]System, 0, len(r.systems))

	for _, system := range r.systems {
		systems = append(systems, system)
	}

	return systems
}
