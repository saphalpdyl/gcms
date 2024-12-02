package helpers

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

func RenderDoctorResult(t string, isSuccess bool, helpText string) string {
	prefixStyle := lipgloss.
		NewStyle()

	var prefix string

	if isSuccess {
		prefix = prefixStyle.Foreground(lipgloss.Color("#04B575")).Render("[✓]")
	} else {
		prefix = prefixStyle.Foreground(lipgloss.Color("9")).Render("[×]")
	}

	titleText := lipgloss.
		NewStyle().
		Bold(true).
		Underline(true).
		Render(t)

	var renderedHelpText string

	if helpText != "" {
		renderedHelpText = lipgloss.
			NewStyle().
			Foreground(lipgloss.Color("#888")).
			Render(fmt.Sprintf("\n\t%s", helpText))
	}

	return fmt.Sprintf("%s %s%s", prefix, titleText, renderedHelpText)
}
