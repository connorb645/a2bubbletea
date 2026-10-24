package a2ui

// CreateSurface starts a surface, optionally with its first components.
type CreateSurface struct {
	SurfaceID  string      `json:"surfaceId"`
	Components []Component `json:"components,omitempty"`
}
