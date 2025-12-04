package git

import (
	"fmt"
	"strings"
)

// BranchDiffCommits compares commits from 2 branches and returns of a diff of them.
// Uses git log with exclusion syntax for efficient comparison - finds commits in branchA that are not in branchB.
// This is more efficient than fetching all commits from both branches and comparing them.
func (g *Git) BranchDiffCommits(branchA string, branchB string) ([]Hash, error) {
	// git log branchA ^branchB shows all commits reachable from branchA but not from branchB
	// This is equivalent to: commits in branchA that are not in branchB
	// The ^branchB syntax excludes all commits reachable from branchB
	output, err := g.runGitCommand("log", "--format=%H", branchA, "^"+branchB)
	if err != nil {
		return nil, fmt.Errorf("Failed comparing branches %v and %v: %v", branchA, branchB, err)
	}

	var diffCommits []Hash
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		hash, err := NewHash(line)
		if err != nil {
			// Skip invalid hashes
			continue
		}

		diffCommits = append(diffCommits, hash)
	}

	return diffCommits, nil
}
