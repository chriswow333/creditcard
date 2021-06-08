package delivery

import "example.com/creditcard/models/action"

type Delivery struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Action action.Action `json:"action"`
	Desc   string        `json:"desc"`
}
