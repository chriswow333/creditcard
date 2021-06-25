package streaming

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/streaming"
	"github.com/gin-gonic/gin"
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

}

func (h *streamingHandler) updateByID(ctx *gin.Context) {

}

func (h *streamingHandler) getAll(ctx *gin.Context) {

}
