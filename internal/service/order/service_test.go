package order

import (
	"errors"
	"testing"

	"nichigovsky/solid/internal/repository/model"
)

type MockRepo struct {
	SavedOrder *model.Order
	RepoErr    error
}

func (m *MockRepo) Initialize() error {
	return nil
}

func (m *MockRepo) Save(order *model.Order) error {
	m.SavedOrder = order
	return m.RepoErr
}

type MockNotifier struct {
	Customer string
}

func (m *MockNotifier) Send(customer string) error {
	m.Customer = customer
	return nil
}

func TestCreateOrder(t *testing.T) {
	repo := &MockRepo{}
	notifier := &MockNotifier{Customer: "Ольга"}

	orderSystem := NewOrderSystem(repo, notifier)

	err := orderSystem.CreateOrder(notifier.Customer, []string{"apple, watermelon"}, 26.2)

	if err != nil {
		t.Errorf("Ожидалось nil, получена ошибка %v", err)
	}
}

func TestCreateOrderError(t *testing.T) {
	repo := &MockRepo{RepoErr: errors.New("Ошибка")}
	notifier := &MockNotifier{Customer: "Ольга"}

	orderSystem := NewOrderSystem(repo, notifier)

	err := orderSystem.CreateOrder(notifier.Customer, []string{"apple, watermelon"}, 26.2)

	if err == nil {
		t.Errorf("Ожидалась ошибка: %v, получен nil ", err)
	}
}
