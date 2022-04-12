package feedback

type FeedReturnStatus int32

const (
	ALL FeedReturnStatus = iota + 1
	SOME
	NONE
)

type Feedback struct {
	Cashback *Cashback `json:"cashback,omitempty"`
}

type FeedReturn struct {
	FeedReturnStatus FeedReturnStatus `json:"feedReturnStatus,omitempty"`

	CashReturn *CashReturn `json:"cashReturn,omitempty"`
}
