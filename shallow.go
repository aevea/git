package git

// IsShallow returns true if the repository is a shallow clone.
// Shallow clones (e.g. created with git clone --depth=1) have limited
// history and may not have refs like tags available.
func (g *Git) IsShallow() (bool, error) {
	output, err := g.runGitCommand("rev-parse", "--is-shallow-repository")
	if err != nil {
		return false, err
	}
	return output == "true", nil
}
