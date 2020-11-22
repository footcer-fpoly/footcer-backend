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
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
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

	queryUserExits := `SELECT * FROM users WHERE users.phone = $1`

	user.Email = ""
	user.TokenNotify = ""
	err := u.sql.Db.GetContext(context, &user, queryUserExits, userReq.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("New user -> Insert Data")
			query := `INSERT INTO users(user_id, phone,password, email,role, display_name,birthday,position,level, avatar,verify,token_notify,created_at, updated_at)
       VALUES(:user_id, :phone, :password,:email,:role, :display_name,:birthday, :position,:level,:avatar, :verify, :token_notify ,:created_at, :updated_at)`
			user.CreatedAt = time.Now()
			user.UpdatedAt = time.Now()
			_, err := u.sql.Db.NamedExecContext(context, query, userReq)
			if err != nil {
				log.Error(err.Error())
			}
			return userReq, err
		} else {
			fmt.Println("User Exits -> Return User")
			return user, nil
		}
	}
	fmt.Println("User Error -> Return User")

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
			level = (CASE WHEN LENGTH(:level) = 0 THEN level ELSE :level END),
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

func (u UserRepoImpl) ValidPhone(context context.Context, phoneReq string) (int, model.User, error) {
	var user model.User
	queryUserExits := `SELECT * FROM users WHERE users.phone = $1`

	err := u.sql.Db.GetContext(context, &user, queryUserExits, phoneReq)
	if err == sql.ErrNoRows {
		return 200, user, nil
	}
	if user.Role == 0 {
		return 209, user, message.UserConflict
	}
	if user.Role == 1 {
		return 203, user, message.UserIsAdmin
	}

	return 409, user, message.SomeWentWrong

}

func (u UserRepoImpl) CreateForPhone(context context.Context, user model.User) (model.User, error) {
	query := `INSERT INTO users(user_id, phone, email,password,role, display_name,birthday,position,level, avatar,verify,token_notify,created_at, updated_at)
     VALUES(:user_id, :phone, :email,:password,:role, :display_name,:birthday, :position,:level,:avatar, :verify, :token_notify,:created_at, :updated_at)`

	user.Token = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, query, user)
	if err != nil {
		log.Error(err.Error())
		if err != sql.ErrNoRows {
			return user, message.UserConflict
		}
		return user, message.SignUpFail
	}
	//var stadiumRepo = StadiumRepoImpl{sql: u.sql}

	if user.Role == 1 {
		var stadiumId = uuid.NewV1().String()

		var stadium = model.Stadium{
			StadiumId:   stadiumId,
			StadiumName: "Sân bóng mẫu",
			Address:     "01 Đường Tô Kí, Quận 12, Tp.HCM",
			Description: "Sân cỏ nhân tao",
			Image:       "http://footcer.tk:4000/static/stadium/example.jpg",
			Category:    "Sân cỏ nhân tạo",
			Latitude:    -1,
			Longitude:   -1,
			Ward:        "",
			District:    "",
			City:        "",
			UserId:      user.UserId,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		queryCreateStadium := `INSERT INTO public.stadium(
	stadium_id, user_id, name_stadium, address, description, image, category, latitude, longitude, ward, district, city, created_at, updated_at)
	VALUES 
	(:stadium_id, :user_id, :name_stadium, :address, :description, :image, :category, :latitude, :longitude, :ward, :district, :city , :created_at, :updated_at);`

		_, err := u.sql.Db.NamedExecContext(context, queryCreateStadium, stadium)

		if err != nil {
			log.Error(err.Error())
			queryDeleteUser := `DELETE FROM public.users WHERE user_id = $1`
			row, err := u.sql.Db.ExecContext(context, queryDeleteUser, user.UserId)
			if err != nil {
				log.Error(err.Error())
			}
			count, _ := row.RowsAffected()
			if count == 0 {
				log.Error(err.Error())
			}
			return user, message.SignUpFail
		}

		//var stadiumCollageId = uuid.NewV4().String()
		//
		//var stadiumCollage = model.StadiumCollage{
		//	StadiumCollageId:   stadiumCollageId,
		//	NameStadiumCollage: "Sân số 1",
		//	AmountPeople:       "5",
		//	StartTime:          "19800000", //5:30
		//	EndTime:            "79200000", //22:30
		//	PlayTime:           "5400000",  // 90'
		//	StadiumId:          stadiumId,
		//	CreatedAt:          time.Now(),
		//	UpdatedAt:          time.Now(),
		//}
		//
		//_, errCreateStadiumCollage := stadiumRepo.StadiumCollageAdd(context, stadiumCollage)
		//if errCreateStadiumCollage != nil {
		//	log.Error(errCreateStadiumCollage.Error())
		//
		//	queryDeleteStadium := `DELETE FROM public.stadium WHERE stadium_id = $1`
		//	row, err := u.sql.Db.ExecContext(context, queryDeleteStadium, stadiumCollage.StadiumId)
		//	if err != nil {
		//		log.Error(err.Error())
		//	}
		//	count, _ := row.RowsAffected()
		//	if count == 0 {
		//		log.Error(err.Error())
		//	}
		//
		//	queryDeleteUser := `DELETE FROM public.users WHERE user_id = $1`
		//	row, err = u.sql.Db.ExecContext(context, queryDeleteUser, user.UserId)
		//	if err != nil {
		//		log.Error(err.Error())
		//	}
		//	count, _ = row.RowsAffected()
		//	if count == 0 {
		//		log.Error(err.Error())
		//	}
		//	return user, message.SignUpFail
		//
		//}
		//startTime, _ := strconv.ParseInt(stadiumCollage.StartTime, 10, 64)
		//endTime, _ := strconv.ParseInt(stadiumCollage.EndTime, 10, 64)
		//playTime, _ := strconv.ParseInt(stadiumCollage.PlayTime, 10, 64)
		//
		//start := startTime
		//end := 0
		//
		//amountTimer := math.Floor(float64((endTime - startTime) / playTime))
		//
		//for i := 0; i < int(amountTimer); i++ {
		//	end = int(start + playTime)
		//	var stadiumDetails = model.StadiumDetails{
		//		StadiumDetailsId: uuid.NewV1().String(),
		//		StadiumCollageId: stadiumCollageId,
		//		StartTimeDetails: strconv.Itoa(int(start)),
		//		EndTimeDetails:   strconv.Itoa(end),
		//		Price:            stadiumCollage.DefaultPrice,
		//		Description:      "",
		//		HasOrder:         false,
		//		CreatedAt:        time.Now(),
		//		UpdatedAt:        time.Now(),
		//	}
		//
		//	_, errCreateStadiumDetails := stadiumRepo.StadiumDetailsAdd(context, stadiumDetails)
		//	if errCreateStadiumDetails != nil {
		//		log.Error(errCreateStadiumDetails.Error())
		//		//
		//		queryDeleteStadiumCollage := `DELETE FROM public.stadium_collage WHERE stadium_collage_id = $1`
		//		row, err := u.sql.Db.ExecContext(context, queryDeleteStadiumCollage, stadiumCollageId)
		//		if err != nil {
		//		}
		//		count, _ := row.RowsAffected()
		//		if count == 0 {
		//			log.Error(err.Error())
		//		}
		//
		//		queryDeleteStadium := `DELETE FROM public.stadium WHERE stadium_id = $1`
		//		row, err = u.sql.Db.ExecContext(context, queryDeleteStadium, stadiumId)
		//		if err != nil {
		//		}
		//		count, _ = row.RowsAffected()
		//		if count == 0 {
		//			log.Error(err.Error())
		//		}
		//
		//		queryDeleteUser := `DELETE FROM public.users WHERE user_id = $1`
		//		row, err = u.sql.Db.ExecContext(context, queryDeleteUser, user.UserId)
		//		if err != nil {
		//		}
		//		count, _ = row.RowsAffected()
		//		if count == 0 {
		//			log.Error(err.Error())
		//		}
		//
		//		return user, message.SignUpFail
		//
		//	}
		//
		//	start = int64(end)
		//}

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

func (u UserRepoImpl) ValidEmail(context context.Context, emailReq string) (model.User, error) {
	var user model.User
	queryUserExits := `SELECT * FROM users WHERE users.email = $1`

	err := u.sql.Db.GetContext(context, &user, queryUserExits, emailReq)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if user.Role == 0 {
		return user, message.UserConflict
	}
	if user.Role == 1 {
		return user, message.UserIsAdmin
	}
	return user, message.SomeWentWrong

}

func (u UserRepoImpl) ValidUUID(context context.Context, uuidReq string) (model.User, error) {
	var user model.User
	queryUserExits := `SELECT * FROM users WHERE users.user_id = $1`

	err := u.sql.Db.GetContext(context, &user, queryUserExits, uuidReq)

	if err == sql.ErrNoRows {
		return user, nil
	}
	if user.Role == 0 {
		return user, message.UserConflict

	}
	if user.Role == 1 {
		return user, message.UserIsAdmin
	}
	return user, message.SomeWentWrong

}

func (u UserRepoImpl) UpdatePassword(context context.Context, user model.User) error {
	sqlStatement := `
		UPDATE users
		SET 
			password = (CASE WHEN LENGTH(:password) = 0 THEN password ELSE :password END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE phone    = :phone
	`

	user.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return message.UserNotUpdated
	}
	if count == 0 {
		return message.UserNotUpdated
	}

	return nil
}

func (u UserRepoImpl) DeleteUser(context context.Context, phone string) error {
	sqlStatement := `
		DELETE FROM public.users
	WHERE phone = $1;
	`

	result, err := u.sql.Db.ExecContext(context, sqlStatement, phone)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}
	if count == 0 {
		return message.SomeWentWrong
	}

	return nil
}
