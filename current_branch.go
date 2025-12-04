package git

// CurrentBranch returns the reference HEAD is at right now
func (g *Git) CurrentBranch() (*Reference, error) {
	// Get the symbolic ref name
	refName, err := g.runGitCommand("symbolic-ref", "HEAD")
	if err != nil {
		return nil, err
	}

	// Get the commit hash
	hashStr, err := g.runGitCommand("rev-parse", "HEAD")
	if err != nil {
		return nil, err
	}

	hash, err := NewHash(hashStr)
	if err != nil {
		return nil, err
	}

	return NewReference(refName, hash), nil
}
