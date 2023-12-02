package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("Day 6 part 1 answer is:", FindMarkerIndex(input, 4))
		fmt.Println("Day 6 part 2 answer is:", FindMarkerIndex(input, 14))
	}
}

func FindMarkerIndex(input string, markerCount int) int {
	for i := 0; i < len(input)-markerCount-1; i++ {
		if hasNoDuplicateChars(input[i : i+markerCount]) {
			return i + markerCount
		}
	}
	fmt.Println("No marker found")
	return -1
}

func hasNoDuplicateChars(input string) bool {
	if len(input) == 0 {
		return true
	}
	contains := false
	contains = strings.Contains(input[1:], string(input[0]))
	if contains {
		return false
	}
	contains = hasNoDuplicateChars(input[1:])
	return contains
}
