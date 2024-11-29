package github_service

import github_dao "github.com/saphalpdyl/gcms/internals/common/dao/github"

type IGithubService interface {
	CreateNewRepository(string) (*github_dao.CreateNewRepositoryResponseDAO, error)
	DeleteRepository(string) error
	LinkLocalToRemote(string, string, string)
}
