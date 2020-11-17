package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"time"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) repository.OrderRepository {
	return &OrderRepoImpl{
		sql: sql,
	}
}

func (o OrderRepoImpl) AddOrder(context context.Context, order model.Order) (model.Order, error) {
	queryCreateOrder := `INSERT INTO public.orders (
	order_id, user_id, stadium_detail_id, "time", description, price, finish, accept, order_created_at, order_updated_at)
	VALUES ( :order_id , :user_id , :stadium_detail_id , :time , :description , :price , :finish , :accept , :order_created_at , :order_updated_at );`
	_, err := o.sql.Db.NamedExecContext(context, queryCreateOrder, order)
	if err != nil {
		log.Error(err.Error())
		return order, message.SomeWentWrong
	}
	return order, nil
}

func (o OrderRepoImpl) RefuseOrder(context context.Context, order model.Order) error {
	sqlStatement := `
		UPDATE orders
		SET 
		accept  = (CASE WHEN LENGTH(:accept) = 0 THEN accept ELSE :accept END),
			order_updated_at 	  = COALESCE (:order_updated_at, order_updated_at)
		WHERE order_id    = :order_id
	`

	order.UpdatedAt = time.Now()
	_, err := o.sql.Db.NamedExecContext(context, sqlStatement, order)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}
	return nil
}

func (o OrderRepoImpl) AcceptOrder(context context.Context, order model.Order) error {
	sqlStatement := `
		UPDATE orders
		SET 
		accept  = (CASE WHEN LENGTH(:accept) = 0 THEN accept ELSE :accept END),
			order_updated_at 	  = COALESCE (:order_updated_at, order_updated_at)
		WHERE order_id    = :order_id
	`

	order.UpdatedAt = time.Now()
	_, err := o.sql.Db.NamedExecContext(context, sqlStatement, order)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}

	sqlStatement = `
		UPDATE stadium_details
		SET has_order  = true,
			updated_at 	  = CURRENT_DATE
		WHERE stadium_detail_id    = :stadium_detail_id
	`

	order.UpdatedAt = time.Now()
	_, err = o.sql.Db.NamedExecContext(context, sqlStatement, order)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}

	return nil
}

func (o OrderRepoImpl) FinishOrder(context context.Context, order model.Order) error {
	sqlStatement := `
		UPDATE orders
		SET 
		finish  = (CASE WHEN LENGTH(:finish) = 0 THEN finish ELSE :finish END),
			order_updated_at 	  = COALESCE (:order_updated_at, order_updated_at)
		WHERE order_id    = :order_id
	`

	order.UpdatedAt = time.Now()
	_, err := o.sql.Db.NamedExecContext(context, sqlStatement, order)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}

	sqlStatement = `
		UPDATE stadium_details
		SET has_order  = false,
			updated_at 	  = CURRENT_DATE
		WHERE stadium_detail_id    = :stadium_detail_id
	`

	order.UpdatedAt = time.Now()
	_, err = o.sql.Db.NamedExecContext(context, sqlStatement, order)
	if err != nil {
		log.Error(err.Error())
		return message.SomeWentWrong
	}

	return nil
}

func (o OrderRepoImpl) ListOrderForStadium(context context.Context, stadiumId string) (interface{}, error) {
	type listOrders struct {
		model.Order
		model.Stadium        `json:"stadium"`
		model.StadiumCollage `json:"stadium_collage"`
		model.StadiumDetails `json:"stadium_details"`
		model.User           `json:"user"`
	}
	var orders = []listOrders{}
	sqlStatement := `SELECT orders.*,
	users.user_id,users.display_name,users.avatar,
	stadium_collage.name_stadium_collage,stadium_collage.amount_people,
	stadium.name_stadium,stadium.address,stadium.category, 
	stadium_details.price , stadium_details.start_time_detail , stadium_details.end_time_detail 
	FROM public.orders 
	INNER JOIN users ON users.user_id = orders.user_id 
	INNER JOIN stadium_details ON stadium_details.stadium_detail_id = orders.stadium_detail_id
	INNER JOIN stadium_collage  ON stadium_collage.stadium_collage_id = stadium_details.stadium_collage_id  
	INNER JOIN stadium ON stadium.stadium_id = stadium_collage.stadium_id WHERE stadium.stadium_id = $1;
	`

	err := o.sql.Db.SelectContext(context, &orders, sqlStatement, stadiumId)
	if err != nil {
		log.Error(err.Error())
		return orders, message.SomeWentWrong
	}
	return orders, nil
}

func (o OrderRepoImpl) ListOrderForUser(context context.Context, userId string) (interface{}, error) {
	type listOrders struct {
		model.Order
		model.Stadium        `json:"stadium"`
		model.StadiumCollage `json:"stadium_collage"`
		model.StadiumDetails `json:"stadium_details"`
		model.User           `json:"user"`
	}
	var orders = []listOrders{}
	sqlStatement := `
	SELECT orders.*,
	users.user_id,users.display_name,users.avatar,
	stadium_collage.name_stadium_collage,stadium_collage.amount_people,
	stadium.name_stadium,stadium.address,stadium.category, 
	stadium_details.price , stadium_details.start_time_detail , stadium_details.end_time_detail 
	FROM public.orders 
	INNER JOIN users ON users.user_id = orders.user_id 
	INNER JOIN stadium_details ON stadium_details.stadium_detail_id = orders.stadium_detail_id
	INNER JOIN stadium_collage  ON stadium_collage.stadium_collage_id = stadium_details.stadium_collage_id  
	INNER JOIN stadium ON stadium.stadium_id = stadium_collage.stadium_id WHERE users.user_id = $1;
	`

	err := o.sql.Db.SelectContext(context, &orders, sqlStatement, userId)
	if err != nil {
		log.Error(err.Error())
		return orders, message.SomeWentWrong
	}
	return orders, nil
}
