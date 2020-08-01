package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type NotificationRepoImpl struct {
	sql *db.Sql
}


func NewNotificationRepo(sql *db.Sql) repository.NotificationRepository {
	return &NotificationRepoImpl{sql: sql}
}

func (NotificationRepoImpl) AddNotification(context context.Context, notification model.Notification) (model.Notification, error) {
	panic("implement me")
	return  notification,nil
}
func (NotificationRepoImpl) GetNotification(context context.Context) ([]model.Notification, error) {
	panic("implement me")
	var notification = []model.Notification{}
	return  notification,nil

}