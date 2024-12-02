package handlers

import (
	github_service "github.com/saphalpdyl/gcms/internals/services/github"
)

type IHandler interface {
	ConfigSet(ConfigSetHandlerParams)
	ConfigGet(ConfigGetHandlerParams)
	Detach(DetachHandlerParams)
	Remove(RemoveHandlerParams)
	// Init(bool)
}

type Handler struct {
	githubService github_service.IGithubService
}

func NewHandler(service github_service.IGithubService) *Handler {
	return &Handler{
		githubService: service,
	}
}
