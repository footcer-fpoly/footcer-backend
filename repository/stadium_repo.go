package repository

import (
	"context"
	"footcer-backend/model"
)

type StadiumRepository interface {
	StadiumInfo(context context.Context, userId string) (model.Stadium, error)
	StadiumUpdate(context context.Context, stadium model.Stadium) (model.Stadium, error)
}
