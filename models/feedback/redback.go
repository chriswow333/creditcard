package feedback

type RedCalculateType int32

const (
	RED_TIMES = iota + 1
)

type Redback struct {
	RedCalculateType RedCalculateType `json:"redCalculateType"`

	Min int64 `json:"min"`
	Max int64 `json:"max"`

	Times int64 `json:"times"`
}

type RedReturn struct {
	IsRedGet     bool  `json:"isRedGet"`
	RedbackTimes int64 `json:"redbackTimes"`

	ActualUseCash int64   `json:"actualUseCash"`
	ActualRedback float64 `json:"actualRedback"`

	CurrentCash int64   `json:"current"`
	TotalCash   float64 `json:"total"`
}
