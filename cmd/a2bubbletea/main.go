// Command a2bubbletea renders an A2UI JSONL message stream in the terminal.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/connorb645/a2bubbletea/internal/a2ui"
	"github.com/connorb645/a2bubbletea/internal/tui"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "a2bubbletea:", err)
		os.Exit(1)
	}
}

func run() error {
	file := flag.String("file", "", "path to an A2UI JSONL message stream")
	flag.Parse()

	path := *file
	if path == "" && flag.NArg() > 0 {
		path = flag.Arg(0)
	}
	if path == "" {
		return errors.New("no message file given; pass -file <path>")
	}

	stream, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer func() {
		_ = stream.Close()
	}()

	messages, err := a2ui.DecodeMessages(stream)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	store := a2ui.NewStore()
	for i, message := range messages {
		if err := store.Apply(message); err != nil {
			return fmt.Errorf("message %d: %w", i+1, err)
		}
	}
	if _, ok := store.Current(); !ok {
		return errors.New("message stream created no surface")
	}

	_, err = tea.NewProgram(tui.New(store)).Run()
	return err
}
