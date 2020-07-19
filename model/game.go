package model

import "time"

type Game struct {
	GameId      string    `json:"gameId,omitempty" db:"game_id,omitempty"`
	Date        string `json:"date,omitempty" db:"date,omitempty"`
	Hour        string `json:"hour,omitempty" db:"hour,omitempty"`
	Score        string `json:"score,omitempty" db:"score,omitempty"`
	Type        string    `json:"type,omitempty" db:"type,omitempty"`
	Description string    `json:"description,omitempty" db:"description,omitempty"`
	Finish      string    `json:"finish,omitempty" db:"finish,omitempty"`
	StadiumId   string    `json:"stadiumId,omitempty" db:"stadium_id,omitempty"`
	TeamIdHost  string    `json:"teamIdHost,omitempty" db:"team_id_host,omitempty"`
	TeamIdGuest string    `json:"teamIdGuest,omitempty" db:"team_id_guest,omitempty"`
	CreatedAt   time.Time `json:"created_at" db:"game_created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" db:"game_updated_at, omitempty"`
}
