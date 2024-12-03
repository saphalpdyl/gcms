package helpers

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func CommitCurrentChanges(repositoryFolderPath string, commitMsg string) error {
	// Commit the changes
	g, _ := git.PlainOpen(repositoryFolderPath)
	w, _ := g.Worktree()

	_, err := w.Add(".")
	if err != nil {
		return fmt.Errorf("fatal couldn't stage all files in git")
	}

	_, err = w.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "GCMS Service Worker",
			Email: "worer",
			When:  time.Now(),
		},
	})

	if err != nil {
		return fmt.Errorf("fatal couldn't commit changes")
	}

	return nil
}
