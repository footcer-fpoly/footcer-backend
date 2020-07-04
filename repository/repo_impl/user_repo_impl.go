package repo_impl

import (
	"context"
	"database/sql"
	"fmt"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/model/req"
	"footcer-backend/repository"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
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
			query := `INSERT INTO users(user_id, phone, email,role, display_name,birthday,position,level, avatar,verify,created_at, updated_at)
       VALUES(:user_id, :phone, :email,:role, :display_name,:birthday, :position,:level,:avatar, :verify, :created_at, :updated_at)`
			user.CreatedAt = time.Now()
			user.UpdatedAt = time.Now()
			_, err := u.sql.Db.NamedExecContext(context, query, userReq)
			return userReq, err
		}
	}
	return user, err

}

func (u UserRepoImpl) SelectById(context context.Context, userId string) (model.User, error) {
	var user model.User

	err := u.sql.Db.GetContext(context, &user,
		"SELECT * FROM users WHERE user_id = $1", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, message.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) SelectAll(context context.Context, userId string) ([]model.User, error) {
	panic("implement me")
}

func (u UserRepoImpl) Update(context context.Context, user model.User) (model.User, error) {
	sqlStatement := `
		UPDATE users
		SET 
			display_name  = (CASE WHEN LENGTH(:display_name) = 0 THEN display_name ELSE :display_name END),
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			avatar = (CASE WHEN LENGTH(:avatar) = 0 THEN avatar ELSE :avatar END),
			birthday = (CASE WHEN LENGTH(:birthday) = 0 THEN birthday ELSE :birthday END),
			position = (CASE WHEN LENGTH(:position) = 0 THEN position ELSE :position END),
			level = (CASE WHEN LENGTH(:level) = 0 THEN position ELSE :level END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE user_id    = :user_id
	`

	user.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == "23505" {
			log.Error(err.Error())
			return user, message.EmailExits
		}
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, message.UserNotUpdated
	}
	if count == 0 {
		return user, message.UserNotUpdated
	}

	return user, nil
}

func (u UserRepoImpl) ValidPhone(context context.Context, phoneReq string) error {
	var role string
	queryUserExits := `SELECT role FROM users WHERE users.phone = $1`

	err := u.sql.Db.GetContext(context, &role, queryUserExits, phoneReq)
	if err == sql.ErrNoRows {
		return nil
	}
	if role == "0" {
		return message.UserConflict

	}
	if role == "1" {
		return message.UserIsAdmin
	}
	return message.SomeWentWrong

}

func (u UserRepoImpl) CreateForPhone(context context.Context, user model.User) (model.User, error) {
	query := `INSERT INTO users(user_id, phone, email,password,role, display_name,birthday,position,level, avatar,verify,created_at, updated_at)
     VALUES(:user_id, :phone, :email,:password,:role, :display_name,:birthday, :position,:level,:avatar, :verify, :created_at, :updated_at)`

	_, err := u.sql.Db.NamedExecContext(context, query, user)
	if err != nil {
		log.Error(err.Error())
		return user, message.SignUpFail
	}
	if user.Role == 1 {
		var stadiumId = uuid.NewV1().String()

		var stadium = model.Stadium{
			StadiumId:   stadiumId,
			StadiumName: "Sân bóng mẫu",
			Address:     "123",
			Description: "",
			Image:       "example.jpg",
			PriceNormal: 0,
			PricePeak:   1,
			StartTime:   "5:30",
			EndTime:     "10:00",
			Category:    "Sân cỏ nhân tạo",
			Latitude:    0,
			Longitude:   1,
			Ward:        "",
			District:    "",
			City:        "",
			UserId:      user.UserId,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		queryCreateStadium := `INSERT INTO stadium(
		stadium_id, name_stadium, address, description, image, price_normal, price_peak, start_time, end_time, category, latitude, longitude, ward, district, city, user_id, created_at, updated_at)
		VALUES (:stadium_id, :name_stadium, :address, :description, :image, :price_normal, :price_peak, :start_time, :end_time, :category, :latitude, :longitude, :ward, :district, :city , :user_id, :created_at, :updated_at)`

		_, err := u.sql.Db.NamedExecContext(context, queryCreateStadium, stadium)

		if err != nil {
			log.Error(err.Error())
			return user, message.SignUpFail
		}
		var stadiumCollage = model.StadiumCollage{
			StadiumCollageId:   uuid.NewV1().String(),
			NameStadiumCollage: "Sân số 1",
			AmountPeople:       "5",
			StadiumId:          stadiumId,
			CreatedAt:          time.Time{},
			UpdatedAt:          time.Time{},
		}
		queryCreateStadiumCollage := `INSERT INTO public.stadium_collage(
		stadium_collage_id, name_stadium_collage, amount_people, stadium_id, created_at, updated_at)
		VALUES (:stadium_collage_id, :name_stadium_collage, :amount_people, :stadium_id, :created_at, :updated_at);`
		_, errCreateStadiumCollage := u.sql.Db.NamedExecContext(context, queryCreateStadiumCollage, stadiumCollage)
		if errCreateStadiumCollage != nil {
			log.Error(errCreateStadiumCollage.Error())
		}
	}
	return user, err
}

func (u UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error) {
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
