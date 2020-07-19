package repository

import (
	"context"
	"footcer-backend/model"
)
type GameRepository interface {
	AddGame(context context.Context, game model.Game) (model.Game, error)

	JoinGame(context context.Context, gameTemp model.GameTemp)  error

	AcceptJoin(context context.Context, gameTemp model.GameTemp) error

	RefuseJoin(context context.Context, gameTemp model.GameTemp) error

	UpdateScore(context context.Context, game model.Game) (interface{}, error)

	//SearchGameForDate(context context.Context, var date) (interface{}, error)


}