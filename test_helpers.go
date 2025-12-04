package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// setupRepo creates a temporary git repository for testing
func setupRepo(t *testing.T) (string, *Git) {
	tmpDir, err := os.MkdirTemp("", "git-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Initialize git repo
	cmd := exec.Command("git", "init", tmpDir)
	if err := cmd.Run(); err != nil {
		os.RemoveAll(tmpDir)
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Configure git
	exec.Command("git", "-C", tmpDir, "config", "user.name", "Test User").Run()
	exec.Command("git", "-C", tmpDir, "config", "user.email", "test@example.com").Run()

	git, err := OpenGit(tmpDir)
	if err != nil {
		os.RemoveAll(tmpDir)
		t.Fatalf("Failed to open git repo: %v", err)
	}

	return tmpDir, git
}

// createCommit creates a commit in the test repository
func createCommit(t *testing.T, git *Git, message string) Hash {
	// Create a dummy file to commit
	testFile := filepath.Join(git.Path, "test.txt")
	err := os.WriteFile(testFile, []byte(message), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Add and commit
	exec.Command("git", "-C", git.Path, "add", "test.txt").Run()
	cmd := exec.Command("git", "-C", git.Path, "commit", "-m", message)
	cmd.Env = append(os.Environ(), "GIT_AUTHOR_DATE=2000-01-01T00:00:00Z", "GIT_COMMITTER_DATE=2000-01-01T00:00:00Z")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to create commit: %v", err)
	}

	hashStr, err := git.runGitCommand("rev-parse", "HEAD")
	if err != nil {
		t.Fatalf("Failed to get commit hash: %v", err)
	}

	hash, err := NewHash(hashStr)
	if err != nil {
		t.Fatalf("Failed to parse commit hash: %v", err)
	}

	return hash
}

// createBranch creates a branch in the test repository
func createBranch(t *testing.T, git *Git, branchName string) {
	cmd := exec.Command("git", "-C", git.Path, "checkout", "-b", branchName)
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to create branch: %v", err)
	}
}

// createTestHistory creates a test history with commits and branches
func createTestHistory(t *testing.T, git *Git) Hash {
	// Create initial commit on master
	createCommit(t, git, "test commit on master")

	// Create branch
	createBranch(t, git, "my-branch")

	// Create commits on branch
	createCommit(t, git, "commit on new branch")
	createCommit(t, git, "second commit on new branch\n\n Long message")
	lastHash := createCommit(t, git, "third commit on new branch")

	return lastHash
}
