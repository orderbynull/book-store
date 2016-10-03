package main

import (
	"github.com/orderbynull/book-store/core"
	"os"
)

func main() {
	core.
		NewApp("postgres", "aleksandr", "postgres").
		SetFatalWriter(os.Stdout).
		AddRoutes(makeRoutes()).
		Run(":8081")
}
