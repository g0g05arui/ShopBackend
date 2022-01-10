package models

type Review struct {
	Id          string   `json:"id"`
	CreatedAt   string   `json:"createdAt"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Rating      float64  `json:"rating"`
}
