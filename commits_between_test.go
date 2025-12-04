package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitsBetween(t *testing.T) {
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

	commits, err := testGit.CommitsBetween(headHash, tag.Hash)

	assert.NoError(t, err)
	assert.Len(t, commits, 3)

	middleCommit, _ := testGit.Commit(commits[1])

	assert.Equal(t, "chore: third commit on master\n", middleCommit.Message)
}

func TestNoToCommit(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	lastHash := createTestHistory(t, testGit)

	var emptyHash Hash
	commits, err := testGit.CommitsBetween(lastHash, emptyHash)

	assert.Equal(t, 4, len(commits))

	commit, commitErr := testGit.Commit(commits[0])

	assert.NoError(t, commitErr)
	assert.Equal(t, "third commit on new branch\n", commit.Message)
	assert.Equal(t, err, nil)

	lastCommit, _ := testGit.Commit(commits[3])

	assert.Equal(t, "test commit on master\n", lastCommit.Message)
}

func TestToFromEqual(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	lastHash := createTestHistory(t, testGit)

	commits, err := testGit.CommitsBetween(lastHash, lastHash)

	assert.Equal(t, 0, len(commits))
	assert.NoError(t, err)
}
