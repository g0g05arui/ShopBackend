package models

type Product struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Thumbnail   string     `json:"thumbnail"`
	Photos      []string   `json:"photos"`
	Categories  [][]string `json:"categories"`
	Rating      float64    `json:"rating"`
}
