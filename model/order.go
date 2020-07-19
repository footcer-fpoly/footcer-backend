package model

import "time"

type Order struct {
	OrderId          string    `json:"orderId,omitempty" db:"order_id,omitempty"`
	TimeSlot         string    `json:"timeSlot,omitempty" db:"time_slot,omitempty"`
	Time             string    `json:"time,omitempty" db:"time,omitempty"`
	Price             string    `json:"price,omitempty" db:"price,omitempty"`
	StadiumId        string    `json:"stadiumId,omitempty" db:"stadium_id,omitempty"`
	StadiumCollageId string    `json:"stadiumCollageId,omitempty" db:"stadium_collage_id,omitempty"`
	UserId           string    `json:"userId,omitempty" db:"user_id,omitempty"`
	Finish           string    `json:"finish,omitempty" db:"finish,omitempty"`
	Accept           string    `json:"accept,omitempty" db:"accept,omitempty"`
	CreatedAt        time.Time `json:"createdAt" db:"order_created_at,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt" db:"order_updated_at, omitempty"`
}
