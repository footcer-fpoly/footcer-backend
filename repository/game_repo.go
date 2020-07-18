package repository

import (
	"context"
	"footcer-backend/model"
)
type GameRepository interface {
	AddGame(context context.Context, game model.Game) (model.Game, error)
}