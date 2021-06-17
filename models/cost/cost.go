package cost

type CurrencyType int32

const (
	NTD CurrencyType = iota
)

type Cost struct {
	Current int64 `json:"current"`

	CurrencyType *CurrencyType `json:"currencyType"`
	Limit        *Limit        `json:"limit,omitempty"`
}

type Limit struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
}
