package main

import (
	"fmt"
	"net/http"
	"os"
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

		os.WriteFile("/usr/src/app/images/counter.txt", []byte(strconv.Itoa(counter)), 0644)
	})
	fmt.Println("Server started in port 3001")

	http.ListenAndServe(":3001", r)
}
