package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type githubRepository struct {
	PATToken string
}

var instance *githubRepository = nil

func Initiailize(patToken string) {
	instance = &githubRepository{
		PATToken: patToken,
	}
}

func verifyInitialization() {
	if instance == nil {
		log.Fatalf("fatal GithubRepository has not been initialized yet with New(). Maybe the remote has not been configured yet?")
	}
}

type CreateNewRepositoryResponse struct {
	URL string `json:"html_url"`
}

func prepareRequest(url string, requestType string, payload []byte) *http.Request {
	// Create a new HTTP request
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "token "+instance.PATToken)
	req.Header.Set("Content-Type", "application/json")

	return req
}

func CreateNewRepository(repoName string) (string, error) {
	verifyInitialization()

	// Create the request body
	data := map[string]string{
		"name": repoName,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	url := "https://api.github.com/user/repos"
	req := prepareRequest(url, "POST", jsonData)

	// Make the request
	client := &http.Client{}
	resp, _ := client.Do(req)

	if resp.StatusCode != 201 {
		return "", fmt.Errorf("%s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	defer resp.Body.Close()

	var response CreateNewRepositoryResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("fatal couldn't unmarshal repository response", err)
	}

	return response.URL, nil
}
