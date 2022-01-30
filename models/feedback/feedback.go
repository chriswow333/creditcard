package feedback

type Feedback struct {
	Current int64   `json:"current"`
	Total   float64 `json:"total"`

	IsFeedbackGet  bool           `json:"isFeedbackGet"`
	FeedbackStatus FeedbackStatus `json:"feedbackStatus"` // 是否回饋全拿

	CashBack  *CashBack  `json:"cashBack,omitempty"`
	BonusBack *BonusBack `json:"bonusBack,omitempty"`
}

type FeedbackStatus int32

const (
	Full FeedbackStatus = iota
	PartOf
	None
)
