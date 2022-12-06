package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)

	count := 0
	part2count := 0

	for scanner.Scan() {
		pair := scanner.Text()
		elf1, elf2 := splitPair(pair)
		assignment1 := parseToAssignment(elf1)
		assignment2 := parseToAssignment(elf2)

		if fullyContains(assignment1, assignment2) {
			count++
		}
		if overlaps(assignment1, assignment2) {
			part2count++
		}
	}

	fmt.Println("Part 1 answer is:", count)
	fmt.Println("Part 2 answer is:", part2count)
}

type Assignment struct {
	Start int
	End   int
}

// Check if start and end of b is fully contained within a or vice versa
func fullyContains(a Assignment, b Assignment) bool {
	return a.Start <= b.Start && b.End <= a.End ||
		b.Start <= a.Start && a.End <= b.End
}

func overlaps(a Assignment, b Assignment) bool {
	return a.Start <= b.Start && a.End >= b.Start || b.Start <= a.Start && b.End >= a.Start
}

func splitPair(pair string) (string, string) {
	s := strings.Split(pair, ",")
	return s[0], s[1]
}

func parseToAssignment(input string) Assignment {
	a := strings.Split(input, "-")
	start, err := strconv.Atoi(a[0])
	if err != nil {
		log.Fatal()
	}
	end, err := strconv.Atoi(a[1])
	if err != nil {
		log.Fatal()
	}
	return Assignment{Start: start, End: end}
}
