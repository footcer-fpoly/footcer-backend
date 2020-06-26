package model

type ReqSignIn struct {
	Phone    string `json:"phone,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
