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
		fmt.Println("Day 6 part 1 answer is:", FindMarkerIndex(input))
	}
}

func FindMarkerIndex(input string) int {
	for i := 0; i < len(input)-3; i++ {
		if hasNoDuplicateChars(input[i : i+4]) {
			return i + 4
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
