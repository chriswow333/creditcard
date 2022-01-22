package timeinterval

type TimeType int32

const (
	WeekDay TimeType = iota
)

type TimeInterval struct {
	ID string `json:"id"`

	TimeType TimeType `json:"timeType"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	WeekDayFrom int32 `json:"weekDayFrom"`
	WeekDayTo   int32 `json:"weekDayTo"`

	DayFrom int32 `json:"dayFrom"`
	DayTo   int32 `json:"dayTo"`

	HourFrom int32 `json:"hourFrom"`
	HourTo   int32 `json:"hourTo"`

	MinuteFrom int32 `json:"minuteFrom"`
	MinuteTo   int32 `json:"minuteTo"`
}
