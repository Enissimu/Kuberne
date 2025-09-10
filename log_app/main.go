package main

import (
	"fmt"
	"math/rand"
	"time"
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
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	randomString := generateRandomString(rand.Intn(10) + 1)
	fmt.Println(time.Now().Format(time.RFC3339) + ":" + randomString)
	for range ticker.C {
		fmt.Println(time.Now().Format(time.RFC3339) + ":" + randomString)
	}
}
