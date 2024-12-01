package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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
