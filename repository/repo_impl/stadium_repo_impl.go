package repo_impl

import (
	"context"
	"database/sql"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"time"
)

type StadiumRepoImpl struct {
	sql *db.Sql
}

func NewStadiumRepo(sql *db.Sql) repository.StadiumRepository {
	return &StadiumRepoImpl{sql: sql}
}

func (s StadiumRepoImpl) StadiumInfo(context context.Context, userId string) (model.Stadium, error) {
	stadium := model.Stadium{}

	query := `SELECT stadium.stadium_id, stadium.name_stadium, stadium.address, stadium.description, stadium.image, stadium.price_normal, stadium.price_peak, stadium.start_time, stadium.end_time, stadium.category, stadium.latitude, stadium.longitude, stadium.ward, stadium.district, stadium.city, stadium.user_id,users.display_name,users.avatar,users.phone ,stadium.created_at, stadium.updated_at
	FROM public.stadium INNER JOIN users ON users.user_id = stadium.user_id WHERE stadium.user_id =  $1`
	err := s.sql.Db.GetContext(context, &stadium,
		query, userId)

	if err != nil {
		if err == sql.ErrNoRows{
			log.Error(err.Error())
			return stadium, message.StadiumNotFound
		}
		log.Error(err.Error())
		return stadium, err
	}

	return stadium, nil

}

func (s StadiumRepoImpl) StadiumUpdate(context context.Context, stadium model.Stadium) (model.Stadium, error) {
	sqlStatement := `
		UPDATE stadium
		SET 
			name_stadium  = (CASE WHEN LENGTH(:name_stadium) = 0 THEN name_stadium ELSE :name_stadium END),
			address = (CASE WHEN LENGTH(:address) = 0 THEN address ELSE :address END),
			description = (CASE WHEN LENGTH(:description) = 0 THEN description ELSE :description END),
			image = (CASE WHEN LENGTH(:image) = 0 THEN image ELSE :image END),
			price_normal = (CASE WHEN LENGTH(:price_normal) = 0 THEN price_normal ELSE :price_normal END),
			price_peak = (CASE WHEN LENGTH(:price_peak) = 0 THEN price_peak ELSE :price_peak END),
			start_time = (CASE WHEN LENGTH(:start_time) = 0 THEN start_time ELSE :start_time END),
			end_time = (CASE WHEN LENGTH(:end_time) = 0 THEN end_time ELSE :end_time END),
			category = (CASE WHEN LENGTH(:category) = 0 THEN category ELSE :category END),
			latitude = (CASE WHEN LENGTH(:latitude) = 0 THEN latitude ELSE :latitude END),
			longitude = (CASE WHEN LENGTH(:longitude) = 0 THEN longitude ELSE :longitude END),
			ward = (CASE WHEN LENGTH(:ward) = 0 THEN ward ELSE :ward END),
			district = (CASE WHEN LENGTH(:district) = 0 THEN district ELSE :district END),
			city = (CASE WHEN LENGTH(:city) = 0 THEN city ELSE :city END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE user_id    = :user_id
	`

	stadium.UpdatedAt = time.Now()

	result, err := s.sql.Db.NamedExecContext(context, sqlStatement, stadium)
	if err != nil {
		log.Error(err.Error())
		return stadium, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return stadium, message.UserNotUpdated
	}
	if count == 0 {
		return stadium, message.UserNotUpdated
	}

	return stadium, nil
}
