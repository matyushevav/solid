package notifier

import "fmt"

// SMSSender - структура для отправки SMS уведомлений
type SMSSender struct {
	// Здесь можно добавить поля для:
	// - API ключ SMS-сервиса
	// - Номер телефона отправителя
	// - URL API шлюза
	// Пока заглушка
}


// NewSMSSender - конструктор для создания SMSSender
func NewSMSSender() *SMSSender {
    return &SMSSender{}
}

// Send - реализует интерфейс Notifier
// Отправляет SMS уведомление клиенту
func (s *SMSSender) Send(customer string) error {
	// Здесь должна быть реальная логика отправки SMS. Пока заглушка
	fmt.Printf("📱 SMS уведомление отправлено клиенту %s\n", customer)
	return nil
}