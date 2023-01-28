package cashback

import (
	"context"
	"errors"

	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"

	feedbackM "example.com/creditcard/models/feedback"

	feedbackComp "example.com/creditcard/components/feedback"
)

type impl struct {
	Cashback *feedbackM.Cashback
}

func New(
	cashback *feedbackM.Cashback,
) feedbackComp.Component {
	return &impl{
		Cashback: cashback,
	}
}

func (im *impl) GetFeedback(ctx context.Context) *feedbackM.Feedback {
	return &feedbackM.Feedback{
		Cashback: im.Cashback,
	}
}

// 計算回饋
func (im *impl) Calculate(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.FeedReturn, error) {
	logrus.Info("cashback.Calculate pass :", pass)

	cash := int64(e.Cash)
	cashReturn := &feedbackM.CashReturn{
		// TotalBonus: im.Cashback.Bonus,
		Cash: cash,
	}

	// switch im.Cashback.CashCalculateType {
	// case feedbackM.FIXED_CASH_RETURN:
	// 	cashReturn.TotalBonus = im.Cashback.Fixed
	// 	break
	// case feedbackM.BONUS_MULTIPLY_CASH:
	// 	cashReturn.TotalBonus = im.Cashback.Bonus
	// 	break
	// }

	feedReturn := &feedbackM.FeedReturn{
		// FeedbackType:  ,
		CashReturn: cashReturn,
	}

	// 先定義一下
	var actualUseCash int64 = 0
	var actualCashReturn float64 = 0.0
	var cashReturnStatus feedbackM.CashReturnStatus = feedbackM.NONE_RETURN_CASH
	var cashReturnBonus float64 = 0.0

	if pass {
		// 取得可使用的回饋花費金額
		switch im.Cashback.CashCalculateType {
		case feedbackM.FIXED_CASH_RETURN:
			actualUseCash, actualCashReturn, cashReturnStatus = im.takeFixedCashReturn(ctx, cash)

			if cashReturnStatus != feedbackM.NONE_RETURN_CASH {
				cashReturnBonus = im.Cashback.Fixed
			}
			break
		case feedbackM.BONUS_MULTIPLY_CASH:
			actualUseCash, actualCashReturn, cashReturnStatus = im.multiplyCashReturn(ctx, cash)
			if cashReturnStatus != feedbackM.NONE_RETURN_CASH {
				cashReturnBonus = im.Cashback.Bonus
			}
			break
		default:
			return nil, errors.New("not found suitable im.Cashback.CashCalculateType")
		}

	}

	cashReturn.CashReturnBonus = cashReturnBonus
	cashReturn.CashReturnStatus = cashReturnStatus
	cashReturn.ActualUseCash = actualUseCash
	cashReturn.ActualCashReturn = actualCashReturn

	logrus.Info("cashback.Calculate cashReturn : ", cashReturn)
	logrus.Info("cashback.Calculate feedReturn : ", feedReturn)

	return feedReturn, nil
}

func (im *impl) takeFixedCashReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.CashReturnStatus) {

	if im.Cashback.Min == 0 && im.Cashback.Max == 0 {
		return cash, im.Cashback.Fixed, feedbackM.ALL_RETURN_CASH
	} else if im.Cashback.Min == 0 && im.Cashback.Max != 0 {
		if cash <= im.Cashback.Max {
			return cash, im.Cashback.Fixed, feedbackM.ALL_RETURN_CASH
		} else {
			return int64(im.Cashback.Fixed), im.Cashback.Fixed, feedbackM.SOME_RETURN_CASH
		}
	} else if im.Cashback.Min != 0 && im.Cashback.Max == 0 {
		if im.Cashback.Min <= cash {
			return cash, im.Cashback.Fixed, feedbackM.ALL_RETURN_CASH
		} else {
			return 0, 0, feedbackM.NONE_RETURN_CASH
		}
	} else {
		if im.Cashback.Min <= cash && cash <= im.Cashback.Max {
			return cash, im.Cashback.Fixed, feedbackM.ALL_RETURN_CASH
		} else if cash < im.Cashback.Min {
			return 0, 0, feedbackM.NONE_RETURN_CASH
		} else {
			return int64(im.Cashback.Fixed), im.Cashback.Fixed, feedbackM.SOME_RETURN_CASH
		}
	}
}

// 實際可以用多少錢拿回饋, 回饋多少, 回饋是否全拿
func (im *impl) multiplyCashReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.CashReturnStatus) {
	logrus.Info("multiplyCashReturn, ", im.Cashback.Min, im.Cashback.Max)

	if im.Cashback.Min == 0 && im.Cashback.Max == 0 {
		return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL_RETURN_CASH
	} else if im.Cashback.Min == 0 && im.Cashback.Max != 0 {
		if cash <= im.Cashback.Max {
			// logrus.Info("hello under max", im.Cashback.Max, im.Cashback.Bonus*float64(im.Cashback.Max), feedbackM.SOME_RETURN_CASH)
			return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL_RETURN_CASH
		} else {
			// logrus.Info("hello over max", im.Cashback.Max, im.Cashback.Bonus*float64(im.Cashback.Max), feedbackM.SOME_RETURN_CASH)
			return im.Cashback.Max, im.Cashback.Bonus * float64(im.Cashback.Max), feedbackM.SOME_RETURN_CASH
		}
	} else if im.Cashback.Min != 0 && im.Cashback.Max == 0 {
		if im.Cashback.Min <= cash {
			return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL_RETURN_CASH
		} else {
			return 0, 0, feedbackM.NONE_RETURN_CASH
		}
	} else {

		if im.Cashback.Min <= cash && cash <= im.Cashback.Max {
			return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL_RETURN_CASH
		} else if cash < im.Cashback.Min {
			return 0, 0, feedbackM.NONE_RETURN_CASH
		} else {
			return int64(im.Cashback.Bonus), im.Cashback.Bonus * float64(im.Cashback.Max), feedbackM.SOME_RETURN_CASH
		}

	}

}
