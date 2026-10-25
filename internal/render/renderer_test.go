package render

import (
	"os"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/connorb645/a2bubbletea/internal/a2ui"
)

func TestRenderText(t *testing.T) {
	surface := surfaceWith(a2ui.Component{ID: "root", Type: "Text", Text: "Hello"})

	if got := New().Render(surface, 20); !strings.Contains(got, "Hello") {
		t.Fatalf("rendered = %q", got)
	}
}

func TestRenderHeading(t *testing.T) {
	surface := surfaceWith(a2ui.Component{ID: "root", Type: "Text", Text: "# Title"})

	rendered := New().Render(surface, 20)
	if !strings.Contains(rendered, "Title") || strings.Contains(rendered, "#") {
		t.Fatalf("rendered = %q", rendered)
	}
}

func TestRenderPlaceholder(t *testing.T) {
	surface := surfaceWith(a2ui.Component{ID: "root", Type: "Row"})

	if got := New().Render(surface, 20); !strings.Contains(got, "[Row]") {
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
	if !strings.Contains(rendered, "A2UI v1.0") {
		t.Fatalf("rendered = %q", rendered)
	}
	for _, line := range strings.Split(rendered, "\n") {
		if width := lipgloss.Width(line); width > 80 {
			t.Errorf("line width = %d, want at most 80", width)
		}
	}
}

func surfaceWith(root a2ui.Component) *a2ui.Surface {
	return &a2ui.Surface{Components: map[string]a2ui.Component{"root": root}}
}
