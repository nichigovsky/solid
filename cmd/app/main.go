package main

import (
	"database/sql"
	"log"
	"nichigovsky/solid/internal/notifier"
	"nichigovsky/solid/internal/repository/sqlite"
	"nichigovsky/solid/internal/service/order"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := sqlite.NewSqLiteOrderSystem(db)
	repo.Initialize()

	emailSender := notifier.NewEmailSender()

	orderSystemEmail := order.NewOrderSystem(repo, emailSender)

	err = orderSystemEmail.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)

	if err != nil {
		log.Println("Ошибка:", err)
		return
	}

	smsSender := notifier.NewSMSSender()

	orderSystemSms := order.NewOrderSystem(repo, smsSender)

	err = orderSystemSms.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)

	if err != nil {
		log.Println("Ошибка:", err)
		return
	}
}
