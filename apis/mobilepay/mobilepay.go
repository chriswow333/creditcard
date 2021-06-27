package mobilepay

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	mobilepayM "example.com/creditcard/models/mobilepay"
	"example.com/creditcard/service/mobilepay"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type mobilepayHandler struct {
	dig.In

	mobilepayService mobilepay.Service
}

func NewMobilepayHandler(
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
	var mobilepayModel mobilepayM.Mobilepay

	ctx.BindJSON(mobilepayModel)
	if err := h.mobilepayService.Create(ctx, &mobilepayModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *mobilepayHandler) updateByID(ctx *gin.Context) {

	var mobilepayModel mobilepayM.Mobilepay
	ctx.BindJSON(&mobilepayModel)
	ID := ctx.Param("ID")
	mobilepayModel.ID = ID
	if err := h.mobilepayService.UpdateByID(ctx, &mobilepayModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *mobilepayHandler) getAll(ctx *gin.Context) {
	mobilepays, err := h.mobilepayService.GetAll(ctx)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, mobilepays)
}
