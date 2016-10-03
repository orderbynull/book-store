package models

//Book представляет собой книгу библиотеки
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Books []Book