package styles

import "github.com/charmbracelet/lipgloss"

func RenderBold(t string) string {
	return lipgloss.
		NewStyle().
		Bold(true).
		Render(t)
}

func RenderDanger(t string) string {
	return lipgloss.
		NewStyle().
		Bold(true).
		Underline(true).
		Foreground(lipgloss.Color("9")).
		Render(t)
}
