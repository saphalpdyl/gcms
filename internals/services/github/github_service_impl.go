package github_service

import (
	github_dao "github.com/saphalpdyl/gcms/internals/common/dao/github"
	"github.com/saphalpdyl/gcms/internals/repositories/github"
)

type GithubServiceImpl struct {
	Repository github.IGithubRepository
}

func NewService(repository github.IGithubRepository) IGithubService {
	return &GithubServiceImpl{
		Repository: repository,
	}
}

func (r *GithubServiceImpl) CreateNewRepository(repoName string) (*github_dao.CreateNewRepositoryResponseDAO, error) {
	response, err := r.Repository.CreateNewRepository(repoName)

	if err != nil {
		return nil, err
	}

	responseDAO := github_dao.CreateNewRepositoryResponseDAOFrom(response)

	return &responseDAO, nil
}

func (r *GithubServiceImpl) DeleteRepository(url string) error {
	return r.Repository.DeleteRepository(url)
}

func (r *GithubServiceImpl) LinkLocalToRemote(path string, repoName string, ownerName string) {
	r.Repository.LinkLocalToRemote(path, repoName, ownerName)
}

func (r *GithubServiceImpl) UpdateRepository() {
	r.Repository.UpdateRepository()
}
