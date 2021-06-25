package mobilepay

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/mobilepay"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type mobilepayHandler struct {
	dig.In

	mobilepayService mobilepay.Service
}

func NewEcommerceHandler(
	rg *gin.RouterGroup,

	mobilepayService mobilepay.Service,

) {
	eh := &mobilepayHandler{
		mobilepayService: mobilepayService,
	}

	apis.Handle(rg, http.MethodPost, "", eh.create)
	apis.Handle(rg, http.MethodPost, "/:ID", eh.updateByID)
	apis.Handle(rg, http.MethodGet, "", eh.getAll)
}

func (h *mobilepayHandler) create(ctx *gin.Context) {

}

func (h *mobilepayHandler) updateByID(ctx *gin.Context) {

}

func (h *mobilepayHandler) getAll(ctx *gin.Context) {

}
