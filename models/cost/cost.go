package cost

type CurrencyType int32

const (
	NTD CurrencyType = iota
)

type Cost struct {
	Current  int64        `json:"current"`
	Total    int64        `json:"total"`
	Currency CurrencyType `json:"currency"`
}

type CostLimit struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Desc         string       `json:"desc"`
	CurrencyType CurrencyType `json:"currencyType"`
	AtLeast      int64        `json:"atLeast"`
	AtMost       int64        `json:"atMost"`
}
