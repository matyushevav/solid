package notifier

import "fmt"

// EmailSender - структура для отправки email уведомлений
type EmailSender struct {
	// Здесь должны добавить поля для SMTP настроек. Пока заглушка
	// Например: smtpHost, smtpPort, username, password и т.д.
}

// NewEmailSender - конструктор для создания EmailSender
func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

// Send - реализует интерфейс Notifier
// Отправляет уведомление клиенту по email
func (e *EmailSender) Send(customer string) error {
	// Здесь должна быть реальная логика отправки email. Пока заглушка
	fmt.Printf("📧 Email уведомление отправлено клиенту %s\n", customer)
	return nil
}
