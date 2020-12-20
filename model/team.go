package model

import "time"

type Team struct {
	TeamId      string    `json:"teamId,omitempty" db:"team_id,omitempty"`
	Name        string    `json:"name,omitempty" db:"name,omitempty"`
	Level       string    `json:"level,omitempty" db:"level,omitempty"`
	Place       string    `json:"place,omitempty" db:"place,omitempty"`
	Description string    `json:"description,omitempty" db:"description,omitempty"`
	Avatar      string    `json:"avatar,omitempty" db:"avatar,omitempty"`
	Background  string    `json:"background,omitempty" db:"background,omitempty"`
	LeaderId    string    `json:"leaderId,omitempty" db:"leader_id,omitempty"`
	MemberList  string    `json:"memberList,omitempty" db:"omitempty"`
	CreatedAt   time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at, omitempty"`
	//User        `json:"user,omitempty"`
	Folder      string    `json:"folder,omitempty" validate:"required"`

	NameUser string `json:"nameUser,omitempty"`


}
