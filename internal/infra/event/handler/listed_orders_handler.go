package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/pkg/events"
)

type ListedOrdersHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewListedOrdersHandler(rabbitMQChannel *amqp.Channel) *ListedOrdersHandler {
	return &ListedOrdersHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *ListedOrdersHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Listed orders: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
