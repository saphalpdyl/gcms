package handlers

type UpdateHandlerParams struct {
}

func (h *Handler) Update(params UpdateHandlerParams) {
	h.githubRepostiory.UpdateRepository()
}
