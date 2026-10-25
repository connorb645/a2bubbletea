// Package tui hosts the Bubble Tea program that displays a surface.
package tui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/connorb645/a2bubbletea/internal/a2ui"
	"github.com/connorb645/a2bubbletea/internal/render"
)

// Model displays the store's current surface.
type Model struct {
	store    *a2ui.Store
	renderer *render.Renderer
	width    int
}

// New returns a model that renders the store at the terminal width.
func New(store *a2ui.Store) Model {
	return Model{store: store, renderer: render.New(), width: 80}
}

// Init starts the program.
func (model Model) Init() tea.Cmd {
	return nil
}

// Update handles terminal resizes and quit keys.
func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		model.width = message.Width
	case tea.KeyMsg:
		switch message.String() {
		case "q", "esc", "ctrl+c":
			return model, tea.Quit
		}
	}
	return model, nil
}

// View draws the current surface.
func (model Model) View() string {
	surface, ok := model.store.Current()
	if !ok {
		return "No surface. Press q to quit."
	}
	return model.renderer.Render(surface, model.width)
}
