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
		return renderer.placeholder("waiting for a root component", width).lines[0]
	}
	return strings.Join(renderer.render(root, surface, width, 0).lines, "\n")
}

func (renderer *Renderer) render(component a2ui.Component, surface *a2ui.Surface, width, height int) block {
	width = max(width, 1)

	switch component.Type {
	case "Text":
		return renderer.text(component, width)
	case "Row":
		return renderer.row(component, surface, width)
	case "Column":
		return renderer.column(component, surface, width, height)
	case "Card":
		return renderer.card(component, surface, width)
	case "Divider":
		return renderer.divider(component, width, height)
	default:
		return renderer.placeholder(component.Type, width)
	}
}

func (renderer *Renderer) text(component a2ui.Component, width int) block {
	style := renderer.theme.Body
	if component.Variant == "caption" {
		style = renderer.theme.Caption
	}
	if title, ok := heading(component.Text); ok {
		style = renderer.theme.Heading
		return newBlock(style.Width(width).Render(title))
	}
	return newBlock(style.Width(width).Render(component.Text))
}

func (renderer *Renderer) card(component a2ui.Component, surface *a2ui.Surface, width int) block {
	cardWidth := max(width-2, 1)
	contentWidth := max(cardWidth-4, 1)
	content := renderer.child(component.Child, surface, contentWidth, 0)
	return newBlock(renderer.theme.Card.Width(cardWidth).Render(strings.Join(content.lines, "\n")))
}

func (renderer *Renderer) divider(component a2ui.Component, width, height int) block {
	if component.Axis == "vertical" {
		height = max(height, 1)
		return newBlock(renderer.theme.Divider.Render(strings.Repeat("│\n", height-1) + "│"))
	}
	return newBlock(renderer.theme.Divider.Render(strings.Repeat("─", width)))
}

func (renderer *Renderer) placeholder(name string, width int) block {
	return newBlock(renderer.theme.Caption.Width(width).Render("[" + name + "]"))
}

func (renderer *Renderer) child(id string, surface *a2ui.Surface, width, height int) block {
	child, ok := surface.Components[id]
	if !ok {
		return renderer.placeholder("missing: "+id, width)
	}
	return renderer.render(child, surface, width, height)
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
