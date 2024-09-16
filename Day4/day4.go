package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func findAdventCoin(secretKey string) int {
	number := 1 // Start from 1 and increment

	for {
		// Combine the secret key with the current number and hash it
		data := secretKey + strconv.Itoa(number)
		hash := md5.Sum([]byte(data))
		hashString := hex.EncodeToString(hash[:])

		// Check if the hash starts with "00000"
		if hashString[:5] == "00000" {
			return number // Return the first valid number
		}

		number++ // Increment and try the next number
	}
}

func main() {
	// Secret key given in the problem
	secretKey := "bgvyzdsv"

	// Find the lowest number that produces a hash starting with "00000"
	result := findAdventCoin(secretKey)
	fmt.Println("Lowest number:", result)
}
