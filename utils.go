package main

import (
	"os/exec"
	"strings"
)

func GetBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--sort=-committerdate")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	str := string(stdout)
	branchesStdout := strings.Split(str, "\n")

	// -1 for current branch
	// -1 for empty line in the end
	branchesStdoutSize := len(branchesStdout) - 2
	size := branchesStdoutSize

	// 1 because "current branch" is something we don't show
	if size == 0 {
		return make([]string, 0), nil
	}

	branches := make([]string, size)
	for i, j := 0, 0; i < size+1 && j < size; i++ {
		// Don't count the curret branch
		if branchesStdout[i][0] != '*' {
			branches[j] = strings.TrimSpace(branchesStdout[i])
			j++
		}
	}

	return branches, nil
}

func SwitchBranch(branch string) error {
	cmd := exec.Command("git", "switch", branch)

	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}
