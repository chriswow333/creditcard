package onlinegame

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/onlinegame"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type onlinegameHandler struct {
	dig.In

	onlinegameService onlinegame.Service
}

func NewEcommerceHandler(
	rg *gin.RouterGroup,

	onlinegameService onlinegame.Service,

) {
	eh := &onlinegameHandler{
		onlinegameService: onlinegameService,
	}

	apis.Handle(rg, http.MethodPost, "", eh.create)
	apis.Handle(rg, http.MethodPost, "/:ID", eh.updateByID)
	apis.Handle(rg, http.MethodGet, "", eh.getAll)
}

func (h *onlinegameHandler) create(ctx *gin.Context) {

}

func (h *onlinegameHandler) updateByID(ctx *gin.Context) {

}

func (h *onlinegameHandler) getAll(ctx *gin.Context) {

}
