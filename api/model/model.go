package model

type Item struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}
