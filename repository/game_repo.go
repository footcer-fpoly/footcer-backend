package repository

import (
	"context"
	"footcer-backend/model"
)

type GameRepository interface {
	AddGame(context context.Context, game model.Game) (model.Game, error)

	JoinGame(context context.Context, gameTemp model.GameTemp) error

	AcceptJoin(context context.Context, gameTemp model.GameTemp) error

	RefuseJoin(context context.Context, gameTemp model.GameTemp) error

	GetGames(context context.Context, date string) (interface{}, error)

	GetGame(context context.Context, gameId string) (interface{}, error)

	UpdateScore(context context.Context, game model.Game) (interface{}, error)

	GetGameForUser(context context.Context, userId string) (interface{}, error)

	//SearchGameForDate(context context.Context, var date) (interface{}, error)
}
