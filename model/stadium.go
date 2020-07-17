package model

import "time"

type Stadium struct {
	StadiumId   string    `json:"stadiumId" db:"stadium_id,omitempty"`
	StadiumName string    `json:"stadiumName" db:"name_stadium,omitempty"`
	Address     string    `json:"address" db:"address"`
	Description string    `json:"description" db:"description,omitempty"`
	Image       string    `json:"image" db:"image,omitempty"`
	RateCount   float64   `json:"rateCount" db:"_,omitempty"`
	StartTime   string    `json:"start_time" db:"start_time,omitempty"`
	EndTime     string    `json:"end_time" db:"end_time,omitempty"`
	Category    string    `json:"category" db:"category,omitempty"`
	Latitude    float64   `json:"latitude" db:"latitude,omitempty"`
	Longitude   float64   `json:"longitude" db:"longitude,omitempty"`
	Ward        string    `json:"ward" db:"ward,omitempty"`
	District    string    `json:"district" db:"district,omitempty"`
	City        string    `json:"city" db:"city,omitempty"`
	TimePeak    string    `json:"timePeak" db:"time_peak,omitempty"`
	TimeOrder   string    `json:"timeOrder" db:"_,omitempty" `
	UserId      string    `json:"userId,omitempty" db:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at,omitempty"`
}
