package main

import (
	"github.com/calebtracey/go-htmx/internal/routes"
)

func main() {

	handler := routes.Handler{}

	e := handler.Initialize()

	e.Logger.Fatal(e.Start(localhost))
}

const localhost = "localhost:42069"
const htmxSource = "https://unpkg.com/htmx.org@1.9.6"
