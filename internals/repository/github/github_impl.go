package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
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

func (r *GithubRepositoryImpl) CreateNewRepository(repoName string) (*CreateNewRepositoryResponse, error) {

	// Create the request body
	data := map[string]string{
		"name": repoName,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	url := "https://api.github.com/user/repos"
	req := prepareRequest(url, "POST", jsonData, r.PATToken)

	// Make the request
	client := &http.Client{}
	resp, _ := client.Do(req)

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	defer resp.Body.Close()

	var response *CreateNewRepositoryResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("fatal couldn't unmarshal repository response", err)
	}

	return response, nil
}

func (r *GithubRepositoryImpl) DeleteRepository(url string) error {

	req := prepareRequest(url, "DELETE", []byte{}, r.PATToken)

	// Make the request
	client := &http.Client{}
	resp, _ := client.Do(req)

	if resp.StatusCode != 204 {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}

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
