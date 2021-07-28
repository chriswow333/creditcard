package common

type ValidateTime struct {
	StartTime int64 `json:"startDate"`
	EndTime   int64 `json:"endDate"`
}

type OperatorType int32

const (
	AND OperatorType = iota
	OR
	Union // 全部都要符合
	Cross // 其中一個符合
)
