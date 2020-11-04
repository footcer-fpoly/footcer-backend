package model

import "time"

type StadiumDetails struct {
	StadiumDetailsId string    `json:"stadiumDetailsId" db:"stadium_detail_id,omitempty"`
	StadiumCollageId string    `json:"stadiumCollageId" db:"stadium_collage_id"`
	StartTimeDetails string    `json:"startTimeDetail" db:"start_time_detail,omitempty"`
	EndTimeDetails   string    `json:"endTimeDetail," db:"end_time_detail,omitempty" `
	Price            int       `json:"price," db:"price,omitempty"`
	Description      string    `json:"description," db:"description,omitempty"`
	HasOrder         bool      `json:"hasOrder," db:"has_order,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}
