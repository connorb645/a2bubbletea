package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestView(t *testing.T) {
	if got := New().View(); got != "a2bubbletea" {
		t.Fatalf("View() = %q", got)
	}
}

func TestQuit(t *testing.T) {
	_, command := New().Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	if command == nil {
		t.Fatal("q did not return a command")
	}
}
