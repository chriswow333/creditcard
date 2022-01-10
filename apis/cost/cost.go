package cost

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/cost"

	costM "example.com/creditcard/models/cost"
)

type costHandler struct {
	dig.In

	costService cost.Service
}

func NewCostHandler(
	rg *gin.RouterGroup,

	costService cost.Service,
) {

	ch := &costHandler{
		costService: costService,
	}

	apis.Handle(rg, http.MethodPost, "/rewardID/:ID", ch.update)
	apis.Handle(rg, http.MethodGet, "/rewardID/ID", ch.get)
}

func (h *costHandler) update(ctx *gin.Context) {

	rewardID := ctx.Param("ID")

	var costModel *costM.Cost

	ctx.BindJSON(&costModel)

	if err := h.costService.UpdateByRewardID(ctx, rewardID, costModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *costHandler) get(ctx *gin.Context) {

	rewardID := ctx.Param("ID")

	var costModel *costM.Cost

	ctx.BindJSON(&costModel)

	if err := h.costService.UpdateByRewardID(ctx, rewardID, costModel); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}
