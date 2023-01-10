package cashback

import (
	"context"

	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"

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
	pointReturn := &feedbackM.PointReturn{
		Cash: cash,
		// TotalBonus: im.Pointback.Bonus,
	}

	// switch im.Pointback.PointCalculateType {
	// case feedbackM.FIXED_POINT_RETURN:
	// 	pointReturn.TotalBonus = im.Pointback.Fixed
	// 	break
	// case feedbackM.BONUS_MULTIPLY_POINT:
	// 	pointReturn.TotalBonus = im.Pointback.Bonus
	// 	break
	// }

	feedReturn := &feedbackM.FeedReturn{
		PointReturn: pointReturn,
	}

	// 先定義一下
	var actualUseCash int64 = 0
	var actualPointReturn float64 = 0.0
	var pointReturnStatus feedbackM.PointReturnStatus = feedbackM.NONE_RETURN_POINT
	var pointReturnBonus float64 = 0.0

	if pass {

		switch im.Pointback.PointCalculateType {
		case feedbackM.FIXED_POINT_RETURN:
			actualUseCash, actualPointReturn, pointReturnStatus = im.takeFixedPointReturn(ctx, cash)
			if pointReturnStatus != feedbackM.NONE_RETURN_POINT {
				pointReturnBonus = im.Pointback.Fixed
			}
			break

		case feedbackM.BONUS_MULTIPLY_POINT:
			actualUseCash, actualPointReturn, pointReturnStatus = im.multiplyPointReturn(ctx, cash)
			if pointReturnStatus != feedbackM.NONE_RETURN_POINT {
				pointReturnBonus = im.Pointback.Bonus
			}
			break
		}
		// 取得可使用的回饋花費金額

	}

	pointReturn.PointReturnBonus = pointReturnBonus
	pointReturn.PointReturnStatus = pointReturnStatus
	pointReturn.ActualUseCash = actualUseCash
	pointReturn.ActualPointReturn = actualPointReturn

	feedReturn.PointReturn = pointReturn

	logrus.Info("pointback.Calculate ", pointReturn)
	return feedReturn, nil
}

func (im *impl) takeFixedPointReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.PointReturnStatus) {

	if im.Pointback.Min == 0 && im.Pointback.Max == 0 {

		return cash, im.Pointback.Fixed, feedbackM.ALL_RETURN_POINT

	} else if im.Pointback.Min == 0 && im.Pointback.Max != 0 {

		if cash <= im.Pointback.Max {

			return cash, im.Pointback.Fixed, feedbackM.ALL_RETURN_POINT
		} else {
			return 0, im.Pointback.Fixed, feedbackM.SOME_RETURN_POINT
		}

	} else if im.Pointback.Min != 0 && im.Pointback.Max == 0 {
		if im.Pointback.Min <= cash {
			return cash, im.Pointback.Fixed, feedbackM.ALL_RETURN_POINT
		} else {
			return 0, 0, feedbackM.NONE_RETURN_POINT
		}

	} else {

		if im.Pointback.Min <= cash && cash <= im.Pointback.Max {
			return cash, im.Pointback.Fixed, feedbackM.ALL_RETURN_POINT
		} else if cash < im.Pointback.Min {
			return 0, 0, feedbackM.NONE_RETURN_POINT
		} else {
			return 0, im.Pointback.Fixed, feedbackM.SOME_RETURN_POINT
		}

	}
}

// 實際可以用多少錢拿回饋, 回饋多少, 回饋是否全拿
func (im *impl) multiplyPointReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.PointReturnStatus) {

	if im.Pointback.Min == 0 && im.Pointback.Max == 0 {

		return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL_RETURN_POINT

	} else if im.Pointback.Min == 0 && im.Pointback.Max != 0 {

		if cash <= im.Pointback.Max {
			return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL_RETURN_POINT
		} else {
			return 0, im.Pointback.Bonus * float64(im.Pointback.Max), feedbackM.NONE_RETURN_POINT
		}

	} else if im.Pointback.Min != 0 && im.Pointback.Max == 0 {
		if im.Pointback.Min <= cash {
			return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL_RETURN_POINT
		} else {
			return 0, 0, feedbackM.NONE_RETURN_POINT
		}

	} else {

		if im.Pointback.Min <= cash && cash <= im.Pointback.Max {
			return cash, im.Pointback.Bonus * float64(cash), feedbackM.ALL_RETURN_POINT
		} else if cash < im.Pointback.Min {
			return 0, 0, feedbackM.NONE_RETURN_POINT
		} else {
			return im.Pointback.Max, im.Pointback.Bonus * float64(im.Pointback.Max), feedbackM.SOME_RETURN_POINT
		}

	}

}
