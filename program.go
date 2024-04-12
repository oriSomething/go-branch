package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func InitialModel(branches []string) model {
	ti := textinput.New()
	ti.CharLimit = 156
	ti.Width = 100
	ti.Focus()

	maxWidthBranch := 0
	for i := 0; i < len(branches); i++ {
		branchSize := utf8.RuneCountInString(branches[i])
		if maxWidthBranch < branchSize {
			maxWidthBranch = branchSize
		}
	}

	return model{
		noFilter: filterModel{
			branches: branches,
			index:    0,
			selected: 0,
		},
		filter:        nil,
		quittingError: nil,
		state:         Idle,

		textInput:      ti,
		maxWidthBranch: maxWidthBranch,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		key := msg.String()
		switch key {

		case "ctrl+c", "esc":
			m.state = Quiting
			return m, tea.Quit

		case "up":
			m.Up()
			return m, nil

		case "down":
			m.Down()
			return m, nil

		case "enter", " ":
			c := m.GetCurrent()

			if c.selected == -1 {
				return m, nil
			}

			if err := SwitchBranch(c.branches[c.selected]); err != nil {
				quittingError := err.Error()
				m.quittingError = &quittingError
			}
			m.state = Quiting

			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)

	// TODO: Use fuzzy search
	if query := m.GetQuery(); m.currentQuery != query {
		m.currentQuery = query
		m.filter = m.noFilter.Filter(m.currentQuery)
		m.noFilter.index = 0
		m.noFilter.selected = 0
	}

	return m, cmd
}

func (m model) View() string {
	if m.state == Quiting {
		return ""
	}

	s := ""

	styleChosen := lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		Background(lipgloss.Color("8")).
		Width(m.maxWidthBranch)

	current := m.GetCurrent()

	for i, j := current.index, 0; current.index != -1 && i < len(current.branches) && j < limit; i, j = i+1, j+1 {
		choice := current.branches[i]

		// Render the row
		if current.selected == i {
			s += fmt.Sprintf("  %s\n", styleChosen.Render(choice))
		} else {
			s += fmt.Sprintf("  %s\n", choice)
		}
	}
	return fmt.Sprintf("%s\n\n%s\n", m.textInput.View(), s)
}
