package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type GameRepoImpl struct {
	sql *db.Sql
}


func NewGameRepo(sql *db.Sql) repository.GameRepository {
	return &GameRepoImpl{sql: sql}
}

func (g * GameRepoImpl) AddGame(context context.Context, game model.Game) (model.Game, error) {
	panic("implement me")
}
