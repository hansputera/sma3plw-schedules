package services

import (
	"errors"
	"os"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func PushUpdatesGit() error {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	worktree.AddGlob(".")

	c, err := worktree.Commit("feat(auto): push updates", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Hanif Dwy Putra S",
			Email: "hanifdwyputrasembiring@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	repo.CommitObject(c)

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	pubKeys, err := ssh.NewPublicKeysFromFile("git", path.Join(homedir, ".ssh", "id_ed25519"), "hanif")
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{
		Auth:       pubKeys,
	})

	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil
	}

	return err
}
