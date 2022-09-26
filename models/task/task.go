package task

type TaskType int32

const (
	NONE    TaskType = iota + 1
	WEEKDAY          // WEEKDAY duration
)

type Task struct {
	ID string `json:"id"`

	Name   string   `json:"name"`
	Descs  []string `json:"descs"`
	CardID string   `json:"cardID"`

	TaskType      TaskType       `json:"taskType"`
	TaskTypeModel *TaskTypeModel `json:"taskTypeModel,omitempty"`

	DefaultPass bool `json:"defaultPass"`
}

type TaskTypeModel struct {
	WeekDayLimit *WeekDayLimit `json:"weekDayLimit,omitempty"`
}

type WeekDayLimit struct {
	WeekDays []int `json:"weekDays"`
}
