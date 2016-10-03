package models

//Author представляет собой автора книги
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Authors []Author
