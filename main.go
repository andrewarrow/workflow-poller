package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#0000AA")).
			Padding(0, 1).
			Width(30).
			Align(lipgloss.Center)

	itemStyle = lipgloss.NewStyle().
			Padding(0, 2)

	selectedItemStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#0000FF"))
)

type model struct {
	options  []string
	cursor   int
	selected string
}

func initialModel() model {
	return model{
		options: []string{
			"dev",
			"stg",
			"prd",
			"dev-packer",
			"stg-packer",
			"prd-packer",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.options) - 1
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}

		case "enter", " ":
			m.selected = m.options[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Select an option") + "\n\n"

	for i, option := range m.options {
		if m.cursor == i {
			s += selectedItemStyle.Render(option) + "\n"
		} else {
			s += itemStyle.Render(option) + "\n"
		}
	}

	s += "\n" + itemStyle.Render("Press q to quit, up/down to navigate, enter to select")

	if m.selected != "" {
		s = fmt.Sprintf("You selected: %s\n", m.selected)
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
