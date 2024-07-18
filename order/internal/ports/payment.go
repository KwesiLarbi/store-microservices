package ports

import "github.com/kwesilarbi/store-microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}