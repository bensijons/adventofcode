package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1()
	part2()
}

func part2() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal()
	}

	score := 0
	group := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		group = append(group, rucksack)
		if len(group) == 3 {
			score += findGroupBadgeItemPriority(group)
			group = []string{}
		}
	}

	fmt.Println("Part 2 score is:", score)
}

func findGroupBadgeItemPriority(group []string) int {
	rucksack1 := group[0]
	rucksack2 := group[1]
	rucksack3 := group[2]

	badgeItem := ""
	for _, item1 := range rucksack1 {
		for _, item2 := range rucksack2 {
			if item1 == item2 {
				for _, item3 := range rucksack3 {
					if item1 == item3 {
						badgeItem = string(item1)
						break
					}
				}
				break
			}
		}
	}

	return priorityMap[badgeItem]
}

func part1() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal()
	}

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := scanner.Text()
		firstCompartment := text[:len(text)/2]
		secondCompartment := text[len(text)/2:]

		foundDuplicate := false
		for _, item := range firstCompartment {
			if foundDuplicate {
				break
			}
			for _, item2 := range secondCompartment {
				if item == item2 && !foundDuplicate {
					score += priorityMap[string(item)]
					foundDuplicate = true
				}
			}
		}
	}

	fmt.Println("Part 1 score is:", score)
}

var priorityMap = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}
