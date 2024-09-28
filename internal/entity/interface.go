package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	ListOrders() (orders []Order, err error)
}
