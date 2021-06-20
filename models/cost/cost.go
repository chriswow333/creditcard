package cost

type CurrencyType int32

const (
	NTD CurrencyType = iota
)

type CurrentCost struct {
	Current  int64        `json:"current"`
	Currency CurrencyType `json:"currency"`
}

type CostLimit struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Descs []string `json:"descs"`

	Currency CurrencyType `json:"currency"`
	AtLeast  int64        `json:"atLeast"`
	AtMost   int64        `json:"atMost"`
}
