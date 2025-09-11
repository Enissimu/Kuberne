package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var counter = 0

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong " + strconv.Itoa(counter)))
		counter++

	})
	fmt.Println("Server started in port 3000")

	http.ListenAndServe(":3000", r)
}
