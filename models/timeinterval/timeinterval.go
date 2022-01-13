package timeinterval

type TimeInterval struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	DayFrom int32 `json:"day,omitempty"`
	DayTo   int32 `json:"dayTo,omitempty"`

	WeekDayFrom int32 `json:"weekDay,omitempty"`
	WeekDayTo   int32 `json:"weekDayTo,omitempty"`

	HourFrom int32 `json:"hour,omitempty"`
	HourTo   int32 `json:"hourTo,omitempty"`

	MinuteFrom int32 `json:"minute,omitempty"`
	MinuteTo   int32 `json:"minuteTo,omitempty"`
}
