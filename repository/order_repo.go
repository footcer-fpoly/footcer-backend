package repository

import (
	"context"
	"footcer-backend/model"
)

type OrderRepository interface {
	AddOrder(context context.Context, order model.Order) (model.Order, error)
	AcceptOrder(context context.Context, order model.Order) error
	RefuseOrder(context context.Context, order model.Order) error
	FinishOrder(context context.Context, order model.Order) error
	ListOrderForStadium(context context.Context, stadiumId string) (interface{}, error)
	ListOrderForUser(context context.Context, userId string) (interface{}, error)
}
