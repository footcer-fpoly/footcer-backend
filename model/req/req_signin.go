package req

type ReqSignIn struct {
	Phone    string `json:"phone,omitempty" validate:"required,phone"`
	Password string `json:"password,omitempty" validate:"required"`
}
