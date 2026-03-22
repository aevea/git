package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsShallowFalse(t *testing.T) {
	testGit, err := OpenGit("./testdata/git_tags")
	require.NoError(t, err)

	shallow, err := testGit.IsShallow()
	assert.NoError(t, err)
	assert.False(t, shallow)
}

func TestIsShallowTrue(t *testing.T) {
	// Create a shallow clone of an existing test repo
	tmpDir, err := os.MkdirTemp("", "shallow-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	absTagsPath, err := filepath.Abs("./testdata/git_tags")
	require.NoError(t, err)

	// Use file:// protocol so --depth is respected for local repos
	cmd := exec.Command("git", "clone", "--depth=1", "file://"+absTagsPath, tmpDir)
	require.NoError(t, cmd.Run())

	testGit, err := OpenGit(tmpDir)
	require.NoError(t, err)

	shallow, err := testGit.IsShallow()
	assert.NoError(t, err)
	assert.True(t, shallow)
}

func TestShallowCloneMissingTags(t *testing.T) {
	// Demonstrate that a shallow clone is missing tags
	tmpDir, err := os.MkdirTemp("", "shallow-tags-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	absTagsPath, err := filepath.Abs("./testdata/git_tags")
	require.NoError(t, err)

	// Clone with depth=1 and no tags using file:// protocol
	cmd := exec.Command("git", "clone", "--depth=1", "--no-tags", "file://"+absTagsPath, tmpDir)
	require.NoError(t, cmd.Run())

	testGit, err := OpenGit(tmpDir)
	require.NoError(t, err)

	// Verify the clone is shallow
	shallow, err := testGit.IsShallow()
	require.NoError(t, err)
	assert.True(t, shallow)

	// Tags should be empty in this shallow clone
	tags, err := testGit.getTags()
	assert.NoError(t, err)
	assert.Empty(t, tags)

	// After fetching tags, they should be available
	err = testGit.FetchTags("origin")
	assert.NoError(t, err)

	tags, err = testGit.getTags()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(tags))
}

func TestFetchTagsDefaultRemote(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "fetch-tags-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	absTagsPath, err := filepath.Abs("./testdata/git_tags")
	require.NoError(t, err)

	cmd := exec.Command("git", "clone", "--depth=1", "--no-tags", "file://"+absTagsPath, tmpDir)
	require.NoError(t, cmd.Run())

	testGit, err := OpenGit(tmpDir)
	require.NoError(t, err)

	// Empty string should default to "origin"
	err = testGit.FetchTags("")
	assert.NoError(t, err)

	tags, err := testGit.getTags()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(tags))
}

func TestUnshallow(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "unshallow-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	absTagsPath, err := filepath.Abs("./testdata/git_tags")
	require.NoError(t, err)

	cmd := exec.Command("git", "clone", "--depth=1", "file://"+absTagsPath, tmpDir)
	require.NoError(t, cmd.Run())

	testGit, err := OpenGit(tmpDir)
	require.NoError(t, err)

	shallow, err := testGit.IsShallow()
	require.NoError(t, err)
	assert.True(t, shallow)

	err = testGit.Unshallow("")
	assert.NoError(t, err)

	// Should no longer be shallow
	shallow, err = testGit.IsShallow()
	assert.NoError(t, err)
	assert.False(t, shallow)
}

func TestUnshallowNoop(t *testing.T) {
	// Unshallow on a non-shallow repo should be a no-op
	testGit, err := OpenGit("./testdata/git_tags")
	require.NoError(t, err)

	err = testGit.Unshallow("")
	assert.NoError(t, err)
}
