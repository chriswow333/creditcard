package feedback

type CashBackType int32

const (
	NTD CashBackType = iota
)

type CashBack struct {
	ActualUseCash  int64   `json:"actualUseCash"`  // 實際用多少錢得到回饋
	ActualCashBack float64 `json:"actualCashBack"` // 回饋多少

	CashBackType CashBackType `json:"cashBackType"`

	CashBackLimit CashBackLimit `json:"cashBackLimit"`
}

type CashBackLimit struct {
	Bonus float64 `json:"bonus"`
	Min   int64   `json:"min"`
	Max   int64   `json:"max"`
}
