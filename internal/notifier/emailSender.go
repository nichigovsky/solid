package notifier

import "fmt"

type EmailSender struct{}

func (e *EmailSender) Send(customer string) error {
	fmt.Printf("Уведомление отправлено клиенту %s\n", customer)
	return nil
}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}
