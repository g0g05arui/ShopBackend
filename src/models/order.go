package models

import "time"

type Order struct {
	Id              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Status          string    `json:"status"`
	ShippingAddress Address   `json:"shippingAddress"`
	BillingAddress  Address   `json:"billingAddress"`
	UserId          string    `json:"userId"`
	Total           float64   `json:"total"`
	Currency        string    `json:"currency"`
	Products        []Product `json:"products"`
}
