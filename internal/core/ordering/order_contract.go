package ordering

import "Qexchange/internal/core/ordering/models"

type OrderCoreContract interface {
	Buy(int, int, float64) (int, error)
	Sell(int, int, float64) (int, error)
	Cancel(int, int) error
}

type AutoOrderContract interface {
	StartPolling(int, int, float64, string, float64)
}

type OrderDBContract interface {
	GetUserBalance(int, int) (float64, error)
	GetCoinPrice(int) (float64, error)
	GetCoinCommission(int) (float64, error)
	CreateOrder(*models.Order) (int, error)
	SubmitOrder(*models.Order)
	ChangeOrderStatus(*models.Order, string) error
}
