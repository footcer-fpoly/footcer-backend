package model

type GameTemp struct {
	GameTempId string `json:"gameTempId,omitempty" db:"game_temp_id,omitempty"`
	GameId     string `json:"gameId,omitempty" db:"game_id,omitempty"`
	TeamId     string `json:"teamId,omitempty" db:"team_id,omitempty"`

	UserNotifyId string `json:"userNotifyId,omitempty" db:"-,omitempty"`
	NameHost     string `json:"nameHost,omitempty" db:"-,omitempty"`
	NameInvite   string `json:"nameInvite,omitempty" db:"-,omitempty"`
}
