package constraint

import (
	"net/http"

	"example.com/creditcard/middlewares/apis"
	"example.com/creditcard/service/constraint"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	constraintM "example.com/creditcard/models/constraint"
)

type constraintHandler struct {
	dig.In

	constraintSrc constraint.Service
}

func NewConstraintHandler(
	rg *gin.RouterGroup,

	constraintSrc constraint.Service,
) {
	ch := &constraintHandler{
		constraintSrc: constraintSrc,
	}

	apis.Handle(rg, http.MethodGet, "", ch.getAll)
	apis.Handle(rg, http.MethodPost, "", ch.create)
	apis.Handle(rg, http.MethodGet, "/:ID", ch.get)
	apis.Handle(rg, http.MethodGet, "/privilageID/:privilageID", ch.getByPrivilageID)

}

func (h *constraintHandler) create(ctx *gin.Context) {
	var constraintModel constraintM.Constraint
	ctx.BindJSON(&constraintModel)

	err := h.constraintSrc.Create(ctx, &constraintModel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (h *constraintHandler) get(ctx *gin.Context) {

	ID := ctx.Param("ID")
	constraint, err := h.constraintSrc.GetByID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, constraint)
}

func (h *constraintHandler) getAll(ctx *gin.Context) {
	constraints, err := h.constraintSrc.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, constraints)
}

func (h *constraintHandler) getByPrivilageID(ctx *gin.Context) {
	privilageID := ctx.Param("privilageID")
	constraints, err := h.constraintSrc.GetByPrivilageID(ctx, privilageID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, constraints)
}
