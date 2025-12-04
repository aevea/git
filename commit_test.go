package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommit(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	createTestHistory(t, testGit)

	head, _ := testGit.CurrentCommit()

	commit, err := testGit.Commit(head.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch\n", commit.Message)
	assert.NotEmpty(t, commit.Hash)
}
