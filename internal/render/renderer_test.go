package render

import (
	"os"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/connorb645/a2bubbletea/internal/a2ui"
)

func TestRenderText(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Text", Text: "Hello"},
	)

	if got := New().Render(surface, 20); !strings.Contains(got, "Hello") {
		t.Fatalf("rendered = %q", got)
	}
}

func TestRenderHeading(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Text", Text: "# Title"},
	)

	rendered := New().Render(surface, 20)
	if !strings.Contains(rendered, "Title") || strings.Contains(rendered, "#") {
		t.Fatalf("rendered = %q", rendered)
	}
}

func TestRenderRowSpaceBetween(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Row", Children: []string{"left", "right"}, Justify: "spaceBetween"},
		a2ui.Component{ID: "left", Type: "Text", Text: "A"},
		a2ui.Component{ID: "right", Type: "Text", Text: "B"},
	)

	lines := strings.Split(New().Render(surface, 10), "\n")
	if len(lines) != 1 || lines[0] != "A        B" {
		t.Fatalf("rendered = %q", lines)
	}
}

func TestRenderColumnCenter(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Column", Children: []string{"text"}, Align: "center"},
		a2ui.Component{ID: "text", Type: "Text", Text: "A"},
	)

	if got := New().Render(surface, 5); got != "  A  " {
		t.Fatalf("rendered = %q", got)
	}
}

func TestRenderCard(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Card", Child: "text"},
		a2ui.Component{ID: "text", Type: "Text", Text: "Hi"},
	)

	lines := strings.Split(New().Render(surface, 20), "\n")
	if len(lines) != 5 {
		t.Fatalf("got %d lines, want 5", len(lines))
	}
	for _, line := range lines {
		if width := lipgloss.Width(line); width != 20 {
			t.Fatalf("line width = %d, want 20: %q", width, line)
		}
	}
	if !strings.Contains(lines[0], "╭") || !strings.Contains(lines[4], "╰") {
		t.Fatalf("missing rounded border: %q", strings.Join(lines, "\n"))
	}
}

func TestRenderDivider(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Divider", Axis: "horizontal"},
	)

	if got := New().Render(surface, 10); got != "──────────" {
		t.Fatalf("rendered = %q", got)
	}
}

func TestRenderVerticalDivider(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Row", Children: []string{"left", "rule", "right"}, Justify: "spaceBetween"},
		a2ui.Component{ID: "left", Type: "Text", Text: "A"},
		a2ui.Component{ID: "rule", Type: "Divider", Axis: "vertical"},
		a2ui.Component{ID: "right", Type: "Text", Text: "B"},
	)

	if got := New().Render(surface, 5); !strings.Contains(got, "│") {
		t.Fatalf("rendered = %q", got)
	}
}

func TestRenderPlaceholder(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Button"},
	)

	if got := New().Render(surface, 20); !strings.Contains(got, "[Button]") {
		t.Fatalf("rendered = %q", got)
	}
}

func TestRenderSample(t *testing.T) {
	file, err := os.Open("../../samples/static-layout.jsonl")
	if err != nil {
		t.Fatalf("open sample: %v", err)
	}
	defer func() {
		_ = file.Close()
	}()

	messages, err := a2ui.DecodeMessages(file)
	if err != nil {
		t.Fatalf("decode sample: %v", err)
	}

	store := a2ui.NewStore()
	for i, message := range messages {
		if err := store.Apply(message); err != nil {
			t.Fatalf("message %d: %v", i+1, err)
		}
	}
	surface, ok := store.Current()
	if !ok {
		t.Fatal("no surface")
	}

	rendered := New().Render(surface, 80)
	for _, want := range []string{
		"Austin, TX",
		"72°",
		"58°",
		"Clear skies with light breeze.",
		"Bubble Tea",
		"A2UI v1.0",
		"Terminal",
		"Static layout sample: Text, Row, Column, Card, Divider.",
		"─",
		"│",
		"╭",
	} {
		if !strings.Contains(rendered, want) {
			t.Errorf("rendered output missing %q", want)
		}
	}
	for _, line := range strings.Split(rendered, "\n") {
		if width := lipgloss.Width(line); width > 80 {
			t.Errorf("line width = %d, want at most 80", width)
		}
	}
}

func surfaceWith(components ...a2ui.Component) *a2ui.Surface {
	surface := &a2ui.Surface{Components: map[string]a2ui.Component{}}
	for _, component := range components {
		surface.Components[component.ID] = component
	}
	return surface
}
