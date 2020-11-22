package model

import "time"

type Order struct {
	OrderId          string    `json:"orderId,omitempty" db:"order_id,omitempty"`
	Time             string `json:"time,omitempty" db:"time,omitempty"`
	Price            int       `json:"price,omitempty" db:"price,omitempty"`
	Description      string    `json:"description,omitempty" db:"description,omitempty"`
	StadiumDetailsId string    `json:"stadiumDetailsId,omitempty" db:"stadium_detail_id,omitempty"`
	UserId           string    `json:"userId,omitempty" db:"user_id,omitempty"`
	Finish           bool      `json:"finish,omitempty" db:"finish,omitempty"`
	CreatedAt        time.Time `json:"createdAt" db:"order_created_at,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt" db:"order_updated_at, omitempty"`
}
