package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitsOnBranch(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	lastHash := createTestHistory(t, testGit)

	commits, err := testGit.CommitsOnBranch(lastHash)

	assert.Equal(t, 4, len(commits))

	commit, commitErr := testGit.Commit(commits[0])

	assert.NoError(t, commitErr)
	assert.Equal(t, "third commit on new branch\n", commit.Message)
	assert.Equal(t, err, nil)

	lastCommit, _ := testGit.Commit(commits[3])

	assert.Equal(t, "test commit on master\n", lastCommit.Message)
}

func TestCommitsOnBranchDetachedHead(t *testing.T) {
	testGit, err := OpenGit("./testdata/detached_head")
	assert.NoError(t, err)

	// Verify we're in detached HEAD state
	currentBranch, err := testGit.CurrentBranch()
	assert.NoError(t, err)
	assert.Equal(t, "HEAD", currentBranch.Name())

	// Get HEAD commit hash
	headHash := currentBranch.Hash()

	// CommitsOnBranch should work even in detached HEAD state
	// It takes a commit hash, so HEAD state doesn't matter
	commits, err := testGit.CommitsOnBranch(headHash)

	assert.NoError(t, err)
	// Should have 2 commits (second commit and first commit)
	assert.Equal(t, 2, len(commits))

	commit, err := testGit.Commit(commits[0])
	assert.NoError(t, err)
	assert.Equal(t, "second commit\n", commit.Message)

	lastCommit, err := testGit.Commit(commits[1])
	assert.NoError(t, err)
	assert.Equal(t, "first commit\n", lastCommit.Message)
}
