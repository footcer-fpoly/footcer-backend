package repository

import (
	"context"
	"footcer-backend/model"
)

type StatisticsRepository interface {
	StatisticsDay(context context.Context, date string, userId string) (model.Statistics, error)
	StatisticsMonth(context context.Context, date string, userId string) (model.Statistics, error)
	StatisticsFromTo(context context.Context, dates string, userId string) ([]model.Statistics, error)
}
