package a2ui

import "fmt"

// Store owns the surfaces built from a message stream.
type Store struct {
	surfaces map[string]*Surface
	order    []string
}

// NewStore returns an empty surface store.
func NewStore() *Store {
	return &Store{surfaces: map[string]*Surface{}}
}

// Apply mutates the store with one message.
func (store *Store) Apply(message Message) error {
	if message.Version != Version {
		return fmt.Errorf("unsupported protocol version %q", message.Version)
	}
	if count := message.payloadCount(); count != 1 {
		return fmt.Errorf("message has %d payloads, want exactly one", count)
	}

	switch {
	case message.CreateSurface != nil:
		return store.createSurface(*message.CreateSurface)
	case message.UpdateComponents != nil:
		return store.updateComponents(*message.UpdateComponents)
	default:
		return fmt.Errorf("unknown message")
	}
}

// Current returns the most recently created surface.
func (store *Store) Current() (*Surface, bool) {
	if len(store.order) == 0 {
		return nil, false
	}
	surface, ok := store.surfaces[store.order[len(store.order)-1]]
	return surface, ok
}

func (store *Store) createSurface(create CreateSurface) error {
	if create.SurfaceID == "" {
		return fmt.Errorf("createSurface: missing surfaceId")
	}
	if _, exists := store.surfaces[create.SurfaceID]; exists {
		return fmt.Errorf("createSurface: surface %q already exists", create.SurfaceID)
	}

	surface := &Surface{ID: create.SurfaceID, Components: map[string]Component{}}
	for _, component := range create.Components {
		if component.ID == "" {
			return fmt.Errorf("createSurface: component without id")
		}
		surface.Components[component.ID] = component
	}

	store.surfaces[surface.ID] = surface
	store.order = append(store.order, surface.ID)
	return nil
}

func (store *Store) updateComponents(update UpdateComponents) error {
	surface, ok := store.surfaces[update.SurfaceID]
	if !ok {
		return fmt.Errorf("updateComponents: unknown surface %q", update.SurfaceID)
	}
	for _, component := range update.Components {
		if component.ID == "" {
			return fmt.Errorf("updateComponents: component without id")
		}
		surface.Components[component.ID] = component
	}
	return nil
}
