package feedback

type CashbackType int32

const (
	NTD CashbackType = iota + 1
)

type Cashback struct {
	CashbackType CashbackType `json:"cashbackType,omitempty"`
	Bonus        float64      `json:"bonus,omitempty"`
	Min          int64        `json:"min"`
	Max          int64        `json:"max"`
}

type CashReturn struct {
	IsCashbackGet bool    `json:"isFeedbackGet"`
	CashbackBonus float64 `json:"cashReturnBonus"`

	ActualUseCash    int64   `json:"actualUseCash"`  // 實際用多少錢得到回饋
	ActualCashReturn float64 `json:"actualCashBack"` // 回饋多少cash

	CurrentCash int64   `json:"current"`
	TotalCash   float64 `json:"total"`
}
