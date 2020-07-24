package repo_impl

import (
	"context"
	"database/sql"
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
	game_id, date, hour, type, score, description_game, finish, stadium_id, team_id_host, team_id_guest, game_created_at, game_updated_at)
	VALUES (:game_id, :date, :hour, :type, :score, :description_game, :finish, (CASE WHEN LENGTH(:stadium_id) = 0 THEN null ELSE :stadium_id END) , :team_id_host, null, :game_created_at, :game_updated_at);`
	if len(game.StadiumId) > 0 {
		//game.StadiumId = sql.NullString{}
	}

	_, err := g.sql.Db.NamedExecContext(context, queryCreateGame, game)
	if err != nil {
		log.Error(err.Error())
		return game, message.SomeWentWrong
	}
	return game, nil
}

func (g *GameRepoImpl) JoinGame(context context.Context, gameTemp model.GameTemp) error {

	queryJoinGame := `INSERT INTO public.game_temp(
	game_temp_id, game_id, team_id)
	VALUES (:game_temp_id, :game_id, :team_id);`
	_, err := g.sql.Db.NamedExecContext(context, queryJoinGame, gameTemp)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
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

func (g *GameRepoImpl) GetGames(context context.Context, date string) (interface{}, error) {

	var listGame = []ListGame{}
	if (date == "all") {
		sqlSearch := `SELECT game.game_id, game.date, game.hour, game.type, game.score, game.description_game, game.finish,
 COALESCE(game.stadium_id,'null') stadium_id,  game_created_at, game_updated_at,COALESCE(stadium.name_stadium, '') name_stadium,
  game.team_id_host, COALESCE(game.team_id_guest, 'null') team_id_guest,team_host.name AS team_name_host,team_host.avatar AS team_avatar_host,
  COALESCE(team_guest.name , 'null')  team_name_guest,COALESCE(team_guest.avatar ,'null')  team_avatar_guest FROM public.game 
	LEFT JOIN stadium ON stadium.stadium_id = game.stadium_id 
	INNER JOIN team AS team_host ON team_host.team_id = game.team_id_host 
	LEFT JOIN team AS team_guest ON team_guest.team_id = game.team_id_guest;`
		err := g.sql.Db.SelectContext(context, &listGame, sqlSearch)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Error(err.Error())
				return listGame, message.NotData
			}
			log.Error(err.Error())
			return listGame, err
		}
	} else {
		sqlSearchDate := `SELECT game.game_id, game.date, game.hour, game.type, game.score, game.description_game, 
	game.finish, COALESCE(game.stadium_id,'null') stadium_id,  game_created_at, game_updated_at,COALESCE(stadium.name_stadium, '') name_stadium, 
	game.team_id_host, COALESCE(game.team_id_guest, 'null') team_id_guest,team_host.name AS team_name_host,
	team_host.avatar AS team_avatar_host,COALESCE(team_guest.name , 'null')  team_name_guest,COALESCE(team_guest.avatar ,'')  team_avatar_guest FROM public.game 
	LEFT JOIN stadium ON stadium.stadium_id = game.stadium_id 
	INNER JOIN team AS team_host ON team_host.team_id = game.team_id_host 
	LEFT JOIN team AS team_guest ON team_guest.team_id = game.team_id_guest  WHERE date = $1;`
		err := g.sql.Db.SelectContext(context, &listGame, sqlSearchDate, date)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Error(err.Error())
				return listGame, message.NotData
			}
			log.Error(err.Error())
			return listGame, err
		}

	}

	return listGame, nil

}
func (g *GameRepoImpl) GetGame(context context.Context, gameId string) (interface{}, error) {
	var game = ListGame{}

	sqlGetGame := `SELECT game.game_id, game.date, game.hour, game.type, game.score, game.description_game, 
	game.finish, COALESCE(game.stadium_id,'null') stadium_id,  game_created_at, game_updated_at,COALESCE(stadium.name_stadium, 'null') name_stadium, 
	game.team_id_host, COALESCE(game.team_id_guest, 'null') team_id_guest,team_host.name AS team_name_host,
	team_host.avatar AS team_avatar_host,COALESCE(team_guest.name , 'null')  team_name_guest,COALESCE(team_guest.avatar ,'null')  team_avatar_guest FROM public.game 
	LEFT JOIN stadium ON stadium.stadium_id = game.stadium_id 
	INNER JOIN team AS team_host ON team_host.team_id = game.team_id_host 
	LEFT JOIN team AS team_guest ON team_guest.team_id = game.team_id_guest  WHERE game_id = $1;`
	err := g.sql.Db.GetContext(context, &game, sqlGetGame, gameId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return game, message.NotData
		}
		log.Error(err.Error())
		return game, err
	}
	inviteTeam := game.TeamIdGuest == "null"
	if inviteTeam {
		var inviteTeams = []TeamTemp{}
		sqlGetTeamInvate := `SELECT team.team_id AS team_id_temp,team.name AS team_name_temp,team.avatar AS team_avatar_temp 
	FROM public.game_temp INNER JOIN team ON team.team_id = game_temp.team_id WHERE game_id =$1;`

		err := g.sql.Db.SelectContext(context, &inviteTeams, sqlGetTeamInvate, gameId)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Error(err.Error())
				return game, message.NotData
			}
			log.Error(err.Error())
			return game, err
		}
		game.ArrayTeamTemp = inviteTeams


	}

	return game, nil

}

type TeamHost struct {
	Name   string `json:"teamNameHost,omitempty" db:"team_name_host,omitempty"`
	Avatar string `json:"teamAvatarHost,omitempty" db:"team_avatar_host,omitempty"`
}

type TeamGuest struct {
	Name   string `json:"teamNameGuest,omitempty" db:"team_name_guest,omitempty"`
	Avatar string `json:"teamAvatarGuest,omitempty" db:"team_avatar_guest,omitempty"`
}

type TeamTemp struct {
	TeamId string `json:"teamIdTemp,omitempty" db:"team_id_temp,omitempty"`
	Name   string `json:"teamNameTemp,omitempty" db:"team_name_temp,omitempty"`
	Avatar string `json:"teamAvatarTemp,omitempty" db:"team_avatar_temp,omitempty"`
}
type ArrayTeamTemp [] TeamTemp
type ListGame struct {
	model.Game
	model.Stadium `json:"stadium"`
	TeamHost      `json:"teamHost"`
	TeamGuest     `json:"teamGuest"`
	ArrayTeamTemp `json:"teamInvite"`
}
