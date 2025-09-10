package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func generateRandomString(length int) string {
	var randomString string
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < length; i++ {
		randomString += string(letters[rand.Intn(len(letters))])
	}
	return randomString
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(generateRandomString(rand.Intn(10) + 1)))
	})
	fmt.Println("Server started in port 3000")

	http.ListenAndServe(":3000", r)
}
