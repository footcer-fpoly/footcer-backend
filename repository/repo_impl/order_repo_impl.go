package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) repository.OrderRepository {
	return &OrderRepoImpl{
		sql: sql,
	}
}

func (o OrderRepoImpl) AddOrder(context context.Context, order model.Order) (model.Order, error) {
	panic("implement me")
}
