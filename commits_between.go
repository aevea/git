package git

import (
	"strings"
)

// CommitsBetween returns a slice of commit hashes between two commits
func (g *Git) CommitsBetween(from Hash, to Hash) ([]Hash, error) {
	// If from and to are equal, return empty slice
	if from == to {
		return []Hash{}, nil
	}

	fromStr := from.String()
	toStr := to.String()

	// Check if 'to' is an empty hash (all zeros)
	var emptyHash Hash
	if to == emptyHash {
		// If 'to' is empty, return all commits from 'from'
		output, err := g.runGitCommand("log", "--format=%H", fromStr)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(output, "\n")
		var commits []Hash
		for _, line := range lines {
			if line == "" {
				continue
			}
			hash, err := NewHash(line)
			if err != nil {
				continue
			}
			commits = append(commits, hash)
		}
		return commits, nil
	}

	// Get commits from 'from' to 'to' (excluding 'to')
	// Use ^to to exclude the 'to' commit
	output, err := g.runGitCommand("log", "--format=%H", fromStr, "^"+toStr)
	if err != nil {
		// If the command fails, it might be because there are no commits between
		// Try to check if 'to' is reachable from 'from'
		_, err2 := g.runGitCommand("merge-base", "--is-ancestor", toStr, fromStr)
		if err2 != nil {
			return []Hash{}, nil
		}
		return nil, err
	}

	lines := strings.Split(output, "\n")
	var commits []Hash

	for _, line := range lines {
		if line == "" {
			continue
		}
		hash, err := NewHash(line)
		if err != nil {
			continue
		}
		commits = append(commits, hash)
	}

	return commits, nil
}
