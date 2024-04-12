package main

import "strings"

type filterModel struct {
	branches []string
	index    int
	selected int
}

func (m *filterModel) Filter(query string) *filterModel {
	if query == "" {
		return nil
	}

	branches := make([]string, 0)

	for _, b := range m.branches {
		if strings.Contains(strings.ToLower(b), query) {
			branches = append(branches, b)
		}
	}

	selected := 0
	if len(branches) == 0 {
		selected = -1
	}

	return &filterModel{branches, 0, selected}
}
