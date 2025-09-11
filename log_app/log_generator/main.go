package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	randomString := generateRandomString(rand.Intn(10) + 1)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(randomString))
	})
	http.ListenAndServe(":3000", r)

	fmt.Println(time.Now().Format(time.RFC3339) + ":" + randomString)

}
