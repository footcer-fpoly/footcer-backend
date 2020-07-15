package model

import "time"

type User struct {
	UserId      string    `json:"userId,omitempty" db:"user_id,omitempty"`
	Phone       string    `json:"phone,omitempty" db:"phone,omitempty" valid:"required"`
	Email       string    `json:"email,omitempty" db:"email,omitempty" valid:"required"`
	Password    string    `json:"_,omitempty" db:"password,omitempty"`
	Avatar      string    `json:"avatar,omitempty" db:"avatar,omitempty" valid:"required"`
	DisplayName string    `json:"displayName,omitempty" db:"display_name,omitempty" valid:"required"`
	Role        int8      `json:"role,omitempty" db:"role,omitempty" valid:"required"`
	Birthday    string    `json:"birthday,omitempty" db:"birthday,omitempty"`
	Position    string    `json:"position,omitempty" db:"position,omitempty"`
	Level       string    `json:"level,omitempty" db:"level,omitempty"`
	Verify      string    `json:"verify,omitempty" db:"verify,omitempty"`
	CreatedAt   time.Time `json:"-" db:"created_at, omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at, omitempty"`
	Token       string    `json:"token,omitempty"`
}
