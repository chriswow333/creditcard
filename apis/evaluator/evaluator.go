package evaluator

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"example.com/creditcard/middlewares/apis"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/modules/evaluator"
)

type evaluatorHandler struct {
	dig.In

	evaluatorModule evaluator.Module
}

func NewEvaluatorHandler(
	rg *gin.RouterGroup,

	evaluatorModule evaluator.Module,
) {

	eh := &evaluatorHandler{
		evaluatorModule: evaluatorModule,
	}

	apis.Handle(rg, http.MethodPost, "/evaluate", eh.evaluate)
}

func (h *evaluatorHandler) evaluate(ctx *gin.Context) {

	var eventModel eventM.Event

	ctx.Bind(&eventModel)

	logrus.Info("/evaluate", eventModel)

	resp, err := h.evaluatorModule.Evaluate(ctx, &eventModel)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
