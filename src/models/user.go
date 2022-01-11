package models

import "time"

type User struct {
	Id               string    `json:"id"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	RegistrationDate time.Time `json:"registrationDate"`
	Address          Address   `json:"address"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	Role             string    `json:"role"`
	PhoneNo          string    `json:"phoneNo"`
	FirstName        string    `json:"firstName"`
	Surname          string    `json:"surname"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
