package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	uuid "github.com/satori/go.uuid"
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
	order_id, user_id, stadium_detail_id, "time", description, price, finish, order_created_at, order_updated_at)
	VALUES ( :order_id , :user_id , :stadium_detail_id , :time , :description , :price , :finish , :order_created_at , :order_updated_at );`
	_, err := o.sql.Db.NamedExecContext(context, queryCreateOrder, order)
	if err != nil {
		log.Error(err.Error())
		return order, message.SomeWentWrong
	}
	orderStatus := model.OrderStatus{
		OrderStatusId: uuid.NewV4().String(),
		OrderId:       order.OrderId,
		Status:        "WAITING",
		Reason:        "",
		IsUser:        true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	queryCreateOrderStatus := `INSERT INTO public.orders_status (
	order_status_id, order_id, status, reason, is_user , created_at, updated_at)
	VALUES ( :order_status_id , :orders_status.order_id , :status , :reason , :is_user ,:created_at , :updated_at  );`
	_, err = o.sql.Db.NamedExecContext(context, queryCreateOrderStatus, orderStatus)
	if err != nil {
		log.Error(err.Error())
		return order, message.SomeWentWrong
	}

	return order, nil
}

func (o OrderRepoImpl) UpdateStatusOrder(context context.Context, orderStatus model.OrderStatus) error {
	sqlStatement := `
		UPDATE orders_status
		SET 
		status  = (CASE WHEN LENGTH(:status) = 0 THEN status ELSE :status END),
		reason  = (CASE WHEN LENGTH(:reason) = 0 THEN reason ELSE :reason END),
		is_user  = (CASE WHEN LENGTH(:is_user) = 0 THEN is_user ELSE :is_user END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE orders_status.order_id    = :orders_status.order_id
	`

	orderStatus.UpdatedAt = time.Now()
	_, err := o.sql.Db.NamedExecContext(context, sqlStatement, orderStatus)
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

	return nil
}

func (o OrderRepoImpl) ListOrderForStadium(context context.Context, stadiumId string) (interface{}, error) {
	type listOrders struct {
		model.Order
		model.OrderStatus    `json:"order_status"`
		model.Stadium        `json:"stadium"`
		model.StadiumCollage `json:"stadium_collage"`
		model.StadiumDetails `json:"stadium_details"`
		model.User           `json:"user"`
	}
	var orders = []listOrders{}
	sqlStatement := `SELECT orders.*,
	users.user_id,users.display_name,users.avatar, users.phone,
	stadium_collage.name_stadium_collage,stadium_collage.amount_people,
	stadium.name_stadium,stadium.address,stadium.category, stadium.stadium_id, 
	stadium_details.price , stadium_details.start_time_detail , stadium_details.end_time_detail, orders_status.*
	FROM public.orders 
	INNER JOIN users ON users.user_id = orders.user_id 
	INNER JOIN stadium_details ON stadium_details.stadium_detail_id = orders.stadium_detail_id
	INNER JOIN stadium_collage  ON stadium_collage.stadium_collage_id = stadium_details.stadium_collage_id  
	INNER JOIN stadium ON stadium.stadium_id = stadium_collage.stadium_id 
	INNER JOIN orders_status  ON orders_status.order_id = orders.order_id
	WHERE stadium.stadium_id = $1 ORDER BY orders.time DESC;
	`

	err := o.sql.Db.SelectContext(context, &orders, sqlStatement, stadiumId)
	if err != nil {
		log.Error(err.Error())
		return orders, message.SomeWentWrong
	}
	for i := 0; i < len(orders); i++ {
		orders[i].Stadium.StadiumId = orders[i].StadiumCollage.StadiumId
		orders[i].StadiumCollage.StadiumCollageId = orders[i].StadiumDetails.StadiumCollageId
	}

	return orders, nil
}

func (o OrderRepoImpl) ListOrderForUser(context context.Context, userId string) (interface{}, error) {
	type listOrders struct {
		model.Order
		model.OrderStatus    `json:"order_status"`
		model.Stadium        `json:"stadium"`
		model.StadiumCollage `json:"stadium_collage"`
		model.StadiumDetails `json:"stadium_details"`
		model.User           `json:"user"`
	}
	var orders = []listOrders{}
	sqlStatement := `
	SELECT orders.*, stadium.image ,
	users.user_id,users.display_name,users.avatar, users.phone,stadium_collage.stadium_collage_id,
	stadium_collage.name_stadium_collage,stadium_collage.amount_people,
	stadium.name_stadium,stadium.address,stadium.category, stadium.stadium_id ,
	stadium_details.price , stadium_details.start_time_detail , stadium_details.end_time_detail, orders_status.*
	FROM public.orders 
	INNER JOIN users ON users.user_id = orders.user_id 
	INNER JOIN stadium_details ON stadium_details.stadium_detail_id = orders.stadium_detail_id
	INNER JOIN stadium_collage  ON stadium_collage.stadium_collage_id = stadium_details.stadium_collage_id  
	INNER JOIN stadium ON stadium.stadium_id = stadium_collage.stadium_id 
	INNER JOIN orders_status  ON orders_status.order_id = orders.order_id
	WHERE users.user_id = $1 ORDER BY orders.time DESC;
	`

	err := o.sql.Db.SelectContext(context, &orders, sqlStatement, userId)
	if err != nil {
		log.Error(err.Error())
		return orders, message.SomeWentWrong
	}

	for i := 0; i < len(orders); i++ {
		orders[i].Stadium.StadiumId = orders[i].StadiumCollage.StadiumId
		orders[i].StadiumCollage.StadiumCollageId = orders[i].StadiumDetails.StadiumCollageId
	}

	return orders, nil
}

func (o OrderRepoImpl) OrderDetail(context context.Context, orderId string) (interface{}, error) {
	type listOrders struct {
		model.Order
		model.OrderStatus    `json:"order_status"`
		model.Stadium        `json:"stadium"`
		model.StadiumCollage `json:"stadium_collage"`
		model.StadiumDetails `json:"stadium_details"`
		model.User           `json:"user"`
	}
	var orders = listOrders{}
	sqlStatement := `
	SELECT orders.*,
	users.user_id,users.display_name,users.avatar, users.phone,
	stadium_collage.name_stadium_collage,stadium_collage.amount_people,
	stadium.name_stadium, stadium.image , stadium.latitude, stadium.longitude ,stadium.address,stadium.category, stadium.stadium_id, 
	stadium_details.price , stadium_details.start_time_detail , stadium_details.end_time_detail, orders_status.*
	FROM public.orders 
	INNER JOIN users ON users.user_id = orders.user_id 
	INNER JOIN stadium_details ON stadium_details.stadium_detail_id = orders.stadium_detail_id
	INNER JOIN stadium_collage  ON stadium_collage.stadium_collage_id = stadium_details.stadium_collage_id  
	INNER JOIN stadium ON stadium.stadium_id = stadium_collage.stadium_id 
	INNER JOIN orders_status  ON orders_status.order_id = orders.order_id
	WHERE orders.order_id = $1;
	`

	err := o.sql.Db.GetContext(context, &orders, sqlStatement, orderId)
	if err != nil {
		log.Error(err.Error())
		return orders, message.SomeWentWrong
	}
	return orders, nil
}
