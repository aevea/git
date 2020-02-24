package git

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommit(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo, DebugLogger: log.New(ioutil.Discard, "", 0)}

	head, _ := testGit.CurrentCommit()

	commit, err := testGit.Commit(head.Hash)
	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch", commit.Message)
	assert.NotEmpty(t, commit.Hash)
}