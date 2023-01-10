package redback

import (
	"context"

	feedbackM "example.com/creditcard/models/feedback"

	feedbackComp "example.com/creditcard/components/feedback"

	eventM "example.com/creditcard/models/event"
)

type impl struct {
	Redback *feedbackM.Redback
}

func New(
	redback *feedbackM.Redback,
) feedbackComp.Component {

	return &impl{
		Redback: redback,
	}
}

func (im *impl) GetFeedback(ctx context.Context) *feedbackM.Feedback {

	return &feedbackM.Feedback{
		Redback: im.Redback,
	}
}

func (im *impl) Calculate(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.FeedReturn, error) {

	return nil, nil
	// 	cash := int64(e.Cash)

	// 	redReturn := &feedbackM.RedReturn{
	// 		RedbackTimes: im.Redback.Times,
	// 	}

	// 	feedReturn := &feedbackM.FeedReturn{
	// 		RedReturn: redReturn,
	// 	}

	// 	total := cash

	// 	// define first
	// 	redReturn.TotalCash = float64(total)
	// 	redReturn.CurrentCash = cash

	// 	var actualUseCash int64 = 0
	// 	var actualRedBack int64 = 0
	// 	var feedReturnStatus feedbackM.FeedReturnStatus = feedbackM.NONE

	// 	if pass {

	// 		switch im.Redback.RedCalculateType {
	// 		case feedbackM.RED_TIMES:

	// 			if im.Redback.Min == 0 && im.Redback.Max == 0 {

	// 				actualUseCash = cash
	// 				actualRedBack = cash * im.Redback.Times
	// 				feedReturnStatus = feedbackM.ALL

	// 			} else if im.Redback.Min == 0 && im.Redback.Max != 0 {

	// 				if cash <= im.Redback.Max {

	// 					actualUseCash = cash
	// 					actualRedBack = cash * im.Redback.Times
	// 					feedReturnStatus = feedbackM.ALL

	// 				} else {

	// 					actualUseCash = im.Redback.Max
	// 					actualRedBack = im.Redback.Max * im.Redback.Times
	// 					feedReturnStatus = feedbackM.SOME

	// 				}

	// 			} else if im.Redback.Min != 0 && im.Redback.Max == 0 {
	// 				if cash >= im.Redback.Min {

	// 					actualUseCash = cash
	// 					actualRedBack = cash * im.Redback.Times
	// 					feedReturnStatus = feedbackM.ALL

	// 				}
	// 			} else {
	// 				if im.Redback.Min <= cash && cash <= im.Redback.Max {
	// 					actualUseCash = cash
	// 					actualRedBack = cash * im.Redback.Times
	// 					feedReturnStatus = feedbackM.ALL
	// 				} else if im.Redback.Max < cash {
	// 					actualUseCash = im.Redback.Max
	// 					actualRedBack = im.Redback.Max * im.Redback.Times
	// 					feedReturnStatus = feedbackM.SOME

	// 				}
	// 			}

	// 			break

	// 		default:
	// 			logrus.WithFields(logrus.Fields{
	// 				"": "",
	// 			}).Error("not found red calculate type")
	// 			return nil, errors.New("not found red calculate typ")

	// 		}

	// 	}

	// 	feedReturn.FeedReturnStatus = feedReturnStatus
	// 	if feedReturn.FeedReturnStatus == feedbackM.NONE {
	// 		redReturn.IsRedGet = false
	// 	} else {
	// 		redReturn.IsRedGet = true
	// 		redReturn.RedbackTimes = im.Redback.Times
	// 	}

	// 	feedReturn.RedReturn.ActualUseCash = actualUseCash
	// 	feedReturn.RedReturn.ActualRedback = float64(actualRedBack)

	// 	return feedReturn, nil
}
