package github

type IGithubRepository interface {
	CreateNewRepository(string) (*CreateNewRepositoryResponse, error)
	DeleteRepository(string) error
	LinkLocalToRemote(string, string, string)
	UpdateRepository()
}
