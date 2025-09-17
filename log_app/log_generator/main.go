package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"
)

func main() {
	file, error := os.Create("/usr/src/app/files/hashes.txt")

	if error != nil {
		fmt.Println("Error creating file")
	}

	defer file.Close()

	for {
		data := []byte(time.Now().String())

		hash := sha256.Sum256(data)
		_, error := file.WriteString(fmt.Sprintf("%x\n", hash))
		if error != nil {
			fmt.Println("Error", error)
		}

		fmt.Printf("Wrote hash: %x\n", hash)

		time.Sleep(5 * time.Second)
	}

}
