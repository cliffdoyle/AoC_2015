package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Check if a string is nice based on the given criteria
func isNiceString(s string) bool {
	// Check for at least three vowels
	vowelCount := 0
	vowels := "aeiou"
	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			vowelCount++
		}
	}
	if vowelCount < 3 {
		return false
	}

	// Check for at least one double letter (e.g. "xx")
	hasDouble := false
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			hasDouble = true
			break
		}
	}
	if !hasDouble {
		return false
	}

	// Check for forbidden substrings
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, substr := range forbidden {
		if strings.Contains(s, substr) {
			return false
		}
	}

	// If all checks passed, the string is nice
	return true
}

func main() {
	// Open the input file
	file, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize a counter for nice strings
	niceCount := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line is nice
		if isNiceString(line) {
			niceCount++
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Output the number of nice strings
	fmt.Println("Number of nice strings:", niceCount)
}
