package feedback

type FeedReturnStatus int32

const (
	Full FeedReturnStatus = iota
	PartOf
	None
)

type Feedback struct {
	Cashback *Cashback `json:"cashback"`
}

type FeedbackResp struct {
	CashbackResp *CashbackResp `json:"cashbackResp"`
}

type FeedReturn struct {
	FeedReturnStatus FeedReturnStatus `json:"feedReturnStatus"`

	CashReturn *CashReturn `json:"cashReturn"`
}

func TransferFeedbackResp(feedback *Feedback) *FeedbackResp {
	feedbackResp := &FeedbackResp{
		CashbackResp: TransferCashbackResp(feedback.Cashback),
	}

	return feedbackResp
}
