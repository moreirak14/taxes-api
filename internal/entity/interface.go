package entity


type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotalTransactions() (int, error)
}
