package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/repository"
)

type AdminRepoImpl struct {
	sql *db.Sql
}

func NewAdminRepo(sql *db.Sql) repository.AdminRepository {
	return &AdminRepoImpl{sql: sql}
}

func (a AdminRepoImpl) AcceptStadium(context context.Context, id string) error {
	sqlStatement := `UPDATE stadium
		SET verify = $1 WHERE stadium_id = $2`

	result, err := a.sql.Db.ExecContext(context, sqlStatement, "1", id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}
	if count == 0 {
		log.Error(err.Error())

		return message.SomeWentWrong
	}
	return nil
}
