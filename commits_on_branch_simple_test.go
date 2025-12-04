package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitsOnBranchSimple(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	lastHash := createTestHistory(t, testGit)

	commits, err := testGit.CommitsOnBranchSimple(lastHash)

	assert.Equal(t, 4, len(commits))

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch", commits[0].Message)
	assert.Equal(t, err, nil)

	assert.NoError(t, err)
	assert.Equal(t, "second commit on new branch\n\n Long message", commits[1].Message)
	assert.Equal(t, err, nil)
}
