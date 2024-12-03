package handlers

import (
	"fmt"
	"os"

	"github.com/saphalpdyl/gcms/helpers"
)

type DeleteLocalHandlerParams struct {
	RepositoryFolderPath string
}

func (h *Handler) DeleteLocal(params DeleteLocalHandlerParams) {
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
