package repository

import (
	"context"
	"footcer-backend/model"
	"footcer-backend/model/req"
)

type UserRepository interface {
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	ValidPhone(context context.Context, phone string) (int, model.User, error)
	CreateForPhone(context context.Context, user model.User) (model.User, error)
	Create(context context.Context, user model.User) (model.User, error)
	SelectById(context context.Context, userId string) (model.User, error)
	SelectAll(context context.Context, userId string) ([]model.User, error)
	Update(context context.Context, user model.User) (model.User, error)
	ValidEmail(context context.Context, email string) (model.User, error)
	ValidUUID(context context.Context, uuid string) (model.User, error)
	UpdatePassword(context context.Context, user model.User) error
	DeleteUser(context context.Context, phone string) error

	UpdateTokenNotify(context context.Context, tokenNotify req.TokenNotify) error
	GetToken(context context.Context, userId string) (string, error)
	GetTokenForTeam(context context.Context, teamId string) ([]model.User, error)
}
