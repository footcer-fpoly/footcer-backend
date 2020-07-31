package repository

import (
	"context"
	"footcer-backend/model"
)

type ServiceRepository interface {
	AddService(context context.Context, service model.Service) (model.Service, error)
	DeleteService(context context.Context, serviceId string) error
}
