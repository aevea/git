package git

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// Git is the struct used to house all methods in use in Commitsar.
type Git struct {
	Path string
}

// OpenGit loads Repo on path and returns a new Git struct to work with.
func OpenGit(path string) (*Git, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// Verify it's a git repository
	cmd := exec.Command("git", "-C", absPath, "rev-parse", "--git-dir")
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return &Git{Path: absPath}, nil
}

// runGitCommand runs a git command in the repository directory
func (g *Git) runGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", append([]string{"-C", g.Path}, args...)...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
