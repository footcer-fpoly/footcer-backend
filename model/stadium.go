package model

import "time"

type Stadium struct {
	StadiumId   string    `json:"stadiumId,omitempty" db:"stadium_id,omitempty"`
	StadiumName string    `json:"stadiumName,omitempty" db:"name_stadium,omitempty"`
	Address     string    `json:"address,omitempty" db:"address"`
	Description string    `json:"description,omitempty" db:"description,omitempty"`
	Image       string    `json:"image,omitempty" db:"image,omitempty"`
	RateCount   float64   `json:"rateCount,omitempty" db:"_,omitempty"`
	StartTime   string    `json:"start_time,omitempty" db:"start_time,omitempty"`
	EndTime     string    `json:"end_time,omitempty" db:"end_time,omitempty"`
	Category    string    `json:"category,omitempty" db:"category,omitempty"`
	Latitude    float64   `json:"latitude,omitempty" db:"latitude,omitempty"`
	Longitude   float64   `json:"longitude,omitempty" db:"longitude,omitempty"`
	Ward        string    `json:"ward,omitempty" db:"ward,omitempty"`
	District    string    `json:"district,omitempty" db:"district,omitempty"`
	City        string    `json:"city,omitempty" db:"city,omitempty"`
	TimePeak    string    `json:"timePeak,omitempty" db:"time_peak,omitempty"`
	TimeOrder   string    `json:"timeOrder,omitempty" db:"_,omitempty" `
	UserId      string    `json:"userId,omitempty" db:"user_id,omitempty"`
	CreatedAt   time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at,omitempty"`
}
