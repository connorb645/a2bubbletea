package render

import "github.com/charmbracelet/lipgloss"

// Theme holds the styles the renderer applies to each component kind.
type Theme struct {
	Body    lipgloss.Style
	Caption lipgloss.Style
	Heading lipgloss.Style
}

// DefaultTheme returns a theme that adapts to light and dark terminals.
func DefaultTheme() Theme {
	muted := lipgloss.AdaptiveColor{Light: "242", Dark: "245"}

	return Theme{
		Body:    lipgloss.NewStyle(),
		Caption: lipgloss.NewStyle().Foreground(muted).Italic(true),
		Heading: lipgloss.NewStyle().Bold(true),
	}
}
