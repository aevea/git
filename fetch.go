package git

import (
	"fmt"
)

// FetchTags fetches all tags from the specified remote.
// This is useful for shallow clones where tags are not available by default.
// Pass an empty string for remote to use "origin".
func (g *Git) FetchTags(remote string) error {
	if remote == "" {
		remote = "origin"
	}

	_, err := g.runGitCommand("fetch", remote, "--tags", "--force")
	if err != nil {
		return fmt.Errorf("failed to fetch tags from %s: %w", remote, err)
	}

	return nil
}

// Unshallow converts a shallow clone into a full clone by fetching the
// complete history. This is useful when operations require full commit
// history that is not available in shallow clones.
// Pass an empty string for remote to use "origin".
func (g *Git) Unshallow(remote string) error {
	if remote == "" {
		remote = "origin"
	}

	shallow, err := g.IsShallow()
	if err != nil {
		return fmt.Errorf("failed to check if repository is shallow: %w", err)
	}

	if !shallow {
		return nil
	}

	_, err = g.runGitCommand("fetch", remote, "--unshallow")
	if err != nil {
		return fmt.Errorf("failed to unshallow repository from %s: %w", remote, err)
	}

	return nil
}
