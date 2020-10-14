package model

import "time"

type StadiumCollage struct {
	StadiumCollageId   string    `json:"stadiumCollageId" db:"stadium_collage_id,omitempty"`
	NameStadiumCollage string    `json:"stadiumCollageName" db:"name_stadium_collage,omitempty"`
	AmountPeople       string    `json:"amountPeople" db:"amount_people"`
	StartTime          string    `json:"startTime" db:"start_time"`
	EndTime            string    `json:"endTime" db:"end_time"`
	PlayTime           string    `json:"playTime" db:"play_time"`
	StadiumId          string    `json:"stadiumId," db:"stadium_id,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}
