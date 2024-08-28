package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile("puzzle.txt")
}

//readFile function opens the file using os.Open function
//Then reads it line by line using bufio scanner
func readFile(s string) {
	// Open the file
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	results := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()

		line := strings.Split(lines, "x")
		length := line[0]
		width := line[1]
		height := line[2]

		len, err := strconv.Atoi(length)
		if err != nil {
			log.Fatal(err)
		}

		wid, err := strconv.Atoi(width)
		if err != nil {
			log.Fatal(err)
		}

		hei, err := strconv.Atoi(height)
		if err != nil {
			log.Fatal(err)
		}

		area1 := 2 * (len * wid)
		area2 := 2 * (len * hei)
		area3 := 2 * (wid * hei)

		total += area1 + area2 + area3


		results = append(results, area1/2, area2/2, area3/2)

		sort.Ints(results)
		total+=results[0]
		//fmt.Println(results)
		results=nil
	}
	fmt.Println(total)
}
