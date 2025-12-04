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
