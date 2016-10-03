package core

import (
	"log"
	"net/http"

	"database/sql"

	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io"
)

type Action func(app *App, w http.ResponseWriter, r *http.Request)

//App представляет приложение
type App struct {
	Router      *mux.Router
	DB          *sql.DB
	FatalWriter io.Writer
}

//NewApp возвращает адрес инстанса нового приложения
func NewApp(driver, user, dbname string) *App {
	db, err := sql.Open(driver, fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname))
	if err != nil {
		log.Fatal(err)
	}
	return &App{
		Router: mux.NewRouter().StrictSlash(true),
		DB:     db}
}

//SetFatalWriter устанавливает источник для фатальных ошибок
func (app *App) SetFatalWriter(w io.Writer) *App {
	app.FatalWriter = w

	return app
}

//Fatal пишет фатальную ошибку
func (app *App) Fatal(err error) {
	fmt.Fprintf(app.FatalWriter, err)
}

//AddRoutes добавляет роуты приложения
func (app *App) AddRoutes(routes *Routes) *App {

	for _, route := range *routes {
		app.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(app.WithApp(route.Handler))
	}

	app.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return app
}

//Run запускает приложение
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

//WithApp это middleware
func (app *App) WithApp(fn Action) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		fn(app, w, r)
	}
}
