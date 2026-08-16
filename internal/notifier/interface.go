package notifier

// Notifier определяет контракт для отправки уведомлений
type Notifier interface {
	Send(customer string) error
}
