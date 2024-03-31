package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	branches, err := GetBranches(10)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// If no branches to select from we don't care
	if len(branches) == 0 {
		return
	}

	p := tea.NewProgram(InitialModel(branches))

	m, err := p.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if quittingError := m.(model).quittingError; quittingError != nil {
		fmt.Fprintln(os.Stderr, *quittingError)
		os.Exit(1)
	}
}
