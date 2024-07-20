package api

import (
	"github.com/kwesilarbi/store-microservices/order/internal/application/core/domain"
	"github.com/kwesilarbi/store-microservices/order/internal/ports"
)

// The API depends on DBPort
type Application struct {
	db 		ports.DBPort
	payment ports.PaymentPort
}

// DBPort is passed during the app's initialization
func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db: db,
		payment: payment,
	}
}

// function saves order through the DB port
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}

	return order, nil
}