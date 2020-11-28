package model

type StatisticsAdmin struct {
	TotalUser    int `json:"totalUser" db:"total_user,omitempty"`
	TotalStadium int `json:"totalStadium" db:"total_stadium,omitempty"`
	TotalOrder   int `json:"totalOrder" db:"total_order,omitempty"`
	TotalTeam    int `json:"totalTeam" db:"total_team,omitempty"`
}
