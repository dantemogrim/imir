package pager

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/dantemogrim/imir/pkg/utils"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).PaddingBottom(1).Render

type pager struct {
	viewport viewport.Model
}

func newPager() (*pager, error) {
	dimensions, err := utils.TerminalDimensions()
	if err != nil {
		fmt.Println("Error getting terminal dimensions:", err)
	}

	var width int = dimensions[0] - 5
	var height int = 20

	vp := viewport.New(width, height)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &pager{
		viewport: vp,
	}, nil
}

func (p pager) Init() tea.Cmd {
	return nil
}

func (p pager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return p, tea.Quit
		default:
			var cmd tea.Cmd
			p.viewport, cmd = p.viewport.Update(msg)
			return p, cmd
		}
	default:
		return p, nil
	}
}

func (p pager) View() string {
	return p.viewport.View() + p.helpView()
}

func (p pager) helpView() string {
	return helpStyle("\n ↑/↓: Navigate • q: Quit\n")
}

func Render() {
	model, err := newPager()
	if err != nil {
		fmt.Println("Could not initialize Bubble Tea model:", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
	}
}
