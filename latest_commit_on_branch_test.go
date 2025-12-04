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
