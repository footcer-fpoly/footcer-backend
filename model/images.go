package model

import (
	"time"
)

type Images struct {
	ImageId   string    `json:"imageId,omitempty" db:"image_id,omitempty"`
	GeneralId string    `json:"generalId,omitempty" db:"general_id,omitempty"`
	Url       string    `json:"url,omitempty" db:"url,omitempty"`
	CreatedAt time.Time `json:"created_at_img" db:"created_at_img"`
	UpdatedAt time.Time `json:"updated_at_img" db:"updated_at_img, omitempty"`
}
