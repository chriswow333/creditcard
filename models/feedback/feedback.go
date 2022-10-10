package feedback

type FeedReturnStatus int32

const (
	ALL FeedReturnStatus = iota + 1
	SOME
	NONE
)

type FeedbackDesc struct {
	ID                    string   `json:"id"`
	FeedbackType          int32    `json:"feedbackType"`
	FeedbackCalculateType int32    `json:"feedbackCalculateType"`
	Name                  string   `json:"name"`
	Descs                 []string `json:"descs"`
}

type Feedback struct {
	Cashback  *Cashback  `json:"cashback,omitempty"`
	Pointback *Pointback `json:"pointback,omitempty"`
	Redback   *Redback   `json:"redback,omitempty"`
}

type FeedReturn struct {
	FeedReturnStatus FeedReturnStatus `json:"feedReturnStatus,omitempty"`

	CashReturn  *CashReturn  `json:"cashReturn,omitempty"`
	PointReturn *PointReturn `json:"pointReturn,omitempty"`
	RedReturn   *RedReturn   `json:"redReturn,omitempty"`
}
