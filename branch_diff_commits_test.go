package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBranchDiffCommits(t *testing.T) {
	tmpDir, testGit := setupRepo(t)
	defer os.RemoveAll(tmpDir)

	createTestHistory(t, testGit)

	commits, err := testGit.BranchDiffCommits("my-branch", "master")

	commit, _ := testGit.Commit(commits[0])

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch\n", commit.Message)
	assert.Equal(t, 3, len(commits))
}

func TestBranchDiffCommitsWithMasterMerge(t *testing.T) {
	testGit, err := OpenGit("./testdata/commits_on_branch")
	assert.NoError(t, err)

	commits, err := testGit.BranchDiffCommits("behind-master", "origin/master")

	assert.Equal(t, 2, len(commits))

	commit, _ := testGit.Commit(commits[1])

	assert.Equal(t, "chore: commit on branch\n", commit.Message)

	assert.Equal(t, err, nil)

}
