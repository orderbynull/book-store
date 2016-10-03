package main

import (
	ctrl "github.com/orderbynull/book-store/controllers"
	"github.com/orderbynull/book-store/core"
)

//Свои роуты нужно прописывать здесь
func makeRoutes() *core.Routes {
	return &core.Routes{
		core.Route{Name: "GETBookshelfs", Method: "GET", Pattern: "/bookshelfs", Handler: ctrl.GetBookshelfs},
		core.Route{Name: "GETBookshelf", Method: "GET", Pattern: "/bookshelfs/{id}", Handler: ctrl.GetBookshelf},
		core.Route{Name: "POSTBookshelf", Method: "POST", Pattern: "/bookshelfs", Handler: ctrl.PostBookshelf},

		core.Route{Name: "GETBooks", Method: "GET", Pattern: "/books/{id}", Handler: ctrl.GetBooks},
		core.Route{Name: "POSTBooks", Method: "POST", Pattern: "/books", Handler: ctrl.PostBook},
		core.Route{Name: "DELETEBooks", Method: "DELETE", Pattern: "/books/{id}", Handler: ctrl.DeleteBook},

		core.Route{Name: "GETAuthors", Method: "GET", Pattern: "/authors", Handler: ctrl.GetAuthors},
	}
}
