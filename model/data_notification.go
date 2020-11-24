package model

type DataNotification struct {
	Type string           `json:"type,omitempty"`
	Body BodyNotification `json:"body,omitempty"`
}

type BodyNotification struct {
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	GeneralId string `json:"generalId,omitempty"`
}
