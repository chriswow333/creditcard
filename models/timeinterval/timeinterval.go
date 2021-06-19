package timeinterval

type TimeInterval struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Desc []string `json:"desc"`

	DayFrom string `json:"day"`
	DayTo   string `json:"dayTo"`

	WeekDayFrom string `json:"weekDay"`
	WeekDayTo   string `json:"weekDayTo"`

	HourFrom string `json:"hour"`
	HourTo   string `json:"hourTo"`

	MinuteFrom string `json:"minute"`
	MinuteTo   string `json:"minuteTo"`
}
