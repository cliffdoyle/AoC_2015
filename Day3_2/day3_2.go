package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

func countHousesWithRoboSanta(directions string) int {
	// Map to track visited houses
	visitedHouses := make(map[Point]bool)

	// Starting positions for both Santa and Robo-Santa
	santa := Point{0, 0}
	roboSanta := Point{0, 0}

	// Mark the starting house as visited twice (Santa and Robo-Santa both visit the start)
	visitedHouses[santa] = true

	// Iterate over the directions, alternating between Santa and Robo-Santa
	for i, dir := range directions {
		if i%2 == 0 {
			// Santa's turn
			santa = move(santa, dir)
			visitedHouses[santa] = true
		} else {
			// Robo-Santa's turn
			roboSanta = move(roboSanta, dir)
			visitedHouses[roboSanta] = true
		}
	}

	// The number of unique visited houses is the size of the map
	return len(visitedHouses)
}

// Helper function to move Santa or Robo-Santa based on the direction
func move(pos Point, dir rune) Point {
	switch dir {
	case '^': // Move north
		pos.y++
	case 'v': // Move south
		pos.y--
	case '>': // Move east
		pos.x++
	case '<': // Move west
		pos.x--
	}
	return pos
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
	fmt.Println(countHousesWithRoboSanta(line))         // Output: 3
	fmt.Println(countHousesWithRoboSanta("^>v<"))       // Output: 3
	fmt.Println(countHousesWithRoboSanta("^v^v^v^v^v")) // Output: 11
}
