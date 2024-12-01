package github

import (
	"bytes"
	"log"
	"net/http"
)

type GithubRepositoryImpl struct {
	PATToken string
}

func NewRepository(patToken string) IGithubRepository {
	return &GithubRepositoryImpl{
		PATToken: patToken,
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
