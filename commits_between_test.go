package git

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/stretchr/testify/assert"
)

func TestCommitsBetween(t *testing.T) {
	repo, _ := git.PlainOpen("./testdata/git_tags")
	testGit := &Git{repo: repo}

	head, err := repo.Head()

	assert.NoError(t, err)

	tag, err := testGit.PreviousTag(head.Hash())

	assert.NoError(t, err)

	commit, err := repo.CommitObject(tag.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "chore: first commit on master\n", commit.Message)

	commits, err := testGit.CommitsBetween(head.Hash(), tag.Hash)

	assert.NoError(t, err)
	assert.Len(t, commits, 3)

	middleCommit, _ := repo.CommitObject(commits[1])

	assert.Equal(t, "chore: third commit on master\n", middleCommit.Message)
}

func TestNoToCommit(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	head, _ := repo.Head()

	testGit := &Git{repo: repo}

	commits, err := testGit.CommitsBetween(head.Hash(), plumbing.Hash{})

	assert.Equal(t, 4, len(commits))

	commit, commitErr := repo.CommitObject(commits[0])

	assert.NoError(t, commitErr)
	assert.Equal(t, "third commit on new branch", commit.Message)
	assert.Equal(t, err, nil)

	lastCommit, _ := repo.CommitObject(commits[3])

	assert.Equal(t, "test commit on master", lastCommit.Message)
}

func TestToFromEqual(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	head, _ := repo.Head()

	testGit := &Git{repo: repo}

	commits, err := testGit.CommitsBetween(head.Hash(), head.Hash())

	assert.Equal(t, 0, len(commits))
	assert.NoError(t, err)
}
