package feedback

// type FeedReturnStatus int32

// const (
// 	ALL FeedReturnStatus = iota + 1
// 	SOME
// 	NONE
// )

// type FeedbackType int32

// const (
// 	CASH = iota + 1
// 	POINT
// 	RED
// )

// type FeedbackDesc struct {
// 	ID                    string   `json:"id"`
// 	FeedbackType          int32    `json:"feedbackType"`
// 	FeedbackCalculateType int32    `json:"feedbackCalculateType"`
// 	Name                  string   `json:"name"`
// 	Descs                 []string `json:"descs"`
// }

type Feedback struct {
	// FeedbackType FeedbackType `json:"feedbackType"`

	Cashback  *Cashback  `json:"cashback,omitempty"`
	Pointback *Pointback `json:"pointback,omitempty"`
	Redback   *Redback   `json:"redback,omitempty"`
}

type FeedReturn struct {
	// FeedbackType FeedbackType `json:"feedbackType"`

	CashReturn  *CashReturn  `json:"cashReturn,omitempty"`
	PointReturn *PointReturn `json:"pointReturn,omitempty"`
	RedReturn   *RedReturn   `json:"redReturn,omitempty"`
}

type FeedbackBonus struct {
	CashFeedbackBonus  *CashFeedbackBonus  `json:"cashFeedbackBonus,omitempty"`
	PointFeedbackBonus *PointFeedbackBonus `json:"pointFeedbackBonus,omitempty"`
}
