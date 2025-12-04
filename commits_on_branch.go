package git

import (
	"github.com/apex/log"
	"strings"
)

// CommitsOnBranch iterates through all references and returns commit hashes on given branch. \n
// Important to note is that this will provide all commits from anything the branch is connected to.
func (g *Git) CommitsOnBranch(
	branchCommit Hash,
) ([]Hash, error) {
	hashStr := branchCommit.String()

	// Get all commit hashes
	output, err := g.runGitCommand("log", "--format=%H", hashStr)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(output, "\n")
	var branchCommits []Hash

	for _, line := range lines {
		if line == "" {
			continue
		}
		hash, err := NewHash(line)
		if err != nil {
			log.Debugf("Failed to parse hash %s: %v", line, err)
			continue
		}
		branchCommits = append(branchCommits, hash)
	}

	return branchCommits, nil
}
