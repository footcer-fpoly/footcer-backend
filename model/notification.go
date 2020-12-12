package model

import "time"

type Notification struct {
	NotifyID  string    `json:"notifyId,omitempty" db:"notify_id,omitempty"`
	Key       string    `json:"key,omitempty" db:"key,omitempty"`
	Title     string    `json:"title,omitempty" db:"title,omitempty"`
	Content   string    `json:"content,omitempty" db:"content,omitempty"`
	Icon      string    `json:"icon,omitempty" db:"icon,omitempty"`
	GeneralID string    `json:"generalId,omitempty" db:"general_id,omitempty"`
	UserId    string    `json:"userId,omitempty" db:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at" db:"created_at_notify"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at_notify, omitempty"`
}
