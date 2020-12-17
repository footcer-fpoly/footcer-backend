package repo_impl

import (
	"context"
	"database/sql"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"strings"
)

type StatisticsRepoImpl struct {
	sql *db.Sql
}

func NewStatisticsRepo(sql *db.Sql) repository.StatisticsRepository {
	return &StatisticsRepoImpl{sql: sql}
}

func (s StatisticsRepoImpl) StatisticsDay(context context.Context, date string, userId string) (model.Statistics, error) {
	var statistics = model.Statistics{}
	var stadium = model.Stadium{}
	statistics.Date = date
	queryStadiumID := `
	SELECT stadium.stadium_id FROM stadium
	INNER JOIN users ON users.user_id = stadium.user_id
 	WHERE users.user_id = $1`
	errDet := s.sql.Db.GetContext(context, &stadium, queryStadiumID, userId)
	if errDet != nil {
		if errDet == sql.ErrNoRows {
			log.Error(errDet.Error())
			return statistics, message.StadiumNotFound
		}
		log.Error(errDet.Error())
		return statistics, errDet
	}

	queryTotalDetails := `
	SELECT COUNT(stadium_detail_id) as total_details FROM stadium_details as details
	INNER JOIN stadium_collage as collage ON details.stadium_collage_id = collage.stadium_collage_id
	INNER JOIN stadium ON stadium.stadium_id = collage.stadium_id
 	WHERE stadium.stadium_id = $1`
	errDet = s.sql.Db.GetContext(context, &statistics, queryTotalDetails, stadium.StadiumId)
	if errDet != nil {
		log.Error(errDet.Error())
		return statistics, errDet
	}

	queryTotalDetailsOrder := `
	SELECT COUNT(o.order_id) as total_details_order, COUNT(DISTINCT o.user_id) as total_customer,  COALESCE(SUM(o.price), 0) as total_price
	FROM public.stadium_details as details 
	INNER JOIN orders as o ON details.stadium_detail_id = o.stadium_detail_id
	INNER join orders_status  on orders_status.order_id = o.order_id 
	INNER JOIN stadium_collage as collage ON details.stadium_collage_id = collage.stadium_collage_id
	INNER JOIN stadium ON stadium.stadium_id = collage.stadium_id
	WHERE stadium.stadium_id = $1
	AND orders_status.status LIKE $2 
	AND CAST(time as DATE) = CAST($3 AS DATE)`
	errDetailsOrder := s.sql.Db.GetContext(context, &statistics, queryTotalDetailsOrder, stadium.StadiumId, "%ACCEPT%", date)
	if errDetailsOrder != nil {
		log.Error(errDetailsOrder.Error())
		return statistics, errDetailsOrder
	}

	return statistics, nil
}

func (s StatisticsRepoImpl) StatisticsMonth(context context.Context, date string, userId string) (model.Statistics, error) {
	var statistics = model.Statistics{}
	var stadium = model.Stadium{}
	statistics.Date = date
	queryStadiumID := `
	SELECT stadium.stadium_id FROM stadium
	INNER JOIN users ON users.user_id = stadium.user_id
 	WHERE users.user_id = $1`
	errDet := s.sql.Db.GetContext(context, &stadium, queryStadiumID, userId)
	if errDet != nil {
		if errDet == sql.ErrNoRows {
			log.Error(errDet.Error())
			return statistics, message.StadiumNotFound
		}
		log.Error(errDet.Error())
		return statistics, errDet
	}

	queryTotalDetails := `
	SELECT COUNT(stadium_detail_id) as total_details FROM stadium_details as details
	INNER JOIN stadium_collage as collage ON details.stadium_collage_id = collage.stadium_collage_id
	INNER JOIN stadium ON stadium.stadium_id = collage.stadium_id
 	WHERE stadium.stadium_id = $1`
	errDet = s.sql.Db.GetContext(context, &statistics, queryTotalDetails, stadium.StadiumId)
	if errDet != nil {
		log.Error(errDet.Error())
		return statistics, errDet
	}

	queryTotalDetailsOrder := `
	SELECT COUNT(o.order_id) as total_details_order, COUNT(DISTINCT o.user_id) as total_customer,  COALESCE(SUM( o.price), 0 )as total_price
	FROM public.stadium_details as details 
	INNER JOIN orders as o ON details.stadium_detail_id = o.stadium_detail_id
	INNER join orders_status  on orders_status.order_id = o.order_id 
	INNER JOIN stadium_collage as collage ON details.stadium_collage_id = collage.stadium_collage_id
	INNER JOIN stadium ON stadium.stadium_id = collage.stadium_id
	WHERE stadium.stadium_id = $1
	AND orders_status.status LIKE $2 
	AND to_char( time, 'yyyy-mm' ) = $3`
	errDetailsOrder := s.sql.Db.GetContext(context, &statistics, queryTotalDetailsOrder, stadium.StadiumId, "%ACCEPT%", date)
	if errDetailsOrder != nil {
		log.Error(errDetailsOrder.Error())
		return statistics, errDetailsOrder
	}

	return statistics, nil
}

func (s StatisticsRepoImpl) StatisticsFromTo(context context.Context, dates string, userId string) ([]model.Statistics, error) {
	date := strings.Split(dates, ",")
	var statisticss []model.Statistics

	for i := 0; i < len(date); i++ {
		println(date[i])
		var statistics = model.Statistics{}
		var stadium = model.Stadium{}
		statistics.Date =  date[i]
		queryStadiumID := `
	SELECT stadium.stadium_id FROM stadium
	INNER JOIN users ON users.user_id = stadium.user_id
 	WHERE users.user_id = $1`
		errDet := s.sql.Db.GetContext(context, &stadium, queryStadiumID, userId)
		if errDet != nil {
			if errDet == sql.ErrNoRows {
				log.Error(errDet.Error())
				return statisticss, message.StadiumNotFound
			}
			log.Error(errDet.Error())
			return statisticss, errDet
		}

		queryTotalDetails := `
	SELECT COUNT(stadium_detail_id) as total_details FROM stadium_details as details
	INNER JOIN stadium_collage as collage ON details.stadium_collage_id = collage.stadium_collage_id
	INNER JOIN stadium ON stadium.stadium_id = collage.stadium_id
 	WHERE stadium.stadium_id = $1`
		errDet = s.sql.Db.GetContext(context, &statistics, queryTotalDetails, stadium.StadiumId)
		if errDet != nil {
			log.Error(errDet.Error())
			return statisticss, errDet
		}

		queryTotalDetailsOrder := `
	SELECT COUNT(o.order_id) as total_details_order, COUNT(DISTINCT o.user_id) as total_customer,  COALESCE(SUM(o.price), 0) as total_price
	FROM public.stadium_details as details 
	INNER JOIN orders as o ON details.stadium_detail_id = o.stadium_detail_id
	INNER join orders_status  on orders_status.order_id = o.order_id 
	INNER JOIN stadium_collage as collage ON details.stadium_collage_id = collage.stadium_collage_id
	INNER JOIN stadium ON stadium.stadium_id = collage.stadium_id
	WHERE stadium.stadium_id = $1
	AND orders_status.status LIKE $2 
	AND CAST(time as DATE) = CAST($3 AS DATE)`
		errDetailsOrder := s.sql.Db.GetContext(context, &statistics, queryTotalDetailsOrder, stadium.StadiumId, "%ACCEPT%", date[i])
		if errDetailsOrder != nil {
			log.Error(errDetailsOrder.Error())
			return statisticss, errDetailsOrder
		}
		statisticss = append(statisticss, statistics)

	}
	return statisticss, nil

}
