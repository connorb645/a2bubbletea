// Command a2bubbletea renders an A2UI message stream in the terminal.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/connorb645/a2bubbletea/internal/tui"
)

func main() {
	if _, err := tea.NewProgram(tui.New()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "a2bubbletea:", err)
		os.Exit(1)
	}
}
