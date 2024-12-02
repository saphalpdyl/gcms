package handlers

import (
	"fmt"
	"os"

	"github.com/saphalpdyl/gcms/helpers"
)

type RemoveHandlerParams struct {
	RepositoryFolderPath string
}

func (h *Handler) RemoveHandler(params RemoveHandlerParams) {
	var deleteConfirmationAnswer string

	// Confirmation message
	fmt.Printf(
		"\n%s\n%s\n%s\n",
		helpers.RenderDiff("Deleting the local repository", false, ""),
		"This action is irreversible.",
		helpers.RenderBold("Are you sure you want to continue?[y/N] "),
	)
	fmt.Scan(&deleteConfirmationAnswer)

	if deleteConfirmationAnswer != "y" && deleteConfirmationAnswer != "Y" {
		return
	}

	os.RemoveAll(params.RepositoryFolderPath)

}
