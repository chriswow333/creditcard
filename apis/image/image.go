package image

import (
	"net/http"
	"path/filepath"

	"example.com/creditcard/middlewares/apis"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type imageHandler struct {
	dig.In
}

const DOWNLOADS_PATH = "static/card_image/"

func NewImageHandler(
	rg *gin.RouterGroup,
) {

	ih := &imageHandler{}

	apis.Handle(rg, http.MethodGet, "/:imageName", ih.downloadImage)
}

func (h *imageHandler) downloadImage(ctx *gin.Context) {
	imageName := ctx.Param("imageName")

	targetPath := filepath.Join(DOWNLOADS_PATH, imageName)
	//This ckeck is for example, I not sure is it can prevent all possible filename attacks - will be much better if real filename will not come from user side. I not even tryed this code
	// if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
	// 	ctx.String(403, "Look like you attacking me")
	// 	return
	// }
	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+imageName)
	ctx.Header("Content-Type", "application/octet-stream")

	ctx.File(targetPath)

}
