package render

import "github.com/charmbracelet/lipgloss"

// Theme holds the styles the renderer applies to each component kind.
type Theme struct {
	Body    lipgloss.Style
	Caption lipgloss.Style
	Heading lipgloss.Style
	Card    lipgloss.Style
	Divider lipgloss.Style
}

// DefaultTheme returns a theme that adapts to light and dark terminals.
func DefaultTheme() Theme {
	muted := lipgloss.AdaptiveColor{Light: "242", Dark: "245"}
	border := lipgloss.AdaptiveColor{Light: "250", Dark: "238"}

	return Theme{
		Body:    lipgloss.NewStyle(),
		Caption: lipgloss.NewStyle().Foreground(muted).Italic(true),
		Heading: lipgloss.NewStyle().Bold(true),
		Card:    lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(border).Padding(1, 2),
		Divider: lipgloss.NewStyle().Foreground(border),
	}
}
