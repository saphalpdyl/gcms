package github

import (
	"bytes"
	"log"
	"net/http"

	"github.com/saphalpdyl/gcms/utils"
)

type GithubRepositoryImpl struct {
	PATToken             string
	RepositoryFolderPath string
}

func NewRepository(patToken string, repositoryFolderPath string) IGithubRepository {
	return &GithubRepositoryImpl{
		PATToken:             patToken,
		RepositoryFolderPath: repositoryFolderPath,
	}
}

func prepareRequest(url string, requestType string, payload []byte, patToken string) *http.Request {
	// Create a new HTTP request
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "token "+patToken)
	req.Header.Set("Content-Type", "application/json")

	return req
}

func (r *GithubRepositoryImpl) validateRepositoryExistence() {
	if !utils.PathExists(r.RepositoryFolderPath) {
		log.Fatal("fatal Repository validation failed:  ")
	}
}
