package bonus

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/cost"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	costM "example.com/creditcard/models/cost"
)

type bonusHandler struct {
	dig.In

	costSrc cost.Service
}

func NewBonusHandler(
	rg *gin.RouterGroup,

	costSrc cost.Service,

) {

	bh := &bonusHandler{
		costSrc: costSrc,
	}

	apis.Handle(rg, http.MethodGet, "/rewardID/:ID", bh.get)
	apis.Handle(rg, http.MethodPost, "/rewardID/:ID", bh.update)

}

func (h *bonusHandler) update(ctx *gin.Context) {
	rewardID := ctx.Param("ID")
	var costModel *costM.Cost
	ctx.BindJSON(&costModel)

	if err := h.costSrc.UpdateByRewardID(ctx, rewardID, costModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *bonusHandler) get(ctx *gin.Context) {

	rewardID := ctx.Param("ID")
	bonus, err := h.costSrc.GetByRewardID(ctx, rewardID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, bonus)
}
