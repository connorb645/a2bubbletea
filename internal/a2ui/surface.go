package a2ui

// Surface is one rendered UI region and its flat component list.
type Surface struct {
	ID         string
	Components map[string]Component
}
