package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatestCommitOnBranch(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	createTestHistory(t, testGit)

	commit, err := testGit.LatestCommitOnBranch("my-branch")

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch\n", commit.Message)
	assert.Equal(t, err, nil)
}

func TestLatestCommitOnBranchDetachedHead(t *testing.T) {
	testGit, err := OpenGit("./testdata/detached_head")
	assert.NoError(t, err)

	// Verify we're in detached HEAD state
	currentBranch, err := testGit.CurrentBranch()
	assert.NoError(t, err)
	assert.Equal(t, "HEAD", currentBranch.Name())

	// LatestCommitOnBranch should work even in detached HEAD state
	commit, err := testGit.LatestCommitOnBranch("origin/master")

	assert.NoError(t, err)
	assert.Equal(t, "third commit\n", commit.Message)
}
