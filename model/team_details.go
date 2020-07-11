package model

import "time"

type TeamDetails struct {
	TeamDetailsId string    `json:"teamDetailId,omitempty" db:"team_details_id,omitempty"`
	TeamId        string    `json:"teamId,omitempty" db:"teams_id,omitempty,prefix=team_details."`
	UserId        string    `json:"userId,omitempty" db:"user_id,omitempty"`
	Role          string    `json:"role,omitempty" db:"role_team,omitempty"`
	Accept        string    `json:"accept,omitempty" db:"accept,omitempty"`
	CreatedAt     time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at, omitempty"`
	//User        `json:"user,omitempty"`
}
