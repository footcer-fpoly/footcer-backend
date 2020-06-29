package repository

import (
	"context"
	"footcer-backend/model"
)

type StadiumRepository interface {
	StadiumInfo(context context.Context, userId string) (model.Stadium, error)
}
