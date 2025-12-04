package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentCommit(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	createTestHistory(t, testGit)

	currentCommit, err := testGit.CurrentCommit()

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch\n", currentCommit.Message)
}
