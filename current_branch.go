package git

// CurrentBranch returns the reference HEAD is at right now.
// In detached HEAD state, it returns a reference with name "HEAD".
func (g *Git) CurrentBranch() (*Reference, error) {
	// Get the commit hash first (works in both normal and detached HEAD state)
	hashStr, err := g.runGitCommand("rev-parse", "HEAD")
	if err != nil {
		return nil, err
	}

	hash, err := NewHash(hashStr)
	if err != nil {
		return nil, err
	}

	// Try to get the symbolic ref name (fails in detached HEAD state)
	refName, err := g.runGitCommand("symbolic-ref", "HEAD")
	if err != nil {
		// In detached HEAD state, use "HEAD" as the reference name
		refName = "HEAD"
	}

	return NewReference(refName, hash), nil
}
