package a2ui

// Component is one entry in a surface's flat component list.
type Component struct {
	ID      string `json:"id"`
	Type    string `json:"component"`
	Text    string `json:"text,omitempty"`
	Variant string `json:"variant,omitempty"`
}
