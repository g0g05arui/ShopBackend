package models

type HttpError struct {
	Error string `json:"error"`
	Code  int32  `json:"code"`
}
