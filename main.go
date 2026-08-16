package main

import (
	"database/sql"
	"log"

	"example/solid/internal/notifier"
	"example/solid/internal/repository"
	"example/solid/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 1. Подключение к БД
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	// 2. Создание репозитория (работа с БД)
	repo := repository.NewSQLiteOrderRepository(db)

	// 3. Инициализация схемы БД
	if err := repo.InitSchema(); err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	log.Println("=== Демонстрация пункта 4: OrderService ===")
	log.Println()

	// ============================================
	// ДЕМОНСТРАЦИЯ: OrderService с разными отправителями
	// ============================================

	// 4. Создаем Email отправитель
	emailSender := notifier.NewEmailSender()

	// 5. Создаем OrderService с EmailSender
	orderServiceWithEmail := service.NewOrderService(repo, emailSender)

	log.Println("--- Тест 1: OrderService с Email уведомлением ---")
	err = orderServiceWithEmail.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Println("Ошибка:", err)
	}
	log.Println()

	// 6. Создаем SMS отправитель
	smsSender := notifier.NewSMSSender()

	// 7. Создаем OrderService с SMSSender
	orderServiceWithSMS := service.NewOrderService(repo, smsSender)

	log.Println("--- Тест 2: OrderService с SMS уведомлением ---")
	err = orderServiceWithSMS.CreateOrder("Анатолий", []string{"orange", "grape"}, 15.75)
	if err != nil {
		log.Println("Ошибка:", err)
	}
	log.Println()

}
