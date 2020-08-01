package repository

import (
	"context"
	"footcer-backend/model"
)

type NotificationRepository interface {
	AddNotification(context context.Context, notification model.Notification) (model.Notification, error)
	GetNotification(context context.Context) ([]model.Notification, error)
}
