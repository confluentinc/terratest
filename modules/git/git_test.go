package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func TestGetCurrentBranchName(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dir, _ := filepath.Abs(filepath.Dir(filename))
	dir = filepath.Dir(filepath.Dir(dir))
	name, _ := GetCurrentBranchName(t, dir)
	assert.NotEmpty(t, name)
}

func TestCloneRepository(t *testing.T) {
	t.Parallel()
	dir, _ := ioutil.TempDir("", "terratesttemp")
	name, err := CloneRepository(t, "https://github.com/githubtraining/hellogitworld.git", dir, "master")
	assert.Empty(t, err)
	assert.Equal(t, "refs/heads/master", name)
}

func TestCheckoutCommitName(t *testing.T) {
	t.Parallel()
	dir, _ := ioutil.TempDir("", "terratesttemp")
	//Clone repository
	CloneRepository(t, "https://github.com/githubtraining/hellogitworld.git", dir, "master")
	//Test checking out
	name, err := CheckoutCommitName(t, dir, "gh-pages")
	assert.Empty(t, err)
	assert.Equal(t, "refs/heads/master", name)
}
