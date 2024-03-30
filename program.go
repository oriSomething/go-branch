package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	branches     []string
	selected     int
	quittingText string
	quitting     bool
}

func InitialModel() model {
	branches, err := GetBranches(10)

	if err != nil {
		panic(err)
	}

	return model{
		branches:     branches,
		selected:     0,
		quittingText: "",
		quitting:     false,
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
				m.quittingText = err.Error()
			} else {
				m.quittingText = ""
			}

			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return m.quittingText
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
