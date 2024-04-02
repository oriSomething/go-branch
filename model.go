package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
)

const limit = 10

type State int

const (
	Idle State = iota + 1
	Quiting
)

type model struct {
	currentQuery  string
	noFilter      filterModel
	filter        *filterModel
	quittingError *string
	state         State

	textInput      textinput.Model
	maxWidthBranch int
}

func (m *model) GetCurrent() *filterModel {
	if m.filter == nil {
		return &m.noFilter
	} else {
		return m.filter
	}
}

func (m *model) GetQuery() string {
	return strings.ToLower(strings.TrimSpace(m.textInput.Value()))
}

func (m *model) Up() {
	if c := m.GetCurrent(); c.selected > 0 {
		c.selected--

		if c.selected == c.index-1 {
			c.index--
		}
	}
}

func (m *model) Down() {
	if c := m.GetCurrent(); c.selected < len(c.branches)-1 {
		c.selected++

		if c.selected >= c.index+limit {
			c.index++
		}
	}
}
