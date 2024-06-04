package styles

import "github.com/charmbracelet/lipgloss"

func Err(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff7d7d")).
		Padding(1).
		Width(50)

	return styles.Render(text)
}

func Retrograde(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ffbd5a")).
		Padding(1).
		Align(lipgloss.Center).
		Width(25)

	return styles.Render(text)
}

func NotRetrograde(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Foreground(lipgloss.Color("#d5ffcb")).
		Padding(1).
		Align(lipgloss.Center).
		Width(25)

	return styles.Render(text)
}
