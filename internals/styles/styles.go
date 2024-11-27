package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func RenderBold(t string) string {
	return lipgloss.
		NewStyle().
		Bold(true).
		Render(t)
}

func RenderDiff(t string, isPositiveChange bool, prefix string) string {
	var textColorCode string

	if isPositiveChange {
		textColorCode = "#04B575"
	} else {
		textColorCode = "9"
	}

	return lipgloss.
		NewStyle().
		Bold(true).
		Underline(true).
		Foreground(lipgloss.Color(textColorCode)).
		Render(fmt.Sprintf("%s%s", prefix, t))
}
