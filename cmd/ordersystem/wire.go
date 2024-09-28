//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/entity"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/infra/database"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/infra/web"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/usecase"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	//event.NewOrderCreated,
	//event.NewListedOrders,
	//wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	//wire.Bind(new(events.EventInterface), new(*event.ListedOrders)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

/*
var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setListedOrdersEvent = wire.NewSet(
	event.NewListedOrders,
	wire.Bind(new(events.EventInterface), new(*event.ListedOrders)),
)
*/

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		//setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		//setListedOrdersEvent,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		//setOrderCreatedEvent,
		//setListedOrdersEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
