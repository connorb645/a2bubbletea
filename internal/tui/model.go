// Package tui hosts the Bubble Tea program that displays a surface.
package tui

import "github.com/charmbracelet/bubbletea"

// Model is the terminal program.
type Model struct{}

// New returns the starting model.
func New() Model {
	return Model{}
}

// Init starts the program.
func (model Model) Init() tea.Cmd {
	return nil
}

// Update handles quit keys.
func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := message.(tea.KeyMsg); ok {
		switch key.String() {
		case "q", "esc", "ctrl+c":
			return model, tea.Quit
		}
	}
	return model, nil
}

// View draws the program.
func (model Model) View() string {
	return "a2bubbletea"
}
