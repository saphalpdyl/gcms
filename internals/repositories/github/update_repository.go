package github

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func (r *GithubRepositoryImpl) UpdateRepository() {
	r.validateRepositoryExistence()

	// Pull from remote
	g, err := git.PlainOpen(r.RepositoryFolderPath)

	if err != nil {
		log.Fatal("fatal couldn't open repository: ", err)
	}

	w, err := g.Worktree()
	if err != nil {
		log.Fatal("fatal couldn't open working tree: ", err)
	}

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})

	if err != nil {
		log.Fatal("fatal couldn't pull from remote: ", err)
	}

	// Push
	err = g.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "Personal Access Token",
			Password: r.PATToken,
		},
	})

	if err != nil {
		log.Fatal("fatal couldn't push to repository: ", err)
	}
}
