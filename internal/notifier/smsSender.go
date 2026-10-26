package notifier

import "fmt"

type SMSSender struct{}

func (s *SMSSender) Send(customer string) error {
	fmt.Printf("Уведомление отправлено клиенту %s\n", customer)
	return nil
}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}
