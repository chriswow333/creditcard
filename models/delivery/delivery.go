package delivery

import "example.com/creditcard/models/action"

type Delivery struct {
	ID         string            `json:"id"`
	Name       string            `json:"name,omitempty"`
	ActionType action.ActionType `json:"actionType"`
	Desc       string            `json:"desc,omitempty"`
}
