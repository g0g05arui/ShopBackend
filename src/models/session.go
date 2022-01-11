package models

import "time"

type Session struct {
	Id        string    `json:"id"`
	Expires   time.Time `json:"expires"`
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
}
