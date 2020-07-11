package repository

import (
	"context"
	"footcer-backend/model"
)

type OrderRepository interface {
	AddOrder(context context.Context, order model.Order) (model.Order, error)
}
