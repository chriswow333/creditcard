package supermarket

type ActionType int32

const (
	Shopping ActionType = iota
	Deposit
)

type Supermarket struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Action ActionType `json:"action"`
	Desc   string     `json:"desc"`
}
