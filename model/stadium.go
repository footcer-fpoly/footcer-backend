package model

import "time"

type Stadium struct {
	StadiumId   string  `json:"stadiumId" db:"stadium_id,omitempty"`
	StadiumName string  `json:"stadiumName" db:"name_stadium,omitempty"`
	Address     string  `json:"address" db:"address"`
	Description string  `json:"description" db:"description,omitempty"`
	Image       string  `json:"image" db:"image,omitempty"`
	PriceNormal int     `json:"priceNormal" db:"price_normal,omitempty"`
	PricePeak   int     `json:"pricePeak" db:"price_peak,omitempty" `
	StartTime   string  `json:"start_time" db:"start_time,omitempty"`
	EndTime     string  `json:"end_time" db:"end_time,omitempty"`
	Category    string  `json:"category" db:"category,omitempty"`
	Latitude    float64 `json:"latitude" db:"latitude,omitempty"`
	Longitude   float64 `json:"longitude" db:"longitude,omitempty"`
	Ward        string  `json:"ward" db:"ward,omitempty"`
	District    string  `json:"district" db:"district,omitempty"`
	City        string  `json:"city" db:"city,omitempty"`
	UserId      string  `json:"user_id" db:"user_id,omitempty"`
	User        struct {
		DisplayName string `json:"display_name" db:"display_name,prefix=users.display_name"`
		Avatar      string `json:"avatar" db:"avatar,omitempty"`
		Phone       string `json:"phone" db:"phone,omitempty"`
	} `json:"user"`
	CreatedAt time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at,omitempty"`
}
