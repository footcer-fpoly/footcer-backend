package repo_impl

import (
	"context"
	"database/sql"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type NotificationRepoImpl struct {
	sql *db.Sql
}

func NewNotificationRepo(sql *db.Sql) repository.NotificationRepository {
	return &NotificationRepoImpl{sql: sql}
}

func (n *NotificationRepoImpl) AddNotification(context context.Context, notification model.Notification) (model.Notification, error) {
	queryCreateOrder := `INSERT INTO public.notifications(
	notify_id, key, title, content, icon, general_id, user_id, created_at_notify, updated_at_notify)
	VALUES (:notify_id, :key, :title, :content, :icon, :general_id, :user_id, :created_at_notify, :updated_at_notify);`
	_, err := n.sql.Db.NamedExecContext(context, queryCreateOrder, notification)
	if err != nil {
		log.Error(err.Error())
		return notification, err
	}
	return notification, nil
}
func (n *NotificationRepoImpl) GetNotification(context context.Context, userId string) ([]model.Notification, error) {
	var notification []model.Notification

	query := `select * from notifications WHERE user_id = $1 ORDER BY updated_at_notify DESC;`
	err := n.sql.Db.SelectContext(context, &notification, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return notification, message.StadiumNotFound
		}
		log.Error(err.Error())
		return notification, err
	}
	return notification, nil

}
