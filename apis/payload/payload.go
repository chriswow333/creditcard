package payload

import (
	"net/http"

	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/payload"

	payloadM "example.com/creditcard/models/payload"
)

type rewardHandler struct {
	dig.In

	payloadSrc payload.Service
}

func NewrewardHandler(
	rg *gin.RouterGroup,

	payloadSrc payload.Service,
) {

	ph := &rewardHandler{
		payloadSrc: payloadSrc,
	}

	apis.Handle(rg, http.MethodPost, "/rewardID/:ID", ph.updateByID)
}

func (h *rewardHandler) updateByID(ctx *gin.Context) {
	var payloadsModel []*payloadM.Payload

	ctx.BindJSON(&payloadsModel)
	ID := ctx.Param("ID")

	if err := h.payloadSrc.UpdateByID(ctx, ID, payloadsModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
