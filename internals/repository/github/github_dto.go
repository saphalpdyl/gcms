package github

// Create Response Data Transfer Object(DTO)
type CreateNewRepositoryResponse struct {
	URL                string `json:"html_url"`
	RepositoryName     string `json:"name"`
	RepositoryFullName string `json:"full_name"`
	RepositoryOwner    struct {
		RepositoryOwnerName string `json:"login"`
	} `json:"owner"`
}
