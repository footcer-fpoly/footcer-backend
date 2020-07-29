package model

type Service struct {
	ServiceId string `json:"serviceId,omitempty" db:"service_id,omitempty"`
	Name      string `json:"name,omitempty" db:"name_service,omitempty"`
	Price     string `json:"price,omitempty" db:"price_service,omitempty"`
	Image     string `json:"image,omitempty" db:"image,omitempty"`
	StadiumId string `json:"stadiumId,omitempty" db:"stadium_id,omitempty"`
}
