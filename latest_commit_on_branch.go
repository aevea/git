package git

// LatestCommitOnBranch resolves a revision and then returns the latest commit on it.
func (g *Git) LatestCommitOnBranch(desiredBranch string) (*Commit, error) {
	hashStr, err := g.runGitCommand("rev-parse", desiredBranch)
	if err != nil {
		return nil, err
	}

	hash, err := NewHash(hashStr)
	if err != nil {
		return nil, err
	}

	return g.Commit(hash)
}
