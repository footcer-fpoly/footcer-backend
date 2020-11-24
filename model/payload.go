package model

type PayLoad struct {
	RegistrationIds []string         `json:"registration_ids"`
	Data            DataNotification `json:"data"`
}
