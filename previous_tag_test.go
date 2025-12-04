package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreviousTag(t *testing.T) {
	testGit, err := OpenGit("./testdata/git_tags")
	assert.NoError(t, err)

	headHashStr, err := testGit.runGitCommand("rev-parse", "HEAD")
	assert.NoError(t, err)

	headHash, err := NewHash(headHashStr)
	assert.NoError(t, err)

	tag, err := testGit.PreviousTag(headHash)

	assert.NoError(t, err)

	commit, err := testGit.Commit(tag.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "chore: first commit on master\n", commit.Message)

}
