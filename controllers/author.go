package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/orderbynull/book-store/core"
	"github.com/orderbynull/book-store/repository"
	"net/http"
)

//GetAuthors возвращает всех авторов
//GET /authors
func GetAuthors(app *core.App, w http.ResponseWriter, r *http.Request) {

	//Репозиторий, который отвечает за получение сущностей
	repo := repository.NewAuthorRepository(app.DB)

	//Получаем книги по ID библиотеки
	authors, err := repo.All()
	if err != nil && err != sql.ErrNoRows {
		app.Fatal(err)
	}

	//Возвращаем результат на фронт
	if err := json.NewEncoder(w).Encode(authors); err != nil {
		app.Fatal(err)
	}
}
