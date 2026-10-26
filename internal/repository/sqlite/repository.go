package sqlite

import (
	"database/sql"
	"fmt"
	"nichigovsky/solid/internal/repository/model"
)

type orderSystem struct {
	db *sql.DB
}

func NewSqLiteOrderSystem(db *sql.DB) *orderSystem {
	return &orderSystem{db: db}
}

func (s *orderSystem) Initialize() error {
	_, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer TEXT NOT NULL,
        products TEXT NOT NULL,
        total REAL NOT NULL,
        status TEXT NOT NULL
    )`)

	return err
}

func (s *orderSystem) Save(order *model.Order) error {
	// Создание заказа в БД
	_, err := s.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer, fmt.Sprintf("%v", order.Products), order.Total, "pending",
	)

	return err
}
