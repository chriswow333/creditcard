package cashback

import (
	"context"

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

	if pass {
		// 取得可使用的回饋花費金額
		actualUseCash, actualCashReturn, feedReturnStatus = im.takeCashReturn(ctx, total)

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
	// set cache
	// im.Feedback.Total = float64(total)
	// im.Feedback.Current = cash

	return feedReturn, nil
}

// 實際可以用多少錢拿回饋, 回饋多少, 回饋是否全拿
func (im *impl) takeCashReturn(ctx context.Context, cash int64) (int64, float64, feedbackM.FeedReturnStatus) {

	if im.Cashback.Min <= cash && cash <= im.Cashback.Max {
		return cash, im.Cashback.Bonus * float64(cash), feedbackM.ALL
	} else if cash < im.Cashback.Min {
		return 0, 0, feedbackM.NONE
	} else {
		return im.Cashback.Max, im.Cashback.Bonus * float64(im.Cashback.Max), feedbackM.SOME
	}
}
