package styles

import "github.com/charmbracelet/lipgloss"

func ApiError(text string) string {
	styles := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	return styles.Render(text)
}

func BackwardRotation(text string) string {
	styles := lipgloss.NewStyle().
		Foreground(lipgloss.Color("201")).
		Bold(true)

	return styles.Render(text)
}

func ForwardRotation(text string) string {
	styles := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#EEEEEE")).
		Bold(true)

	return styles.Render(text)
}
