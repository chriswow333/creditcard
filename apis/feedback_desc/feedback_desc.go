package feedback_desc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"example.com/creditcard/middlewares/apis"

	feedbackDescSrc "example.com/creditcard/service/feedback_desc"
)

type feedbackDescHandler struct {
	dig.In

	feedbackDescSrc feedbackDescSrc.Service
}

func NewFeedbackDescHandler(
	rg *gin.RouterGroup,
	feedbackDescSrc feedbackDescSrc.Service,
) {

	fh := &feedbackDescHandler{
		feedbackDescSrc: feedbackDescSrc,
	}

	apis.Handle(rg, http.MethodPost, "", fh.create)
	apis.Handle(rg, http.MethodGet, "", fh.getAll)
	apis.Handle(rg, http.MethodGet, "/:ID", fh.getByID)
}

func (h *feedbackDescHandler) create(ctx *gin.Context) {
	// var feedbackDescModel feedbackM.FeedbackDesc

	// ctx.BindJSON(&feedbackDescModel)

	// if err := h.feedbackDescSrc.Create(ctx, &feedbackDescModel); err != nil {
	// 	logrus.Error(err)
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *feedbackDescHandler) getAll(ctx *gin.Context) {

	// resp, err := h.feedbackDescSrc.GetAll(ctx)
	// if err != nil {
	// 	logrus.Error(err)
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// ctx.JSON(http.StatusOK, resp)
}

func (h *feedbackDescHandler) getByID(ctx *gin.Context) {
	// ID := ctx.Param("ID")

	// resp, err := h.feedbackDescSrc.GetByID(ctx, ID)
	// if err != nil {
	// 	logrus.Error(err)
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// ctx.JSON(http.StatusOK, resp)
}
