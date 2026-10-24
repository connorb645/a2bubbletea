// Package render turns an A2UI surface into terminal text with Lip Gloss.
package render

import (
	"strings"

	"github.com/connorb645/a2bubbletea/internal/a2ui"
)

// Renderer draws A2UI components for a fixed terminal width.
type Renderer struct {
	theme Theme
}

// New returns a renderer using the default theme.
func New() *Renderer {
	return &Renderer{theme: DefaultTheme()}
}

// Render draws the surface's root component at the given width.
func (renderer *Renderer) Render(surface *a2ui.Surface, width int) string {
	width = max(width, 1)

	root, ok := surface.Components["root"]
	if !ok {
		return renderer.theme.Caption.Width(width).Render("waiting for a root component")
	}
	return renderer.render(root, width)
}

func (renderer *Renderer) render(component a2ui.Component, width int) string {
	if component.Type != "Text" {
		return renderer.theme.Caption.Width(width).Render("[" + component.Type + "]")
	}

	style := renderer.theme.Body
	if component.Variant == "caption" {
		style = renderer.theme.Caption
	}
	if title, ok := heading(component.Text); ok {
		style = renderer.theme.Heading
		return style.Width(width).Render(title)
	}
	return style.Width(width).Render(component.Text)
}

func heading(text string) (string, bool) {
	for level := 1; level <= 3; level++ {
		prefix := strings.Repeat("#", level) + " "
		if title, ok := strings.CutPrefix(text, prefix); ok {
			return strings.TrimSpace(title), true
		}
	}
	return text, false
}
