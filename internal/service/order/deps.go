package order

import (
	"nichigovsky/solid/internal/repository/model"
)

type OrderRepo interface {
	Initialize() error
	Save(order *model.Order) error
}

type Notifier interface {
	Send(customer string) error
}
