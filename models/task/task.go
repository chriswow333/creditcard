package task

import "example.com/creditcard/models/label"

type TaskType int32

const (
	NONE    TaskType = iota + 1
	WEEKDAY          // WEEKDAY duration
	LABEL            // 通路標籤
	// CHANNEL                // 通路
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
	Label        *label.Label  `json:"label,omitempty"`
	// ChannelLabelLimit *ChannelLabelLimit `json:"channelLabelLimit,omitempty"`
	// ChannelLimit      *ChannelLimit      `json:"channelLimit,omitempty"`
}

type WeekDayLimit struct {
	WeekDays []int `json:"weekDays"`
}

// type ChannelLabelLimit struct {
// 	ChannelLabels []int32 `json:"channelLabels"`
// }

// type ChannelLimit struct {
// 	Channels []int32 `json:"channels"`
// }
