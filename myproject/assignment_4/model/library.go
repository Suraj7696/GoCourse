package model

type Book struct {
	Title            string `json:"title"`
	Author           string `json:"author"`
	Publication_year int    `json:"year"`
}
