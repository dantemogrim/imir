package styles

import "github.com/charmbracelet/lipgloss"

func ApiError(text string) string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	return style.Render(text)
}

func Backward(text string) string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("201")).
		Bold(true)

	return style.Render(text)
}

func Forward(text string) string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#EEEEEE")).
		Bold(true)

	return style.Render(text)
}
