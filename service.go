package main

import (
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	r := buildRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	n := negroni.New()
	n.UseHandler(r)
	n.Run(":" + port)
}
