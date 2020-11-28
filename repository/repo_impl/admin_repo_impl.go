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

func (a AdminRepoImpl) Statistics(context context.Context) (model.StatisticsAdmin, error) {
	var statistics = model.StatisticsAdmin{}
	queryUser := `SELECT COUNT(user_id) as total_user FROM users`

	errUser := a.sql.Db.GetContext(context, &statistics, queryUser)
	if errUser != nil {
		if errUser == sql.ErrNoRows {
			log.Error(errUser.Error())
			return statistics, message.StadiumNotFound
		}
		log.Error(errUser.Error())
		return statistics, errUser
	}

	queryStadium := `SELECT COUNT(stadium_id) as total_stadium FROM stadium`
	errStadium := a.sql.Db.GetContext(context, &statistics, queryStadium)
	if errStadium != nil {
		log.Error(errStadium.Error())
		return statistics, errStadium
	}

	queryOrder := `SELECT COUNT(order_id) as total_order FROM orders`
	errOrder := a.sql.Db.GetContext(context, &statistics, queryOrder)
	if errOrder != nil {
		log.Error(errOrder.Error())
		return statistics, errOrder
	}

	queryTeam := `SELECT COUNT(team_id) as total_team FROM team`
	errTeam := a.sql.Db.GetContext(context, &statistics, queryTeam)
	if errTeam != nil {
		log.Error(errTeam.Error())
		return statistics, errTeam
	}

	return statistics, nil
}
