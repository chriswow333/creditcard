package timeinterval

type TimeInterval struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Desc []string `json:"desc"`

	DayFrom int32 `json:"day"`
	DayTo   int32 `json:"dayTo"`

	WeekDayFrom int32 `json:"weekDay"`
	WeekDayTo   int32 `json:"weekDayTo"`

	HourFrom int32 `json:"hour"`
	HourTo   int32 `json:"hourTo"`

	MinuteFrom int32 `json:"minute"`
	MinuteTo   int32 `json:"minuteTo"`
}
