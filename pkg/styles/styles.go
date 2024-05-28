package styles

import "github.com/charmbracelet/lipgloss"

func Err(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff6b6b")).
		Padding(1, 1, 2, 1).
		Align(lipgloss.Center).
		MarginBottom(1).
		Width(25)

	return styles.Render(text)
}

func Retrograde(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ffbd5a")).
		Padding(1, 1, 2, 1).
		Align(lipgloss.Center).
		MarginBottom(1).
		Width(25)

	return styles.Render(text)
}

func NotRetrograde(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Foreground(lipgloss.Color("#d5ffcb")).
		Padding(1, 1, 2, 1).
		Align(lipgloss.Center).
		MarginBottom(1).
		Width(25)

	return styles.Render(text)
}
