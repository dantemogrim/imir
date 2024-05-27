package styles

import "github.com/charmbracelet/lipgloss"

// var myCuteBorder = lipgloss.Border{
// 	Top:         "._.:*:",
// 	Bottom:      "._.:*:",
// 	Left:        "|*",
// 	Right:       "|*",
// 	TopLeft:     "*",
// 	TopRight:    "*",
// 	BottomLeft:  "*",
// 	BottomRight: "*",
// }

func Err(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	return styles.Render(text)
}

func Retrograde(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#f456aa")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)

	return styles.Render(text)
}

func NotRetrograde(text string) string {
	var styles lipgloss.Style = lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Foreground(lipgloss.Color("#fafafa")).
		Background(lipgloss.Color("#5a40a7")).
		Padding(1, 1, 2, 1).
		Align(lipgloss.Center).
		MarginBottom(1).
		Width(25)

	return styles.Render(text)
}
