package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numericStrings = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	// part1()
	part2()
}

func part1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		cv := CalibrationValue{}
		text := scanner.Text()
		for _, t := range text {
			if isInt(t) {
				cv.First = t
				break
			}
		}
		for i := len(text) - 1; i >= 0; i-- {
			r := rune(text[i])
			if isInt(r) {
				cv.Last = r
				break
			}
		}
		n := cv.toTwoDigitNumber()
		sum += n
	}

	fmt.Println("day1 part1 answer is: ", sum)
}

func part2() {
	file, err := os.Open("./input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		cv := CalibrationValue{}
		text := scanner.Text()
		lowestNumericStringIndex := -1
		numericStringValue := 0
		// left to right
		for i, nstring := range numericStrings {
			fmt.Println(nstring)
			if substringIndex := strings.Index(text, nstring); substringIndex != -1 {
				fmt.Println("FOUND")
				fmt.Println(substringIndex, lowestNumericStringIndex)
				if substringIndex > lowestNumericStringIndex && lowestNumericStringIndex != -1 {
					lowestNumericStringIndex = substringIndex
					numericStringValue = i + 1
				}
			}

		}
		fmt.Println(lowestNumericStringIndex)
		fmt.Println(numericStringValue)
		fmt.Println("DONE")
		for i, t := range text {
			if isInt(t) {
				if i < lowestNumericStringIndex {
					cv.First = t
				} else {
					cv.First = rune(numericStringValue)
				}
				break
			}
		}

		// right to left
		lowestNumericStringIndex = -1
		numericStringValue = 0
		reverseText := reverse(text)
		for i, nstring := range numericStrings {
			if substringIndex := strings.Index(reverseText, reverse(nstring)); substringIndex != -1 {
				if substringIndex < lowestNumericStringIndex && lowestNumericStringIndex != -1 {
					lowestNumericStringIndex = substringIndex
					numericStringValue = i + 1
				}
			}
		}

		for i := len(text) - 1; i >= 0; i-- {
			r := rune(text[i])
			if isInt(r) {
				if i < lowestNumericStringIndex && lowestNumericStringIndex != -1 {
					cv.Last = rune(numericStringValue)
					break
				}
				cv.Last = r
				break
			}
		}

		n := cv.toTwoDigitNumber()
		sum += n
	}

	fmt.Println("day1 part2 answer is:", sum)
}

type CalibrationValue struct {
	First rune
	Last  rune
}

func (c CalibrationValue) toTwoDigitNumber() int {
	if !isInt(c.First) {
		log.Fatal("invalid c.First found: ", c.First)
	}

	if !isInt(c.Last) {
		log.Fatal("invalid c.Last found: ", c.Last)
	}

	f := fmt.Sprint(runeToInt(c.First))
	l := fmt.Sprint(runeToInt(c.Last))

	val := f + l

	v, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("error during strconv")
	}
	return v
}

func isInt(r rune) bool {
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func runeToInt(r rune) int {
	return int(r - 48)
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
