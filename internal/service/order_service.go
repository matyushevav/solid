package service

import (
	"fmt"
	"log"

	"example/solid/internal/notifier"
	"example/solid/internal/repository"
)

// OrderService - сервис для управления заказами
// Зависит от интерфейсов, а не от конкретных реализаций
type OrderService struct {
	repo     repository.RepositoryWriter
	notifier notifier.Notifier
}

// NewOrderService - конструктор с внедрением зависимостей
func NewOrderService(repo repository.RepositoryWriter, notifier notifier.Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

// CreateOrder - создание нового заказа
func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	// 1. Валидация входных данных
	if customer == "" {
		return fmt.Errorf("Имя клиента не может быть пустым")
	}
	if len(products) == 0 {
		return fmt.Errorf("Список продуктов не может быть пустым")
	}
	if total <= 0 {
		return fmt.Errorf("Общая сумма должна быть > 0")
	}

	// 2. Сохраняем заказ в БД через репозиторий (интерфейс)
	err := s.repo.SaveOrder(customer, products, total)
	if err != nil {
		return fmt.Errorf("Ошибка при сохранении заказа: %w", err)
	}

	// 3. Отправляем уведомление через нотификатор (интерфейс)
	err = s.notifier.Send(customer)
	if err != nil {
		// Логируем ошибку, но не отменяем создание заказа
		log.Printf("Ошибка отправки уведомления клиенту %s: %v", customer, err)
	}

	return nil
}
