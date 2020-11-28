package repository

import (
	"context"
	"footcer-backend/model"
)

type AdminRepository interface {
	AcceptStadium(context context.Context, id string) error
	Statistics(context context.Context) (model.StatisticsAdmin, error)
}
