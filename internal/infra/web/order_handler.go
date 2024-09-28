package web

import (
	"encoding/json"
	"net/http"

	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/entity"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/internal/usecase"
	"github.com/wendellnd/graduate-go-expert-classes/Clean_Architecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher events.EventDispatcherInterface
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher: EventDispatcher,
		OrderRepository: OrderRepository,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository, h.EventDispatcher)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
