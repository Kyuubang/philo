package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type MDModel struct {
	viewport viewport.Model
}

func MDReader(content string) (*MDModel, error) {
	const width = 78

	vp := viewport.New(width, 20)
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

	return &MDModel{
		viewport: vp,
	}, nil
}

func (e MDModel) Init() tea.Cmd {
	return nil
}

func (e MDModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return e, tea.Quit
		default:
			var cmd tea.Cmd
			e.viewport, cmd = e.viewport.Update(msg)
			return e, cmd
		}
	default:
		return e, nil
	}
}

func (e MDModel) View() string {
	return e.viewport.View() + e.helpView()
}

func (e MDModel) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • q: Quit\n")
}

//func main() {
//	modelPager, err := newExample()
//	if err != nil {
//		fmt.Println("Could not initialize Bubble Tea modelPager:", err)
//		os.Exit(1)
//	}
//
//	if _, err := tea.NewProgram(modelPager).Run(); err != nil {
//		fmt.Println("Bummer, there's been an error:", err)
//		os.Exit(1)
//	}
//}
