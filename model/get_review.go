package model

import "time"

type GetReview struct {
	ReviewId  string    `json:"reviewId,omitempty" db:"review_id,omitempty"`
	Content   string    `json:"content,omitempty" db:"content,omitempty"`
	Rate      float64   `json:"rate,omitempty" db:"rate,omitempty"`
	UserId    string    `json:"userId,omitempty" db:"user_id,omitempty"`
	StadiumId string    `json:"stadiumId,omitempty" db:"stadium_id,omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at, omitempty"`
	//User      `json:"user,omitempty"`
}