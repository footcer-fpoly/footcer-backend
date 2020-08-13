package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type AdRepoImpl struct {
	sql *db.Sql
}

func NewAdRepo(sql *db.Sql) repository.AdRepository {
	return &AdRepoImpl{sql: sql}
}
func (a *AdRepoImpl) AddAd(context context.Context, ad model.Ad) (model.Ad, error) {
	panic("implement me")
}
