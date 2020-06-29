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

type StadiumRepoImpl struct {
	sql *db.Sql
}

func NewStadiumRepo(sql *db.Sql) repository.StadiumRepository {
	return &StadiumRepoImpl{sql: sql}
}

func (s StadiumRepoImpl) StadiumInfo(context context.Context, userId string) (model.Stadium, error) {
	var stadium model.Stadium

	err := s.sql.Db.GetContext(context, &stadium,
		"SELECT * FROM stadium WHERE user_id = $1", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return stadium, message.UserNotFound
		}
		log.Error(err.Error())
		return stadium, err
	}

	return stadium, nil

}
