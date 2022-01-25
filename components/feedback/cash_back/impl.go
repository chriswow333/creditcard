package cash_back

import (
	"context"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"

	feedbackComp "example.com/creditcard/components/feedback"
)

type impl struct {
	Feedback *feedbackM.Feedback
	CashBack *feedbackM.CashBack
}

func New(
	feedback *feedbackM.Feedback,
	cashBack *feedbackM.CashBack,
) feedbackComp.Component {
	return &impl{
		Feedback: feedback,
		CashBack: cashBack,
	}
}

// func (im *impl) GetFeedback(ctx context.Context, e *eventM.Event) *feedbackM.Feedback {

// 	return &feedbackM.Feedback{
// 		FeedbackType: im.Feedback.FeedbackType,
// 		Total:        im.Feedback.Total,
// 		DollarType:   im.Feedback.DollarType,
// 		IsRewardGet:  im.Feedback.IsRewardGet,
// 		CashBack: &feedbackM.CashBack{
// 			ActualUseCash:  0,
// 			ActualCashBack: 0.0,
// 			CashBackLimit:  im.CashBack.CashBackLimit,
// 		},
// 	}
// }

// 計算回饋
func (im *impl) Calculate(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.Feedback, error) {

	cash := int64(e.Cash)

	feedback := &feedbackM.Feedback{
		FeedbackType:   im.Feedback.FeedbackType,
		FeedbackStatus: feedbackM.None,
		CashBack: &feedbackM.CashBack{
			ActualUseCash:  0,
			ActualCashBack: 0.0,
			CashBackLimit:  im.CashBack.CashBackLimit,
		},
	}

	// total := cost.Dollar.Total + cash // no increment for now
	total := cash

	// 先定義一下
	var actualUseCash int64 = 0
	var actualCashBack float64 = 0.0
	var feedbackStatus feedbackM.FeedbackStatus = feedbackM.None

	if pass {
		// 取得可使用的回饋花費金額
		actualUseCash, actualCashBack, feedbackStatus = im.takeDollarBonusBack(ctx, total)
	}

	feedback.Current = cash
	feedback.Total = float64(total)

	feedback.CashBack.ActualUseCash = actualUseCash
	feedback.CashBack.ActualCashBack = actualCashBack
	feedback.FeedbackStatus = feedbackStatus

	if feedbackStatus == feedbackM.None {
		feedback.IsRewardGet = false
	} else {
		feedback.IsRewardGet = true
	}

	// set cache
	im.Feedback.Total = float64(total)
	im.Feedback.Current = cash

	return feedback, nil
}

// 實際可以用多少錢拿回饋, 回饋多少, 回饋是否全拿
func (im *impl) takeDollarBonusBack(ctx context.Context, cash int64) (int64, float64, feedbackM.FeedbackStatus) {

	if im.CashBack.CashBackLimit.Min <= cash && cash <= im.CashBack.CashBackLimit.Max {
		return cash, im.CashBack.CashBackLimit.Bonus * float64(cash), feedbackM.Full
	} else if cash < im.CashBack.CashBackLimit.Min {
		return 0, 0, feedbackM.None
	} else {
		return im.CashBack.CashBackLimit.Max, im.CashBack.CashBackLimit.Bonus * float64(im.CashBack.CashBackLimit.Max), feedbackM.PartOf
	}
}
