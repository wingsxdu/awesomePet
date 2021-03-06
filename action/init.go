package action

import (
	"awesomePet/api"
	. "awesomePet/models"
	"os"
)

func init() {
	err := os.MkdirAll(OriginalPPPath, os.ModePerm) // mkdir
	api.PanicErr(err)
	err = os.MkdirAll(ThumbnailPPPath, os.ModePerm) // mkdir
	api.PanicErr(err)
}
