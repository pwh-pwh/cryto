package main

import (
	"fmt"
	"log"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
	// ?
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	return fmt.Sprintf("%x", bytes), err
}

// don't touch below this line

func main() {
	rand.Seed(0)
	for i := 16; i < 33; i++ {
		key, err := generateRandomKey(i)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v-byte key: %v\n", i, key)
	}
}
