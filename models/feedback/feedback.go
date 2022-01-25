package feedback

type FeedbackType int32

const (
	Cash FeedbackType = iota
	Point
)

type Feedback struct {
	Current int64   `json:"current"`
	Total   float64 `json:"total"`

	IsRewardGet    bool           `json:"isRewardGet"`
	FeedbackType   FeedbackType   `json:"feedbackType"`
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
