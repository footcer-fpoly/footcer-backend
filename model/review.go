package model

import (
	"time"
)

type Review struct {
	ReviewId  string  `json:"reviewId,omitempty" db:"review_id,omitempty"`
	Content   string  `json:"content,omitempty" db:"content,omitempty"`
	Rate      float64 `json:"rate,omitempty" db:"rate,omitempty"`
	UserId    string  `json:"userId,omitempty" db:"user_id,omitempty"`
	StadiumId string  `json:"stadiumId,omitempty" db:"stadium_id,omitempty"`
	User      `json:"user"`
	CreatedAt time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at, omitempty"`
}
