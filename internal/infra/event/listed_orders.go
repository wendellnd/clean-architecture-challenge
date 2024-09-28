package event

import "time"

type ListedOrders struct {
	Name    string
	Payload interface{}
}

func NewListedOrders() *ListedOrders {
	return &ListedOrders{
		Name: "ListedOrders",
	}
}

func (e *ListedOrders) GetName() string {
	return e.Name
}

func (e *ListedOrders) GetPayload() interface{} {
	return e.Payload
}

func (e *ListedOrders) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ListedOrders) GetDateTime() time.Time {
	return time.Now()
}
