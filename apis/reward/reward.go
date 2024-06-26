package reward

import (
	"net/http"
	"runtime/debug"

	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"example.com/creditcard/middlewares/apis"
	rewardM "example.com/creditcard/models/reward"
	"example.com/creditcard/service/reward"
)

type rewardHandler struct {
	dig.In

	rewardSrc reward.Service
}

func NewrewardHandler(
	rg *gin.RouterGroup,

	rewardSrc reward.Service,
) {

	ph := &rewardHandler{
		rewardSrc: rewardSrc,
	}

	apis.Handle(rg, http.MethodPost, "", ph.create)
	apis.Handle(rg, http.MethodPost, "/:ID", ph.updateByID)
}

func (h *rewardHandler) create(ctx *gin.Context) {
	var rewardModel rewardM.Reward

	ctx.BindJSON(&rewardModel)

	if err := h.rewardSrc.Create(ctx, &rewardModel); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *rewardHandler) updateByID(ctx *gin.Context) {
	var rewardModel rewardM.Reward

	ctx.BindJSON(&rewardModel)
	if err := h.rewardSrc.UpdateByID(ctx, &rewardModel); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}
