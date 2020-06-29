package model

import "time"

type Stadium struct {
	StadiumId   string    `json:"stadiumId" db:"stadium_id,omitempty"`
	StadiumName string    `json:"stadiumName" db:"name_stadium,omitempty"  valid:"required"`
	Address     string    `json:"address" db:"address,omitempty"  valid:"required"`
	Description string    `json:"description" db:"description,omitempty"  valid:"required"`
	Image       string    `json:"image" db:"image,omitempty"`
	PriceNormal int       `json:"priceNormal" db:"price_normal,omitempty"  valid:"required"`
	PricePeak   int       `json:"pricePeak" db:"price_peak,omitempty"  valid:"required"`
	StartTime   string    `json:"start_time" db:"start_time,omitempty"`
	EndTime     string    `json:"end_time" db:"end_time,omitempty"`
	Category    string    `json:"category" db:"category,omitempty"`
	Latitude    float64   `json:"latitude" db:"latitude,omitempty"`
	Longitude   float64   `json:"longitude" db:"longitude,omitempty"`
	Ward        string    `json:"ward" db:"ward,omitempty"`
	District    string    `json:"district" db:"district,omitempty"`
	City        string    `json:"city" db:"city,omitempty"`
	UserId      string    `json:"user_id" db:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at" db:"created_at, omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at, omitempty"`
}
