package handlers

import (
	"github.com/saphalpdyl/gcms/internals/repositories/github"
)

type IHandler interface {
	ConfigSet(ConfigSetHandlerParams)
	ConfigGet(ConfigGetHandlerParams)
	Detach(DetachHandlerParams)
	DeleteLocal(DeleteLocalHandlerParams)
	Init(InitHandlerParams)
	Update(UpdateHandlerParams)
	Doctor(DoctorHandlerParams)
	Push(PushHandlerParams)
	Remove(RemoveHandlerParams)
	List(ListHandlerParams)

	// Schema Specific
	SchemaCreateNew(SchemaCreateNewHandlerParams)
}

type Handler struct {
	githubRepostiory github.IGithubRepository
}

func NewHandler(githubRepository github.IGithubRepository) *Handler {
	return &Handler{
		githubRepostiory: githubRepository,
	}
}
