package constraint

import (
	"fmt"
	"net/http"

	"go.uber.org/dig"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/constraint"

	constraintM "example.com/creditcard/models/constraint"
)

type constraintHandler struct {
	dig.In

	constraintService constraint.Service
}

func NewConstraintHandler(
	rg *gin.RouterGroup,

	constraintService constraint.Service,

) {
	ch := &constraintHandler{
		constraintService: constraintService,
	}

	apis.Handle(rg, http.MethodPost, "/rewardID/:ID", ch.update)
	apis.Handle(rg, http.MethodGet, "/rewardID/:ID", ch.get)
}

func (h *constraintHandler) update(ctx *gin.Context) {

	rewardID := ctx.Param("ID")

	var constraintModels []*constraintM.Constraint
	ctx.BindJSON(&constraintModels)

	fmt.Println(constraintModels)

	if err := h.constraintService.UpdateByRewardID(ctx, rewardID, constraintModels); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (h *constraintHandler) get(ctx *gin.Context) {

	rewardID := ctx.Param("ID")

	constraints, err := h.constraintService.GetByRewardID(ctx, rewardID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, constraints)
}
