package feedback

type CashCalculateType int32

const (
	FIXED_CASH_RETURN = iota + 1
	BONUS_MULTIPLY_CASH
)

type Cashback struct {
	CashCalculateType CashCalculateType `json:"cashCalculateType,omitempty"`

	Fixed float64 `json:"fixed,omitempty"`
	Bonus float64 `json:"bonus,omitempty"`

	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

type CashReturn struct {
	IsCashbackGet bool    `json:"isCashbackGet"`
	CashbackBonus float64 `json:"cashbackBonus"`

	ActualUseCash    int64   `json:"actualUseCash"`  // 實際用多少錢得到回饋
	ActualCashReturn float64 `json:"actualCashBack"` // 回饋多少cash

	CurrentCash int64   `json:"current"`
	TotalCash   float64 `json:"total"`
}
