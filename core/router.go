package core

//Route представляет роут приложения
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler Action
}

type Routes []Route
