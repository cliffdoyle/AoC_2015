package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Struct to represent a 2D coordinate (x, y)
type Point struct {
	x, y int
}

func countHouses(directions string) int {
	// Use a map to track visited houses
	visitedHouses := make(map[Point]bool)

	// Start at the origin (0, 0)
	x, y := 0, 0

	// Mark the starting house as visited
	visitedHouses[Point{x, y}] = true

	// Iterate over each direction in the input string
	for _, dir := range directions {
		switch dir {
		case '^': // Move north
			y++
		case 'v': // Move south
			y--
		case '>': // Move east
			x++
		case '<': // Move west
			x--
		}
		// Mark the new house as visited
		visitedHouses[Point{x, y}] = true
	}

	// The number of unique visited houses is the size of the map
	return len(visitedHouses)
}

func main() {
	file, _ := os.Open("testfile.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Test cases
	fmt.Println(countHouses(line))         // Output: 2
	fmt.Println(countHouses("^>v<"))       // Output: 4
	fmt.Println(countHouses("^v^v^v^v^v")) // Output: 2
}
