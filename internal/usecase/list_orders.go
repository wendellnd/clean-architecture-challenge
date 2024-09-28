package usecase

import (
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/entity"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/pkg/events"
)

type ListOrderOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	ListedOrders    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	orderRepository entity.OrderRepositoryInterface,
	listedOrders events.EventInterface,
	eventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
		ListedOrders:    listedOrders,
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

	c.ListedOrders.SetPayload(result)
	c.EventDispatcher.Dispatch(c.ListedOrders)

	return ListOrderOutputDTO{
		Orders: result,
	}, nil
}
