package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
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
