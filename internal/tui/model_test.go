package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/connorb645/a2bubbletea/internal/a2ui"
)

func TestView(t *testing.T) {
	store := a2ui.NewStore()
	create := a2ui.Message{
		Version: a2ui.Version,
		CreateSurface: &a2ui.CreateSurface{
			SurfaceID:  "main",
			Components: []a2ui.Component{{ID: "root", Type: "Text", Text: "Hello"}},
		},
	}
	if err := store.Apply(create); err != nil {
		t.Fatalf("Apply: %v", err)
	}

	if got := New(store).View(); !strings.Contains(got, "Hello") {
		t.Fatalf("View() = %q", got)
	}
}

func TestQuit(t *testing.T) {
	_, command := New(a2ui.NewStore()).Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	if command == nil {
		t.Fatal("q did not return a command")
	}
}
