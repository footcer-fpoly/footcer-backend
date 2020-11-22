package model

type Statistics struct {
	TotalDetails      int `json:"totalDetails" db:"total_details,omitempty"`
	TotalDetailsOrder int `json:"totalDetailsOrder" db:"total_details_order,omitempty"`
	TotalCustomer     int `json:"totalCustomer" db:"total_customer,omitempty"`
	TotalPrice        int `json:"totalPrice" db:"total_price,omitempty"`
}
