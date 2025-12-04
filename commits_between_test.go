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

func TestCommitsBetweenDetachedHead(t *testing.T) {
	testGit, err := OpenGit("./testdata/detached_head")
	assert.NoError(t, err)

	// Verify we're in detached HEAD state
	currentBranch, err := testGit.CurrentBranch()
	assert.NoError(t, err)
	assert.Equal(t, "HEAD", currentBranch.Name())

	// Get HEAD commit hash (second commit)
	headHash := currentBranch.Hash()

	// Get origin/master commit hash (third commit)
	masterHashStr, err := testGit.runGitCommand("rev-parse", "origin/master")
	assert.NoError(t, err)
	masterHash, err := NewHash(masterHashStr)
	assert.NoError(t, err)

	// CommitsBetween should work even in detached HEAD state
	// Get commits between HEAD (second commit) and origin/master (third commit)
	commits, err := testGit.CommitsBetween(masterHash, headHash)

	assert.NoError(t, err)
	// Should have 1 commit (the third commit)
	assert.Equal(t, 1, len(commits))

	commit, err := testGit.Commit(commits[0])
	assert.NoError(t, err)
	assert.Equal(t, "third commit\n", commit.Message)
}
