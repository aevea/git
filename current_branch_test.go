package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentBranch(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	createTestHistory(t, testGit)

	currentBranch, err := testGit.CurrentBranch()

	assert.NoError(t, err)
	assert.Equal(t, "refs/heads/my-branch", currentBranch.Name())
}

func TestCurrentBranchDetachedHead(t *testing.T) {
	testGit, err := OpenGit("./testdata/detached_head")
	assert.NoError(t, err)

	currentBranch, err := testGit.CurrentBranch()

	assert.NoError(t, err)
	assert.Equal(t, "HEAD", currentBranch.Name())
	
	// Verify it has a valid commit hash
	assert.NotEmpty(t, currentBranch.Hash().String())
}
