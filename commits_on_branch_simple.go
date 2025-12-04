package git

import (
	"strings"

	"github.com/apex/log"
)

// SimpleCommit is a slimed down commit object of just Hash and Message
type SimpleCommit struct {
	Hash    Hash
	Message string
}

// CommitsOnBranchSimple iterates through all references and returns simpleCommits on given branch. \n
// Important to note is that this will provide all commits from anything the branch is connected to.
func (g *Git) CommitsOnBranchSimple(
	branchCommit Hash,
) ([]SimpleCommit, error) {
	hashStr := branchCommit.String()

	// Get all commit hashes first
	hashOutput, err := g.runGitCommand("log", "--format=%H", hashStr)
	if err != nil {
		return nil, err
	}

	hashLines := strings.Split(hashOutput, "\n")
	var branchCommits []SimpleCommit

	for _, hashLine := range hashLines {
		if hashLine == "" {
			continue
		}
		hash, err := NewHash(hashLine)
		if err != nil {
			log.Debugf("Failed to parse hash %s: %v", hashLine, err)
			continue
		}

		// Get message for this commit
		message, err := g.runGitCommand("log", "-1", "--format=%B", hashLine)
		if err != nil {
			log.Debugf("Failed to get message for commit %s: %v", hashLine, err)
			continue
		}

		branchCommits = append(branchCommits, SimpleCommit{
			Hash:    hash,
			Message: strings.TrimSpace(message),
		})
	}

	return branchCommits, nil
}
