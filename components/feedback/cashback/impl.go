package cashback

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	eventM "example.com/creditcard/models/event"

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

	cash := int64(e.Cash)
	cashReturn := &feedbackM.CashReturn{}

	feedReturn := &feedbackM.FeedReturn{
		CashReturn: cashReturn,
	}

	// total := cost.Dollar.Total + cash // no increment for now
	total := cash

	cashReturn.TotalCash = float64(total)
	cashReturn.CurrentCash = cash

	// 先定義一下
	var actualUseCash int64 = 0
	var actualCashReturn float64 = 0.0
	var feedReturnStatus feedbackM.FeedReturnStatus = feedbackM.NONE
	fmt.Println("pass " + strconv.FormatBool(pass))

	if pass {
		// 取得可使用的回饋花費金額
		switch im.Cashback.CashCalculateType {
		case feedbackM.FIXED_CASH_RETURN:
			actualUseCash, actualCashReturn, feedReturnStatus = im.takeFixedCashReturn(ctx, total)
			break
		case feedbackM.BONUS_MULTIPLY_CASH:
			actualUseCash, actualCashReturn, feedReturnStatus = im.multiplyCashReturn(ctx, total)
			break
		default:
			return nil, errors.New("not found suitable im.Cashback.CashCalculateType")

		}

	}

	feedReturn.FeedReturnStatus = feedReturnStatus
	if feedReturnStatus == feedbackM.NONE {
		cashReturn.IsCashbackGet = false
	} else {
		cashReturn.IsCashbackGet = true
		cashReturn.CashbackBonus = (im.Cashback.Bonus) * 100
	}

	cashReturn.ActualUseCash = actualUseCash
	cashReturn.ActualCashReturn = actualCashReturn

	return feedReturn, nil
}

func (im *impl) takeFixedCashReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.FeedReturnStatus) {
	if im.Cashback.Min == 0 && im.Cashback.Max == 0 {
		return cash, im.Cashback.Fixed, feedbackM.ALL
	} else if im.Cashback.Min == 0 && im.Cashback.Max != 0 {
		if cash <= im.Cashback.Max {
			return cash, im.Cashback.Fixed, feedbackM.ALL
		} else {
			return int64(im.Cashback.Fixed), im.Cashback.Fixed, feedbackM.SOME
		}
	} else if im.Cashback.Min != 0 && im.Cashback.Max == 0 {
		if im.Cashback.Min <= cash {
			return cash, im.Cashback.Fixed, feedbackM.ALL
		} else {
			return 0, 0, feedbackM.NONE
		}
	} else {
		if im.Cashback.Min <= cash && cash <= im.Cashback.Max {
			return cash, im.Cashback.Fixed, feedbackM.ALL
		} else if cash < im.Cashback.Min {
			return 0, 0, feedbackM.NONE
		} else {
			return int64(im.Cashback.Fixed), im.Cashback.Fixed, feedbackM.SOME
		}

	}
}

// 實際可以用多少錢拿回饋, 回饋多少, 回饋是否全拿
func (im *impl) multiplyCashReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.FeedReturnStatus) {

	if im.Cashback.Min == 0 && im.Cashback.Max == 0 {
		return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL
	} else if im.Cashback.Min == 0 && im.Cashback.Max != 0 {
		if cash <= im.Cashback.Max {
			return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL
		} else {
			return cash, im.Cashback.Bonus * float64(im.Cashback.Max), feedbackM.SOME
		}
	} else if im.Cashback.Min != 0 && im.Cashback.Max == 0 {
		if im.Cashback.Min <= cash {
			return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL
		} else {
			return 0, 0, feedbackM.NONE
		}
	} else {
		if im.Cashback.Min <= cash && cash <= im.Cashback.Max {
			return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL
		} else if cash < im.Cashback.Min {
			return 0, 0, feedbackM.NONE
		} else {
			return int64(im.Cashback.Bonus), im.Cashback.Bonus * float64(cash), feedbackM.SOME
		}
	}

}
