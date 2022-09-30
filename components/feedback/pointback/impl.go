package cashback

import (
	"context"

	eventM "example.com/creditcard/models/event"

	feedbackM "example.com/creditcard/models/feedback"

	feedbackComp "example.com/creditcard/components/feedback"
)

type impl struct {
	Pointback *feedbackM.Pointback
}

func New(
	pointback *feedbackM.Pointback,
) feedbackComp.Component {
	return &impl{
		Pointback: pointback,
	}
}

func (im *impl) GetFeedback(ctx context.Context) *feedbackM.Feedback {
	return &feedbackM.Feedback{
		Pointback: im.Pointback,
	}
}

// 計算回饋
func (im *impl) Calculate(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.FeedReturn, error) {

	cash := int64(e.Cash)
	pointReturn := &feedbackM.PointReturn{}

	feedReturn := &feedbackM.FeedReturn{
		PointReturn: pointReturn,
	}

	// total := cost.Dollar.Total + cash // no increment for now
	total := cash

	pointReturn.TotalCash = float64(total)
	pointReturn.CurrentCash = cash

	// 先定義一下
	var actualUseCash int64 = 0
	var actualPointBack float64 = 0.0
	var feedReturnStatus feedbackM.FeedReturnStatus = feedbackM.NONE

	if pass {

		switch im.Pointback.PointCalculateType {
		case feedbackM.FIXED_POINT_RETURN:

			actualUseCash, actualPointBack, feedReturnStatus = im.takeFixedPointReturn(ctx, total)
			break

		case feedbackM.BONUS_MULTIPLY_POINT:

			actualUseCash, actualPointBack, feedReturnStatus = im.multiplyPointReturn(ctx, total)
			break
		}
		// 取得可使用的回饋花費金額

	}

	feedReturn.FeedReturnStatus = feedReturnStatus
	if feedReturnStatus == feedbackM.NONE {
		pointReturn.IsPointbackGet = false
	} else {
		pointReturn.IsPointbackGet = true
		pointReturn.PointbackBonus = (im.Pointback.Bonus) * 100
	}

	pointReturn.ActualUseCash = actualUseCash
	pointReturn.ActualPointBack = actualPointBack

	feedReturn.PointReturn = pointReturn
	// set cache
	// im.Feedback.Total = float64(total)
	// im.Feedback.Current = cash

	return feedReturn, nil
}

func (im *impl) takeFixedPointReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.FeedReturnStatus) {

	if im.Pointback.Min == 0 && im.Pointback.Max == 0 {

		return cash, im.Pointback.Fixed, feedbackM.ALL

	} else if im.Pointback.Min == 0 && im.Pointback.Max != 0 {

		if cash <= im.Pointback.Max {
			return cash, im.Pointback.Fixed, feedbackM.ALL
		} else {
			return 0, im.Pointback.Fixed, feedbackM.SOME
		}

	} else if im.Pointback.Min != 0 && im.Pointback.Max == 0 {

		if im.Pointback.Min <= cash {
			return cash, im.Pointback.Fixed, feedbackM.ALL
		} else {
			return 0, 0, feedbackM.NONE
		}

	} else {

		if im.Pointback.Min <= cash && cash <= im.Pointback.Max {
			return cash, im.Pointback.Fixed, feedbackM.ALL
		} else if cash < im.Pointback.Min {
			return 0, 0, feedbackM.NONE
		} else {
			return 0, im.Pointback.Fixed, feedbackM.SOME
		}

	}
}

// 實際可以用多少錢拿回饋, 回饋多少, 回饋是否全拿
func (im *impl) multiplyPointReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.FeedReturnStatus) {

	if im.Pointback.Min == 0 && im.Pointback.Max == 0 {

		return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL

	} else if im.Pointback.Min == 0 && im.Pointback.Max != 0 {

		if cash <= im.Pointback.Max {
			return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL
		} else {
			return 0, im.Pointback.Bonus * float64(im.Pointback.Max), feedbackM.NONE
		}

	} else if im.Pointback.Min != 0 && im.Pointback.Max == 0 {

		if im.Pointback.Min <= cash {
			return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL
		} else {
			return 0, 0, feedbackM.NONE
		}

	} else {

		if im.Pointback.Min <= cash && cash <= im.Pointback.Max {
			return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL
		} else if cash < im.Pointback.Min {
			return 0, 0, feedbackM.NONE
		} else {
			return im.Pointback.Max, im.Pointback.Bonus * float64(im.Pointback.Max), feedbackM.SOME
		}

	}

}
