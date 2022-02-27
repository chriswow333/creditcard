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

type FeedReturn struct {
	FeedReturnStatus FeedReturnStatus `json:"feedReturnStatus"`

	CashReturn *CashReturn `json:"cashReturn"`
}
