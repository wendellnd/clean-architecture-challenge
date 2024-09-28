package usecase

import (
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/entity"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/infra/event"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/pkg/events"
)

type ListOrderOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	orderRepository entity.OrderRepositoryInterface,
	eventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
		EventDispatcher: eventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.ListOrders()
	if err != nil {
		return ListOrderOutputDTO{}, err
	}

	var result []OrderOutputDTO
	for _, order := range orders {
		result = append(result, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	listedOrderEvent := event.NewListedOrders()
	listedOrderEvent.SetPayload(result)
	c.EventDispatcher.Dispatch(listedOrderEvent)

	return ListOrderOutputDTO{
		Orders: result,
	}, nil
}
