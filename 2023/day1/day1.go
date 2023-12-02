package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	file, err := os.Open("./input_part2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		cv := CalibrationValue{}
		text := scanner.Text()
		for i, t := range text {
			if isInt(t) {
				cv.First = t
				break
			}
			v, ok := isSpelledOutWithLetters(text[i:])
			if ok {
				cv.First = rune(intToRune(v))
				break
			}
		}

		for i := len(text) - 1; i >= 0; i-- {
			r := rune(text[i])
			if isInt(r) {
				cv.Last = r
				break
			}
			v, ok := isSpelledOutWithLetters(reverse(text[:i+1]))
			if ok {
				cv.Last = rune(intToRune(v))
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

func isSpelledOutWithLetters(text string) (int, bool) {
	for i, word := range numericStrings {
		if len(text) >= len(word) {
			if text[:len(word)] == word || text[:len(word)] == reverse(word) {
				return i + 1, true
			}
		}
	}
	return 0, false
}

func runeToInt(r rune) int {
	return int(r - 48)
}

func intToRune(i int) rune {
	return rune(i + 48)
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
