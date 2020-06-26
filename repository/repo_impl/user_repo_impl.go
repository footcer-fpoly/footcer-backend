package repo_impl

import (
	"context"
	"database/sql"
	"fmt"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepository {
	return &UserRepoImpl{
		sql: sql,
	}
}
func (u UserRepoImpl) Create(context context.Context, userReq model.User) (model.User, error) {
	var user model.User

	queryUserExits := `SELECT * FROM users WHERE users.email = $1`

	err := u.sql.Db.GetContext(context, &user, queryUserExits, userReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("New user -> Insert Data")
			query := `INSERT INTO users(user_id, phone, email,role, display_name,birthday,position,level, avatar,verify)
       VALUES(:user_id, :phone, :email,:role, :display_name,:birthday, :position,:level,:avatar, :verify)`

			_, err := u.sql.Db.NamedExecContext(context, query, userReq)
			return userReq, err
		}
	}
	return user, err

}

func (u UserRepoImpl) SelectById(context context.Context, userId string) (model.User, error) {
	panic("implement me")
}

func (u UserRepoImpl) SelectAll(context context.Context, userId string) ([]model.User, error) {
	panic("implement me")
}

func (u UserRepoImpl) Update(context context.Context, user model.User) error {
	panic("implement me")
}
func (u UserRepoImpl) ValidPhone(context context.Context, phoneReq string) error {
	var phone string
	queryUserExits := `SELECT phone FROM users WHERE users.phone = $1`

	err := u.sql.Db.GetContext(context, &phone, queryUserExits, phoneReq)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
	}
	return err
}
func (u UserRepoImpl) CreateForPhone(context context.Context, user model.User) (model.User, error) {

	query := `INSERT INTO users(user_id, phone, email,password,role, display_name,birthday,position,level, avatar,verify)
       VALUES(:user_id, :phone, :email,:password,:role, :display_name,:birthday, :position,:level,:avatar, :verify)`

	_, err := u.sql.Db.NamedExecContext(context, query, user)
	return user, err
}
func (u UserRepoImpl) CheckLogin(context context.Context, loginReq model.ReqSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE phone=$1", loginReq.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, message.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}
