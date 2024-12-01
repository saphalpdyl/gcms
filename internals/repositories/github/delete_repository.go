package github

import (
	"fmt"
	"net/http"
)

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
