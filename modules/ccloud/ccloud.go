package ccloud

import (
	"testing"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"io/ioutil"
)

func ValidateRepository(t *testing.T, url string, commit string) (string, error) {

	dir, _ := ioutil.TempDir("", "terratest")

	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		return "", err
	}

	w, err := r.Worktree()


	// ... checking out to commit
	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(commit),
	})

	if err != nil {
		return "", err
	}
	ref, err := r.Head()

	return ref.Name().String(), nil
}