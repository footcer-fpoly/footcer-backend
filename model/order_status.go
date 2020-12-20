package model

import "time"

type OrderStatus struct {
	OrderStatusId string    `json:"-" db:"order_status_id,omitempty"`
	OrderId       string    `json:"orderId" db:"orders_status.order_id,omitempty"`
	Status        string    `json:"status,omitempty" db:"status,omitempty"`
	Reason        string    `json:"reason" db:"reason,omitempty"`
	IsUser        bool      `json:"isUser" db:"is_user,omitempty"`
	CreatedAt     time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"-" db:"updated_at, omitempty"`

	UserId      string `json:"userNotifyId" db:"-,omitempty"`
	StadiumName string `json:"stadiumName" db:"-,omitempty"`
}
