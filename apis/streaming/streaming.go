package streaming

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	streamingM "example.com/creditcard/models/streaming"
	"example.com/creditcard/service/streaming"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type streamingHandler struct {
	dig.In

	streamingService streaming.Service
}

func NewEcommerceHandler(
	rg *gin.RouterGroup,

	streamingService streaming.Service,

) {
	eh := &streamingHandler{
		streamingService: streamingService,
	}

	apis.Handle(rg, http.MethodPost, "", eh.create)
	apis.Handle(rg, http.MethodPost, "/:ID", eh.updateByID)
	apis.Handle(rg, http.MethodGet, "", eh.getAll)
}

func (h *streamingHandler) create(ctx *gin.Context) {
	var streamingModel streamingM.Streaming

	ctx.BindJSON(&streamingModel)

	if err := h.streamingService.Create(ctx, &streamingModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *streamingHandler) updateByID(ctx *gin.Context) {
	var streamingModel streamingM.Streaming

	ctx.BindJSON(&streamingModel)
	ID := ctx.Param("ID")
	streamingModel.ID = ID
	if err := h.streamingService.UpdateByID(ctx, &streamingModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *streamingHandler) getAll(ctx *gin.Context) {
	streamings, err := h.streamingService.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, streamings)
}
