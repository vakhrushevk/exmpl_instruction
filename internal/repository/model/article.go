package model

type Article struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
	Prev     int    `json: "prev"`
	Next     int    `json:"next"`
}
