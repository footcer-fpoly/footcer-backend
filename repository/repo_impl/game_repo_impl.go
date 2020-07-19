package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"time"
)

type GameRepoImpl struct {
	sql *db.Sql
}

func NewGameRepo(sql *db.Sql) repository.GameRepository {
	return &GameRepoImpl{sql: sql}
}

func (g *GameRepoImpl) AddGame(context context.Context, game model.Game) (model.Game, error) {

	queryCreateGame := `INSERT INTO public.game(
	game_id, date, hour, type, score, description, finish, stadium_id, team_id_host, team_id_guest, game_created_at, game_updated_at)
	VALUES (:game_id, :date, :hour, :type, :score, :description, :finish, :stadium_id, :team_id_host, null, :game_created_at, :game_updated_at);`
	_, err := g.sql.Db.NamedExecContext(context, queryCreateGame, game)
	if err != nil {
		log.Error(err.Error())
		return game, message.SomeWentWrong
	}
	return game, nil
}

func (g *GameRepoImpl) JoinGame(context context.Context, gameTemp model.GameTemp)  error {

	queryJoinGame := `INSERT INTO public.game_temp(
	game_temp_id, game_id, team_id)
	VALUES (:game_temp_id, :game_id, :team_id);`
	_, err := g.sql.Db.NamedExecContext(context, queryJoinGame, gameTemp)
	if err != nil {
		log.Error(err.Error())
		return  message.SomeWentWrong
	}
	return nil
}

func (g *GameRepoImpl) AcceptJoin(context context.Context, gameTemp model.GameTemp) error {
	queryDelete := `DELETE FROM public.game_temp
	WHERE game_id = $1;`
	row, err := g.sql.Db.ExecContext(context, queryDelete, gameTemp.GameId)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}
	count, _ := row.RowsAffected()
	if count == 0 {
		return message.SomeWentWrong
	}
	var game = model.Game{}
	game.GameId = gameTemp.GameId
	game.TeamIdGuest = gameTemp.TeamId
	game.UpdatedAt = time.Now()
	sqlStatement := `
		UPDATE game
		SET 
		team_id_guest  = (CASE WHEN LENGTH(:team_id_guest) = 0 THEN team_id_guest ELSE :team_id_guest END),
			game_updated_at 	  = COALESCE (:game_updated_at, game_updated_at)
		WHERE game_id    = :game_id
	`

	_, errJoinGame := g.sql.Db.NamedExecContext(context, sqlStatement, game)
	if errJoinGame != nil {
		log.Error(errJoinGame.Error())
		return message.SomeWentWrong
	}

	return nil
}

func (g *GameRepoImpl) RefuseJoin(context context.Context, gameTemp model.GameTemp) error {
	queryDelete := `DELETE FROM public.game_temp
	WHERE game_temp_id = $1;`
	row, err := g.sql.Db.ExecContext(context, queryDelete, gameTemp.GameTempId)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}
	count, _ := row.RowsAffected()
	if count == 0 {
		return message.SomeWentWrong
	}
	return nil
}

func (g *GameRepoImpl) UpdateScore(context context.Context, game model.Game) (interface{}, error) {
	panic("implement me")
}
