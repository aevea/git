package git

import (
	"github.com/apex/log"
)

// CurrentCommit returns the commit that HEAD is at
func (g *Git) CurrentCommit() (*Commit, error) {
	hashStr, err := g.runGitCommand("rev-parse", "HEAD")
	if err != nil {
		return nil, err
	}

	log.Debugf("current commitHash: %v \n", hashStr)

	return g.Commit(MustHash(hashStr))
}
