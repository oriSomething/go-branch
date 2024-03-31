package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	Idle State = iota + 1
	Quiting
)

type model struct {
	branches      []string
	selected      int
	quittingError *string
	state         State
}

func InitialModel(branches []string) model {
	return model{
		branches:      branches,
		selected:      0,
		quittingError: nil,
		state:         Idle,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q", "esc":
			m.state = Quiting
			return m, tea.Quit

		case "up":
			if m.selected > 0 {
				m.selected--
			}

		case "down":
			if m.selected < len(m.branches)-1 {
				m.selected++
			}

		case "enter", " ":
			if err := SwitchBranch(m.branches[m.selected]); err != nil {
				quittingError := err.Error()
				m.quittingError = &quittingError
			}
			m.state = Quiting

			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.state == Quiting {
		return ""
	}

	s := ""

	for i, choice := range m.branches {
		cursor := " [ ] "
		if m.selected == i {
			cursor = " [*] "
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}
