package repository

// RepositoryWriter определяет контракт для записи заказов в хранилище
type RepositoryWriter interface {
    SaveOrder(customer string, products []string, total float64) error
}