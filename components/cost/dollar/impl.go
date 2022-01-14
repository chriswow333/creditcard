package dollar

import (
	"context"

	costM "example.com/creditcard/models/cost"
	dollarM "example.com/creditcard/models/dollar"
	eventM "example.com/creditcard/models/event"

	costComp "example.com/creditcard/components/cost"
)

type impl struct {
	Dollar *dollarM.Dollar
}

func New(
	dollar *dollarM.Dollar,
) costComp.Component {
	return &impl{
		Dollar: dollar,
	}
}

func (im *impl) GetCost(ctx context.Context, e *eventM.Event) *costM.Cost {

	cash := int64(e.Cash)

	return &costM.Cost{
		CostType: costM.Dollar,
		Dollar: &dollarM.Dollar{
			Current:     cash,
			Total:       im.Dollar.Total,
			DollarLimit: im.Dollar.DollarLimit,
			DollarType:  im.Dollar.DollarType,
		},
	}
}

// 計算回饋
func (im *impl) Calculate(ctx context.Context, e *eventM.Event, pass bool) (*costM.Cost, error) {

	cash := int64(e.Cash)

	cost := &costM.Cost{
		CostType: costM.Dollar,
		Dollar: &dollarM.Dollar{
			Current:     cash,
			Total:       im.Dollar.Total,
			DollarLimit: im.Dollar.DollarLimit,
			DollarType:  im.Dollar.DollarType,
		},
	}

	// total := cost.Dollar.Total + cash // no increment for now
	total := cash

	// 先定義一下
	var dollarBonusBack int64 = 0
	var dollarBack float64 = 0.0
	var pointBack dollarM.PointBack = dollarM.None

	if pass {
		// 取得可使用的回饋花費金額
		dollarBonusBack, dollarBack, pointBack = im.takeDollarBonusBack(ctx, total)
	}

	cost.Dollar.Total = total

	cost.Dollar.DollarBonusBack = dollarBonusBack
	cost.Dollar.DollarBack = dollarBack
	cost.Dollar.PointBack = pointBack

	// set cache
	im.Dollar.Total = total
	im.Dollar.Current = cash

	return cost, nil
}

// 實際可以用多少錢拿回饋, 回饋多少, 回饋是否全拿
func (im *impl) takeDollarBonusBack(ctx context.Context, cash int64) (int64, float64, dollarM.PointBack) {

	if im.Dollar.DollarLimit.Min <= cash && cash <= im.Dollar.DollarLimit.Max {
		return cash, im.Dollar.DollarLimit.Point * float64(cash), dollarM.Full
	} else if cash < im.Dollar.DollarLimit.Min {
		return 0, 0, dollarM.None
	} else {
		return im.Dollar.DollarLimit.Max, im.Dollar.DollarLimit.Point * float64(im.Dollar.DollarLimit.Max), dollarM.PartOf
	}
}
