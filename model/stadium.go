package model

import "time"

type Stadium struct {
	StadiumId   string    `json:"stadiumId,omitempty" db:"stadium_id,omitempty"`
	StadiumName string    `json:"stadiumName,omitempty" db:"name_stadium,omitempty"`
	Address     string    `json:"address,omitempty" db:"address"`
	Description string    `json:"description,omitempty" db:"description,omitempty"`
	Image       string    `json:"image,omitempty" db:"image,omitempty"`
	RateCount   float64   `json:"rateCount,omitempty" db:"_,omitempty"`
	Category    string    `json:"category,omitempty" db:"category,omitempty"`
	Latitude    float64   `json:"latitude,omitempty" db:"latitude,omitempty"`
	Longitude   float64   `json:"longitude,omitempty" db:"longitude,omitempty"`
	Ward        string    `json:"ward,omitempty" db:"ward,omitempty"`
	District    string    `json:"district,omitempty" db:"district,omitempty"`
	City        string    `json:"city,omitempty" db:"city,omitempty"`
	UserId      string    `json:"userId,omitempty" db:"user_id,omitempty"`
	Verify      string    `json:"verify" db:"verify,omitempty"`
	CreatedAt   time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at,omitempty"`
	Distance    int       `json:"distance" db:"_,omitempty"`
	Timer       int       `json:"timer" db:"_,omitempty"`
	Folder      string    `json:"folder,omitempty" validate:"required"`

}
