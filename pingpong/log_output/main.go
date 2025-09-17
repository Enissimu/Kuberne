package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.ReadFile("/usr/src/app/files/hashes.txt")
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(file)

	})

	http.ListenAndServe(":3000", r)

}
