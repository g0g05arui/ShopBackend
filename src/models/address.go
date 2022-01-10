package models

type Address struct {
	Country  string `json:"country"`
	County   string `json:"county"`
	City     string `json:"city"`
	Street   string `json:"street"`
	StreetNo int    `json:"streetNo"`
	FlatNo   int    `json:"flatNo"`
	Details  string `json:"details"`
	ZipCode  string `json:"zipCode"`
}
