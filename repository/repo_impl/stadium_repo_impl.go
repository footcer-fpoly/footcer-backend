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

func (s StadiumRepoImpl) StadiumInfo(context context.Context, userId string) (interface{}, error) {

	type StadiumInfo struct {
		model.Stadium
		ArrayStadiumCollage `json:"stadium_collage"`
		ArrayStadiumReview  `json:"review"`
		model.User          `json:"user"`
	}
	stadium := StadiumInfo{}

	query := `SELECT stadium.stadium_id, stadium.name_stadium, stadium.address, stadium.description, stadium.image,  stadium.start_time, stadium.end_time, stadium.category, stadium.latitude, stadium.longitude, stadium.ward, stadium.district, stadium.city,stadium.time_peak,stadium.time_order,stadium.user_id,users.display_name,users.avatar,users.phone ,stadium.created_at, stadium.updated_at
	FROM public.stadium INNER JOIN users ON users.user_id = stadium.user_id  WHERE stadium.user_id =  $1`
	err := s.sql.Db.GetContext(context, &stadium,
		query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return stadium, message.StadiumNotFound
		}
		log.Error(err.Error())
		return stadium, err
	}

	//stadium collage
	var stadiumColl = []model.StadiumCollage{}
	queryColl := `SELECT stadium_collage_id, name_stadium_collage, amount_people, price_normal, price_peak,stadium_id, created_at, updated_at
	FROM public.stadium_collage WHERE stadium_id = $1`
	errColl := s.sql.Db.SelectContext(context, &stadiumColl, queryColl, stadium.StadiumId)
	if errColl != nil {
		if errColl == sql.ErrNoRows {
			log.Error(errColl.Error())
			return stadiumColl, message.StadiumNotFound
		}
		log.Error(errColl.Error())
		return stadiumColl, errColl
	}
	stadium.ArrayStadiumCollage = stadiumColl

	//review
	var review = []model.Review{}

	queryReview := `SELECT review_id, content, rate, stadium_id, review.user_id, review.created_at, review.updated_at, users.display_name,users.avatar
FROM public.review INNER JOIN users ON review.user_id = users.user_id WHERE review.stadium_id = $1;`
	errReview := s.sql.Db.SelectContext(context, &review, queryReview, stadium.StadiumId)
	if errReview != nil {
		if errReview == sql.ErrNoRows {
			log.Error(errReview.Error())
			return review, message.StadiumNotFound
		}
		log.Error(errReview.Error())
		return review, errReview
	}
	var sumRate float64 = 0
	if len(review) > 0 {
		for _, rate := range review {
			sumRate += rate.Rate
		}
		if sumRate > 0 {
			stadium.Stadium.RateCount = sumRate / float64(len(review))
		}
	}

	stadium.ArrayStadiumReview = review

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
			time_peak = (CASE WHEN LENGTH(:time_peak) = 0 THEN time_peak ELSE :time_peak END),
			time_order = (CASE WHEN LENGTH(:time_order) = 0 THEN time_order ELSE :time_order END),
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
		return stadium, message.StadiumNotUpdated
	}
	if count == 0 {
		return stadium, message.StadiumNotUpdated
	}

	return stadium, nil
}

func (s StadiumRepoImpl) StadiumCollageUpdate(context context.Context, stadiumColl model.StadiumCollage) (model.StadiumCollage, error) {
	sqlStatement := `
		UPDATE stadium_collage
		SET 
			name_stadium_collage  = (CASE WHEN LENGTH(:name_stadium_collage) = 0 THEN name_stadium_collage ELSE :name_stadium_collage END),
			amount_people = (CASE WHEN LENGTH(:amount_people) = 0 THEN amount_people ELSE :amount_people END),
			price_normal = (CASE WHEN LENGTH(:price_normal) = 0 THEN price_normal ELSE :price_normal END),
			price_peak = (CASE WHEN LENGTH(:price_peak) = 0 THEN price_peak ELSE :price_peak END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE stadium_collage_id    = :stadium_collage_id
	`

	stadiumColl.UpdatedAt = time.Now()

	result, err := s.sql.Db.NamedExecContext(context, sqlStatement, stadiumColl)
	if err != nil {
		log.Error(err.Error())
		return stadiumColl, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return stadiumColl, message.StadiumNotUpdated
	}
	if count == 0 {
		return stadiumColl, message.StadiumNotUpdated
	}

	return stadiumColl, nil
}

func (s StadiumRepoImpl) StadiumCollageAdd(context context.Context, stadiumColl model.StadiumCollage) (model.StadiumCollage, error) {

	queryCreate := `INSERT INTO public.stadium_collage(
	stadium_collage_id, name_stadium_collage, amount_people,price_normal, price_peak,stadium_id, created_at, updated_at)
	VALUES (:stadium_collage_id, :name_stadium_collage, :amount_people, :price_normal,:price_peak,:stadium_id, :created_at, :updated_at)`

	_, err := s.sql.Db.NamedExecContext(context, queryCreate, stadiumColl)
	if err != nil {
		log.Error(err.Error())
		return stadiumColl, message.StadiumNotUpdated
	}
	return stadiumColl, nil

}

//TODO chưa hợp lí -> xử lí sau
type ArrayStadiumCollage []model.StadiumCollage
type ArrayStadiumReview []model.Review
