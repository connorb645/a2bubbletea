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

func TestRenderPlaceholder(t *testing.T) {
	surface := surfaceWith(
		a2ui.Component{ID: "root", Type: "Card"},
	)

	if got := New().Render(surface, 20); !strings.Contains(got, "[Card]") {
		t.Fatalf("rendered = %q", got)
	}
}

func TestRenderSample(t *testing.T) {
	file, err := os.Open("../../samples/text.jsonl")
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
	for _, want := range []string{"A2UI v1.0", "Bubble Tea"} {
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
