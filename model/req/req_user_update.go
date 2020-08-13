package req

type ReqUpdateUser struct {
	DisplayName string `json:"fullName,omitempty" validate:"required"`
	Phone       string `json:"email,omitempty" validate:"required"`
	Avatar      string `json:"email,omitempty" validate:"required"`
	Birthday    string `json:"email,omitempty" validate:"required"`
}
