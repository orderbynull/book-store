package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/orderbynull/book-store/core"
	"github.com/orderbynull/book-store/repository"
)

//GetBookshelfs возвращает все имеющиеся библиотеки
//GET /bookshelfs
func GetBookshelfs(app *core.App, w http.ResponseWriter, r *http.Request) {

	//Репозиторий, который отвечает за получение сущностей
	repo := repository.NewBookshelfRepository(app.DB)

	//Получаем все библиотеки
	shelfs, err := repo.All()
	if err != nil {
		app.Fatal(err)
	}

	//Возвращаем результат на фронт
	if err := json.NewEncoder(w).Encode(shelfs); err != nil {
		app.Fatal(err)
	}
}

//GetBookshelfs возвращает библиотеку по ее ID
//GET /bookshelfs/{id}
func GetBookshelf(app *core.App, w http.ResponseWriter, r *http.Request) {

	//Репозиторий, который отвечает за получение сущностей
	repo := repository.NewBookshelfRepository(app.DB)

	//ID библиотеки из GET-запроса
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//Получаем библиотеку
	shelf, err := repo.One(id)
	if err != nil && err != sql.ErrNoRows {
		app.Fatal(err)
	}

	//Возвращаем результат на фронт
	if err := json.NewEncoder(w).Encode(shelf); err != nil {
		app.Fatal(err)
	}
}

//PostBookshelf создает новую библиотеку
//POST /bookshelfs
func PostBookshelf(app *core.App, w http.ResponseWriter, r *http.Request) {

	//Репозиторий, который отвечает за создание сущностей
	repo := repository.NewBookshelfRepository(app.DB)

	//Title новой библиотеки из POST-запроса
	title := r.FormValue("title")

	//Создаем библиотеку
	_, err := repo.New(title)
	if err != nil && err != sql.ErrNoRows {
		app.Fatal(err)
	}
}
