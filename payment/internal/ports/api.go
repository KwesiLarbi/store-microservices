package ports

import (
	"context"
	"github.com/kwesilarbi/store-microservices/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}