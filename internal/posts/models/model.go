package models

type Video struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
}

type ID struct {
	ID int `json:"id"`
}
