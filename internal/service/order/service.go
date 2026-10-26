package order

import (
	"fmt"
	"nichigovsky/solid/internal/repository/model"
)

type OrderService struct {
	repo     OrderRepo
	notifier Notifier
}

func NewOrderSystem(repo OrderRepo, notifier Notifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}

func (o *OrderService) CreateOrder(customer string, products []string, total float64) error {
	err := o.repo.Initialize()
	if err != nil {
		return fmt.Errorf("Ошибка при инициализации бд: %w", err)
	}

	order := &model.Order{
		Customer: customer,
		Products: products,
		Total:    total,
		Status:   "pending",
	}

	err = o.repo.Save(order)
	if err != nil {
		return fmt.Errorf("Ошибка при сохранении записи: %w", err)
	}

	err = o.notifier.Send(customer)
	if err != nil {
		return fmt.Errorf("Ошибка при отправке уведомления: %w", err)
	}

	return nil
}
