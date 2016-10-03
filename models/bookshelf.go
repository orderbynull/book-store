package models

//Bookshelf представляет собой библиотеку книг
type Bookshelf struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Bookshelfs []Bookshelf
