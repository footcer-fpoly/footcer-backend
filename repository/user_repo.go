package repository

import (
	"context"
	"footcer-backend/model"
)

type UserRepository interface {
	CheckLogin(context context.Context, loginReq model.ReqSignIn) (model.User, error)
	ValidPhone(context context.Context, phone string) error
	CreateForPhone(context context.Context, user model.User) (model.User, error)
	Create(context context.Context, user model.User) (model.User, error)
	SelectById(context context.Context, userId string) (model.User, error)
	SelectAll(context context.Context, userId string) ([]model.User, error)
	Update(context context.Context, user model.User) error
}
