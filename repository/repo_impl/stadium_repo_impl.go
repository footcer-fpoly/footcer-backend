package repo_impl

import (
	"database/sql"
	"fmt"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"footcer-backend/security/pro"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"sort"
	"strconv"
	"strings"
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

	query := `SELECT stadium.stadium_id, stadium.name_stadium, stadium.address, stadium.description, stadium.image, stadium.category, stadium.latitude, stadium.longitude, stadium.ward, stadium.district, stadium.city, stadium.user_id,users.display_name,users.avatar,users.phone ,stadium.created_at, stadium.updated_at
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
	queryColl := `SELECT stadium_collage_id, stadium_id, name_stadium_collage, amount_people, start_time, end_time, play_time, created_at, updated_at
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

	//stadium details
	var stadiumDet = []model.StadiumDetails{}
	queryDet := `SELECT stadium_collage_id, stadium_id, name_stadium_collage, amount_people, start_time, end_time, play_time, created_at, updated_at
	FROM public.stadium_collage WHERE stadium_id = $1`
	errDet := s.sql.Db.SelectContext(context, &stadiumDet, queryDet, stadium.StadiumId)
	if errDet != nil {
		if errDet == sql.ErrNoRows {
			log.Error(errDet.Error())
			return stadiumDet, message.StadiumNotFound
		}
		log.Error(errColl.Error())
		return stadiumDet, errColl
	}




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
			splitRate := strings.Split(fmt.Sprintf("%.1f", sumRate/float64(len(review))), ".")
			stadium.Stadium.RateCount = CeilRate(splitRate)
		}
	}

	stadium.ArrayStadiumReview = review

	return stadium, nil

}

func (s StadiumRepoImpl) StadiumInfoForID(context context.Context, stadiumID string) (interface{}, error) {

	type StadiumInfo struct {
		model.Stadium
		ArrayStadiumCollage `json:"stadium_collage"`
		ArrayStadiumReview  `json:"review"`
		model.User          `json:"user"`
	}
	stadium := StadiumInfo{}

	query := `SELECT stadium.stadium_id, stadium.name_stadium, stadium.address, stadium.description, stadium.image,  stadium.start_time, stadium.end_time, stadium.category, stadium.latitude, stadium.longitude, stadium.ward, stadium.district, stadium.city,stadium.time_peak,stadium.user_id,users.display_name,users.avatar,users.phone ,stadium.created_at, stadium.updated_at
	FROM public.stadium INNER JOIN users ON users.user_id = stadium.user_id  WHERE stadium.stadium_id =  $1`
	err := s.sql.Db.GetContext(context, &stadium,
		query, stadiumID)
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
			splitRate := strings.Split(fmt.Sprintf("%.1f", sumRate/float64(len(review))), ".")
			stadium.Stadium.RateCount = CeilRate(splitRate)
		}
	}

	stadium.ArrayStadiumReview = review

	//get order time
	var timeOrder = []string{}

	queryTimeOrder := `SELECT time_slot
FROM public.orders WHERE accept = $1 AND finish = $2 AND stadium_id = $3;`
	errTimeOrder := s.sql.Db.SelectContext(context, &timeOrder, queryTimeOrder, "1", "0", stadium.StadiumId)
	if errTimeOrder != nil {
		if errTimeOrder == sql.ErrNoRows {
			log.Error(errTimeOrder.Error())
			return review, message.StadiumNotFound
		}
		log.Error(errTimeOrder.Error())
		return review, errTimeOrder
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
			time_peak = (CASE WHEN LENGTH(:time_peak) = 0 THEN time_peak ELSE :time_peak END),
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
	stadium_collage_id, stadium_id, name_stadium_collage, amount_people,start_time, end_time , play_time,created_at, updated_at)
	VALUES (:stadium_collage_id, :stadium_id, :name_stadium_collage, :amount_people,:start_time,:end_time,:play_time, :created_at, :updated_at);`

	_, err := s.sql.Db.NamedExecContext(context, queryCreate, stadiumColl)
	if err != nil {
		log.Error(err.Error())
		return stadiumColl, message.SomeWentWrong
	}
	return stadiumColl, nil

}

func (s StadiumRepoImpl) SearchStadiumLocation(context context.Context, latitude string, longitude string) ([]model.Stadium, error) {
	var stadium = []model.Stadium{}

	querySQL := `SELECT stadium_id, name_stadium, address, description, image, start_time, end_time, category, latitude, longitude, ward, district, city, time_peak, user_id, created_at, updated_at
	FROM public.stadium;`
	err := s.sql.Db.SelectContext(context, &stadium, querySQL)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return stadium, message.StadiumNotFound
		}
		log.Error(err.Error())
		return stadium, err
	}
	for i, v := range stadium {
		c, err := maps.NewClient(maps.WithAPIKey(pro.GOOGLE_MAP_KEY))
		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}
		latitudeStadium := strconv.FormatFloat(v.Latitude, 'f', -1, 64)
		longitudeStadium := strconv.FormatFloat(v.Longitude, 'f', -1, 64)
		locationStadium := latitudeStadium + "," + longitudeStadium
		locationClient := latitude + "," + longitude
		r := &maps.DistanceMatrixRequest{
			Origins:       []string{locationStadium},
			Destinations:  []string{locationClient},
			Units:         maps.UnitsImperial,
			Language:      "en",
			DepartureTime: "now",
		}
		route, err := c.DistanceMatrix(context, r)

		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}
		//print(route.Rows[0].Elements[0].Distance.Meters / 1000)
		stadium[i].Distance = route.Rows[0].Elements[0].Distance.Meters / 1000
		stadium[i].Timer = int(route.Rows[0].Elements[0].DurationInTraffic.Minutes())

		//review
		var review = []float64{}

		queryReview := `SELECT rate FROM review WHERE review.stadium_id = $1;`
		errReview := s.sql.Db.SelectContext(context, &review, queryReview, stadium[i].StadiumId)
		if errReview != nil {
			if errReview == sql.ErrNoRows {
				log.Error(errReview.Error())
			}
			log.Error(errReview.Error())
			return nil, message.SomeWentWrong
		}
		var sumRate float64 = 0
		if len(review) > 0 {
			for _, rate := range review {
				sumRate += rate
			}
			if sumRate > 0 {
				splitRate := strings.Split(fmt.Sprintf("%.1f", sumRate/float64(len(review))), ".")
				stadium[i].RateCount = CeilRate(splitRate)
			}
		}
	}

	sort.Slice(stadium, func(i, j int) bool {
		return stadium[i].Distance < stadium[j].Distance
	})

	return stadium, nil
}

func (s StadiumRepoImpl) SearchStadiumName(context context.Context, name string) ([]model.Stadium, error) {

	var stadium = []model.Stadium{}

	querySQL := `SELECT stadium_id, name_stadium, address, description, image, start_time, end_time, category, latitude, longitude, ward, district, city, time_peak, user_id, created_at, updated_at
	FROM public.stadium WHERE name_stadium ILIKE $1`
	err := s.sql.Db.SelectContext(context, &stadium, querySQL, "%"+name+"%")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err.Error())
			return stadium, message.StadiumNotFound
		}
		log.Error(err.Error())
		return stadium, err
	}
	return stadium, nil

}

func (s StadiumRepoImpl) StadiumDetailsAdd(context context.Context, stadiumDetails model.StadiumDetails) (model.StadiumDetails, error) {
	queryCreate := `INSERT INTO public.stadium_details(
	stadium_detail_id, stadium_collage_id, start_time_detail, end_time_detail, price, description, has_order, created_at, updated_at)
	VALUES (:stadium_detail_id, :stadium_collage_id, :start_time_detail, :end_time_detail, :price, :description, :has_order, :created_at, :updated_at);`

	_, err := s.sql.Db.NamedExecContext(context, queryCreate, stadiumDetails)
	if err != nil {
		log.Error(err.Error())
		return stadiumDetails, message.SomeWentWrong
	}
	return stadiumDetails, nil
}

func CeilRate(rate []string) float64 {
	var rateCeil float64 = 0.0
	rate1, _ := strconv.Atoi(rate[0])
	rate2, _ := strconv.Atoi(rate[1])

	if rate2 >= 0 && rate2 <= 2 {
		rateCeil = float64(rate1)
	} else if rate2 >= 3 && rate2 <= 7 {
		rateCeil = float64(rate1) + 0.5
	} else {
		rateCeil = float64(rate1) + 1

	}
	return rateCeil
}

//TODO chưa hợp lí -> xử lí sau
type ArrayStadiumCollage []model.StadiumCollage
type ArrayStadiumReview []model.Review

/**
1-2 -> 0
3-7 -> 5
8- -> 0+
*/
