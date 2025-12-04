package git

import (
	"strconv"
	"strings"
	"time"
)

// Commit finds a commit based on commit hash and returns the Commit object
func (g *Git) Commit(hash Hash) (*Commit, error) {
	hashStr := hash.String()

	// Get commit message - use %B to get full body, then normalize
	// go-git's Message field returns the raw commit message which always ends with \n
	fullMessage, err := g.runGitCommand("log", "-1", "--format=%B", hashStr)
	if err != nil {
		return nil, err
	}
	
	// Normalize: remove trailing newlines and add exactly one
	// This matches go-git's behavior where Message always ends with \n
	message := strings.TrimRight(fullMessage, "\n")
	if message != "" {
		message += "\n"
	}

	// Get author info
	authorName, err := g.runGitCommand("log", "-1", "--format=%an", hashStr)
	if err != nil {
		return nil, err
	}

	authorEmail, err := g.runGitCommand("log", "-1", "--format=%ae", hashStr)
	if err != nil {
		return nil, err
	}

	authorDateStr, err := g.runGitCommand("log", "-1", "--format=%at", hashStr)
	if err != nil {
		return nil, err
	}

	authorTimestamp, err := strconv.ParseInt(authorDateStr, 10, 64)
	if err != nil {
		return nil, err
	}

	// Get committer info
	committerName, err := g.runGitCommand("log", "-1", "--format=%cn", hashStr)
	if err != nil {
		return nil, err
	}

	committerEmail, err := g.runGitCommand("log", "-1", "--format=%ce", hashStr)
	if err != nil {
		return nil, err
	}

	committerDateStr, err := g.runGitCommand("log", "-1", "--format=%ct", hashStr)
	if err != nil {
		return nil, err
	}

	committerTimestamp, err := strconv.ParseInt(committerDateStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &Commit{
		Hash:    hash,
		Message: message,
		Author: Signature{
			Name:  authorName,
			Email: authorEmail,
			When:  time.Unix(authorTimestamp, 0),
		},
		Committer: Signature{
			Name:  committerName,
			Email: committerEmail,
			When:  time.Unix(committerTimestamp, 0),
		},
	}, nil
}
