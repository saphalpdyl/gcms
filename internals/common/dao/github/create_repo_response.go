package github_dao

import "github.com/saphalpdyl/gcms/internals/repositories/github"

// Data Access Objects
type CreateNewRepositoryResponseDAO struct {
	URL                string
	RepositoryName     string
	RepositoryFullName string
	RepositoryOwner    struct {
		RepositoryOwnerName string
	}
}

func CreateNewRepositoryResponseDAOFrom(dto_r *github.CreateNewRepositoryResponse) CreateNewRepositoryResponseDAO {
	newResponse := CreateNewRepositoryResponseDAO{
		URL:                dto_r.URL,
		RepositoryName:     dto_r.RepositoryName,
		RepositoryFullName: dto_r.RepositoryFullName,
	}

	newResponse.RepositoryOwner.RepositoryOwnerName = dto_r.RepositoryOwner.RepositoryOwnerName

	return newResponse
}
