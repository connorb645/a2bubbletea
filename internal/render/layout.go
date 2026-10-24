package render

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/connorb645/a2bubbletea/internal/a2ui"
)

func (renderer *Renderer) row(component a2ui.Component, surface *a2ui.Surface, width int) block {
	if len(component.Children) == 0 {
		return newBlock("")
	}

	widths := fitWidths(renderer.naturalWidths(component.Children, surface, width), width)
	blocks := make([]block, len(component.Children))
	height := 1
	for i, id := range component.Children {
		blocks[i] = renderer.child(id, surface, widths[i], 0)
		height = max(height, blocks[i].height())
	}
	for i, id := range component.Children {
		child, ok := surface.Components[id]
		if ok && child.Type == "Divider" && child.Axis == "vertical" {
			blocks[i] = renderer.render(child, surface, widths[i], height)
		}
	}
	for i := range blocks {
		blocks[i] = blocks[i].fit(widths[i], height, component.Align)
	}

	gaps, leading := distribute(width-sum(widths), len(widths), component.Justify)
	return joinRow(blocks, widths, leading, gaps, width)
}

func (renderer *Renderer) column(component a2ui.Component, surface *a2ui.Surface, width, height int) block {
	if len(component.Children) == 0 {
		return newBlock("")
	}

	blocks := make([]block, len(component.Children))
	for i, id := range component.Children {
		childWidth := width
		if component.Align == "center" || component.Align == "end" {
			childWidth = min(renderer.intrinsicWidth(id, surface, width), width)
		}
		blocks[i] = renderer.child(id, surface, childWidth, 0)
		if childWidth < width {
			blocks[i] = blocks[i].place(width, component.Align)
		}
	}

	total := 0
	for _, child := range blocks {
		total += child.height()
	}

	gaps, leading := distribute(max(height-total, 0), len(blocks), component.Justify)
	return joinColumn(blocks, leading, gaps)
}

func (renderer *Renderer) naturalWidths(ids []string, surface *a2ui.Surface, width int) []int {
	widths := make([]int, len(ids))
	for i, id := range ids {
		widths[i] = renderer.intrinsicWidth(id, surface, width)
	}
	return widths
}

func (renderer *Renderer) intrinsicWidth(id string, surface *a2ui.Surface, limit int) int {
	component, ok := surface.Components[id]
	if !ok {
		return min(lipgloss.Width("["+id+"]"), limit)
	}

	switch component.Type {
	case "Text":
		text := component.Text
		if title, ok := heading(text); ok {
			text = title
		}
		width := 0
		for _, line := range strings.Split(text, "\n") {
			width = max(width, lipgloss.Width(line))
		}
		return clamp(width, 1, limit)
	case "Row":
		total := 0
		for _, child := range component.Children {
			total += renderer.intrinsicWidth(child, surface, limit)
		}
		return clamp(total, 1, limit)
	case "Column":
		width := 0
		for _, child := range component.Children {
			width = max(width, renderer.intrinsicWidth(child, surface, limit))
		}
		return clamp(width, 1, limit)
	case "Card":
		inner := renderer.intrinsicWidth(component.Child, surface, max(limit-6, 1))
		return clamp(inner+6, 1, limit)
	case "Divider":
		if component.Axis == "vertical" {
			return 1
		}
		return max(limit, 1)
	default:
		return clamp(lipgloss.Width("["+component.Type+"]"), 1, limit)
	}
}

func fitWidths(widths []int, total int) []int {
	available := 0
	for _, width := range widths {
		available += width
	}
	if available <= total {
		return widths
	}

	fitted := make([]int, len(widths))
	remaining := total
	for i, width := range widths {
		if i == len(widths)-1 {
			fitted[i] = max(remaining, 1)
			break
		}
		share := max(width*total/available, 1)
		reserve := len(widths) - i - 1
		share = min(share, max(remaining-reserve, 1))
		fitted[i] = share
		remaining -= share
	}
	return fitted
}

func distribute(extra, count int, justify string) (gaps []int, leading int) {
	if count < 1 {
		return nil, 0
	}

	gaps = make([]int, max(count-1, 0))
	if extra <= 0 {
		return gaps, 0
	}
	switch justify {
	case "end":
		leading = extra
	case "center":
		leading = extra / 2
	case "spaceBetween":
		if count == 1 {
			return gaps, 0
		}
		base := extra / (count - 1)
		remainder := extra % (count - 1)
		for i := range gaps {
			gaps[i] = base
			if i < remainder {
				gaps[i]++
			}
		}
	case "spaceAround":
		unit := extra / count
		leading = unit / 2
		for i := range gaps {
			gaps[i] = unit
		}
	case "spaceEvenly":
		unit := extra / (count + 1)
		leading = unit
		for i := range gaps {
			gaps[i] = unit
		}
	}
	return gaps, leading
}

func joinRow(blocks []block, widths []int, leading int, gaps []int, total int) block {
	height := 1
	for _, child := range blocks {
		height = max(height, child.height())
	}

	lines := make([]string, height)
	for row := range height {
		var line strings.Builder
		line.WriteString(strings.Repeat(" ", leading))
		for i, child := range blocks {
			if i > 0 {
				line.WriteString(strings.Repeat(" ", gaps[i-1]))
			}
			text := ""
			if row < child.height() {
				text = child.lines[row]
			}
			line.WriteString(padRight(text, widths[i]))
		}
		lines[row] = padRight(line.String(), total)
	}
	return block{lines: lines}
}

func joinColumn(blocks []block, leading int, gaps []int) block {
	var lines []string
	lines = append(lines, make([]string, leading)...)
	for i, child := range blocks {
		if i > 0 {
			lines = append(lines, make([]string, gaps[i-1])...)
		}
		lines = append(lines, child.lines...)
	}
	return block{lines: lines}
}

func sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func clamp(value, low, high int) int {
	return min(max(value, low), high)
}
