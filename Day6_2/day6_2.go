package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Constants for grid size
const gridSize = 1000

// Parse an instruction and update the grid accordingly
func applyInstruction(grid [][]int, instruction string) {
	// Regular expression to parse the instruction
	re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	matches := re.FindStringSubmatch(instruction)
	if len(matches) != 6 {
		fmt.Println("Invalid instruction:", instruction)
		return
	}

	command := matches[1]
	startX, _ := strconv.Atoi(matches[2])
	startY, _ := strconv.Atoi(matches[3])
	endX, _ := strconv.Atoi(matches[4])
	endY, _ := strconv.Atoi(matches[5])

	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			switch command {
			case "turn on":
				grid[y][x]++
			case "turn off":
				if grid[y][x] > 0 {
					grid[y][x]--
				}
			case "toggle":
				grid[y][x] += 2
			}
		}
	}
}

func main() {
	// Initialize the grid with zero brightness
	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	// Open the input file
	file, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read and apply each instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		applyInstruction(grid, instruction)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate the total brightness
	totalBrightness := 0
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			totalBrightness += grid[y][x]
		}
	}

	// Output the total brightness
	fmt.Println("Total brightness of all lights:", totalBrightness)
}
