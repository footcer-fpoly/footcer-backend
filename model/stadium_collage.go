package model

import "time"

type StadiumCollage struct {
	StadiumCollageId   string    `json:"stadiumCollageId" db:"stadium_collage_id,omitempty"`
	NameStadiumCollage string    `json:"stadiumCollageName" db:"name_stadium_collage,omitempty"`
	AmountPeople       string    `json:"amountPeople" db:"amount_people"`
	StadiumId          string    `json:"stadiumId" db:"stadium_id,omitempty"`
	CreatedAt          time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"-" db:"updated_at,omitempty"`
}
