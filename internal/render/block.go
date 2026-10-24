package render

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// block is rendered component output as lines of terminal text.
type block struct {
	lines []string
}

func newBlock(text string) block {
	if text == "" {
		return block{lines: []string{""}}
	}
	return block{lines: strings.Split(text, "\n")}
}

func (b block) width() int {
	width := 0
	for _, line := range b.lines {
		width = max(width, lipgloss.Width(line))
	}
	return width
}

func (b block) height() int {
	return len(b.lines)
}

func (b block) fit(width, height int, verticalAlign string) block {
	offset := 0
	switch verticalAlign {
	case "center":
		offset = (height - b.height()) / 2
	case "end":
		offset = height - b.height()
	}

	lines := make([]string, height)
	for row := range height {
		source := row - offset
		line := ""
		if source >= 0 && source < b.height() {
			line = b.lines[source]
		}
		lines[row] = padRight(line, width)
	}
	return block{lines: lines}
}

func (b block) place(width int, horizontalAlign string) block {
	offset := 0
	switch horizontalAlign {
	case "center":
		offset = max((width-b.width())/2, 0)
	case "end":
		offset = max(width-b.width(), 0)
	}

	lines := make([]string, b.height())
	for i, line := range b.lines {
		lines[i] = padRight(strings.Repeat(" ", offset)+line, width)
	}
	return block{lines: lines}
}

func padRight(line string, width int) string {
	if gap := width - lipgloss.Width(line); gap > 0 {
		return line + strings.Repeat(" ", gap)
	}
	return line
}
