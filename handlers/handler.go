package handlers

import (
	"github.com/saphalpdyl/gcms/internals/repositories/github"
	"github.com/saphalpdyl/gcms/internals/repositories/schema"
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
	SchemaCreateNewGroup(SchemaCreateNewHandlerParams)
}

type Handler struct {
	githubRepostiory github.IGithubRepository
	schemaRepository schema.ISchemaRepository
}

func NewHandler(githubRepository github.IGithubRepository, schemaRepository schema.ISchemaRepository) *Handler {
	return &Handler{
		githubRepostiory: githubRepository,
		schemaRepository: schemaRepository,
	}
}
