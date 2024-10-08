// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, eventDispatcher)
	return createOrderUseCase
}

func NewListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	orderRepository := database.NewOrderRepository(db)
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository, eventDispatcher)
	return listOrdersUseCase
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository)
	return webOrderHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))
