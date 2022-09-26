package payload

import (
	"example.com/creditcard/models/channel"
	"example.com/creditcard/models/feedback"
)

type Payload struct {
	ID string `json:"id"`

	Feedback *feedback.Feedback `json:"feedback,omitempty"`

	Channel *channel.Channel `json:"channel,omitempty"` // 通路 任務 限制條件(限額等)
}
