package req

type TokenNotify struct {
	TokenNotify string `json:"tokenNotify,omitempty" validate:"required" db:"token_notify,omitempty"`
	UserId      string `json:"email,omitempty" db:"user_id,omitempty"`
}
