package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentTagHappy(t *testing.T) {
	testGit, err := OpenGit("./testdata/git_tags")

	assert.NoError(t, err)

	tag, err := testGit.CurrentTag()

	assert.NoError(t, err)
	assert.Equal(t, "refs/tags/v0.0.2", tag.Name)
}

func TestCurrentTagAnnotatedHappy(t *testing.T) {
	testGit, err := OpenGit("./testdata/annotated_git_tags_mix")

	assert.NoError(t, err)

	tag, err := testGit.CurrentTag()

	assert.NoError(t, err)
	assert.Equal(t, "v0.0.3", tag.Name)
}

func TestCurrentTagUnhappy(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo}

	_, err := testGit.CurrentTag()

	assert.Equal(t, ErrCommitNotOnTag, err)
}
