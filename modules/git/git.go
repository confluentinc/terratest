// Package git allows to interact with Git.
package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"testing"
)

func CloneRepository(t *testing.T, url string, path string, commit string) (string, error) {

	r, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		return "", err
	}

	w, err := r.Worktree()

	if err != nil {
		return "", err
	}

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

// GetCurrentBranchName retrieves the current branch name.
func GetCurrentBranchName(t *testing.T, path string) (string, error) {

	// We instance a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)

	if err != nil {
		return "", err
	}

	// ... retrieving the HEAD reference
	ref, err := r.Head()
	return ref.Name().String(), nil

}

// CheckoutCommitName checksout commit name.
func CheckoutCommitName(t *testing.T, path string, commit string) (string, error) {

	// We instance a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)

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
