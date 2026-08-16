package repository

import (
	"database/sql"
	"fmt"
)

// SQLiteOrderRepository - реализация RepositoryWriter для SQLite
type SQLiteOrderRepository struct {
	db *sql.DB
}

// NewSQLiteOrderRepository - конструктор
func NewSQLiteOrderRepository(db *sql.DB) *SQLiteOrderRepository {
	return &SQLiteOrderRepository{db: db}
}

// SaveOrder реализует интерфейс RepositoryWriter
func (r *SQLiteOrderRepository) SaveOrder(customer string, products []string, total float64) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		customer, fmt.Sprintf("%v", products), total, "pending",
	)
	return err
}

// InitSchema - инициализация схемы базы данных
func (r *SQLiteOrderRepository) InitSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer TEXT NOT NULL,
		products TEXT NOT NULL,
		total REAL NOT NULL,
		status TEXT NOT NULL
	)`
	_, err := r.db.Exec(query)
	return err
}
