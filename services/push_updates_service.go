package services

import (
	"errors"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func PushUpdatesGit() error {
	repo, err := git.PlainOpen(".git")
	if err != nil {
		return err
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	pubKeys, err := ssh.NewPublicKeysFromFile("git", path.Join(homedir, ".ssh", "id_ed25519"), "hanif")
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       pubKeys,
	})

	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil
	}

	return err
}
