package controllers

import (
	"net/http"

	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/orderbynull/book-store/core"
	"github.com/orderbynull/book-store/repository"
	"strconv"
)

//GetBooks возвращает книги по ID библиотеки
//GET /books/{id}
func GetBooks(app *core.App, w http.ResponseWriter, r *http.Request) {

	//Репозиторий, который отвечает за получение сущностей
	repo := repository.NewBookRepository(app.DB)

	//ID библиотеки из GET-запроса
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//Получаем книги по ID библиотеки
	books, err := repo.ByBookshelf(id)
	if err != nil && err != sql.ErrNoRows {
		app.Fatal(err)
	}

	//Возвращаем результат на фронт
	if err := json.NewEncoder(w).Encode(books); err != nil {
		app.Fatal(err)
	}
}

//PostBook создает новую книгу
//POST /books
func PostBook(app *core.App, w http.ResponseWriter, r *http.Request) {

	//Репозиторий, который отвечает за создание сущностей
	repo := repository.NewBookRepository(app.DB)

	//Параметры из POST-запроса
	title := r.FormValue("title")
	authorID, _ := strconv.Atoi(r.FormValue("author_id"))
	bookShelfID, _ := strconv.Atoi(r.FormValue("bookshelf_id"))

	//Создание книги
	if _, err := repo.New(title, authorID, bookShelfID); err != nil {
		app.Fatal(err)
	}
}

//DeleteBook удаляет книгу
//DELETE /books/{id}
func DeleteBook(app *core.App, w http.ResponseWriter, r *http.Request) {

	//Репозиторий, который отвечает за удалений сущностей
	repo := repository.NewBookRepository(app.DB)

	//ID книги из DELETE-запроса
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//Создание книги
	if err := repo.Delete(id); err != nil {
		app.Fatal(err)
	}
}
