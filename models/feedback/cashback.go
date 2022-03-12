package feedback

type CashbackType int32

const (
	NTD CashbackType = iota
)

type Cashback struct {
	CashbackType CashbackType `json:"cashbackType"`
	Bonus        float64      `json:"bonus"`
	Min          int64        `json:"min"`
	Max          int64        `json:"max"`
}

type CashbackResp struct {
	CashbackType CashbackType `json:"cashbackType"`
	Bonus        float64      `json:"bonus"`
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

func TransferCashbackResp(cashback *Cashback) *CashbackResp {

	cashbackResp := &CashbackResp{
		CashbackType: cashback.CashbackType,
		Bonus:        cashback.Bonus * 100,
		Max:          cashback.Max,
		Min:          cashback.Min,
	}

	return cashbackResp
}
