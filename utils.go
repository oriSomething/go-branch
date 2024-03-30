package main

import (
	"os/exec"
	"strings"
)

func GetBranches(max int) ([]string, error) {
	cmd := exec.Command("git", "branch", "--sort=-committerdate")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	str := string(stdout)
	branchesStdout := strings.SplitN(str, "\n", max+2)

	size := len(branchesStdout)
	if size > max {
		size = max
	}

	j := 0
	branches := make([]string, size)

	for i := 0; i < size+1 && j < size; i++ {
		// Don't count master branch
		if branchesStdout[i][0] != '*' {
			branches[j] = strings.Trim(branchesStdout[i], " ")
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
