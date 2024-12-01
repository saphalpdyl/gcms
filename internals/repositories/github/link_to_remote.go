package github

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func (r *GithubRepositoryImpl) LinkLocalToRemote(path string, repoName string, ownerName string) {

	g, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal("fatal couldn't open repository: ", err)
	}

	remoteURL := fmt.Sprintf("https://%s:%s@github.com/%s/%s.git", ownerName, r.PATToken, ownerName, repoName)

	_, err = g.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteURL},
	})

	if err != nil {
		log.Fatal("fatal couldn't create remote: ", err)
	}

	// Fetch just once
	g.Fetch(&git.FetchOptions{
		RemoteName: "origin",
	})
}
