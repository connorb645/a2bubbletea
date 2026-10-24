package a2ui

// UpdateComponents adds or replaces components in a surface.
type UpdateComponents struct {
	SurfaceID  string      `json:"surfaceId"`
	Components []Component `json:"components"`
}
