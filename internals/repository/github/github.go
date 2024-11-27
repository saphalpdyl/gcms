package github

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func CreateNewRepository(repoName string) error {
	verifyInitialization()

	// Create the request body
	data := map[string]string{
		"name": repoName,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	// Create a new HTTP request
	url := "https://api.github.com/user/repos"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", "token "+instance.PATToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != 201 {
		return fmt.Errorf("%s", resp.Status)
	}

	return err
}
