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
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	initializing := true
	initialStackLines := []string{}
	crates := map[int][]string{}
	cratesPartTwo := map[int][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if initializing && line == "" {
			initializing = false
			initializeStack(crates, initialStackLines)
			initializeStack(cratesPartTwo, initialStackLines)
		} else if initializing {
			initialStackLines = append(initialStackLines, line)
		} else {
			// initialization done, start moving crates
			count, from, to := readNextMove(line)
			moveCrates(crates, count, from, to)
			moveCratesPartTwo(cratesPartTwo, count, from, to)
		}
	}
	printAnswer(crates)
	printAnswer(cratesPartTwo)
}

func printAnswer(crates map[int][]string) {
	answer := ""
	for i := 1; i <= len(crates); i++ {
		sl := crates[i]
		answer = fmt.Sprintf("%s%s", answer, sl[len(sl)-1])
	}
	fmt.Println(answer)
}

// initializeStack, accepting an empty stack and lines that look like the following:
//
// [N]     [C]                 [Q]
// [W]     [J] [L]             [J] [V]
// [F]     [N] [D]     [L]     [S] [W]
// [R] [S] [F] [G]     [R]     [V] [Z]
// [Z] [G] [Q] [C]     [W] [C] [F] [G]
// [S] [Q] [V] [P] [S] [F] [D] [R] [S]
// [M] [P] [R] [Z] [P] [D] [N] [N] [M]
// [D] [W] [W] [F] [T] [H] [Z] [W] [R]
//  1   2   3   4   5   6   7   8   9
//
// altering the map into a map where the bottom
// line is the key of each entry in the map
// and the values are read from bottom up
func initializeStack(stack map[int][]string, lines []string) {
	for lineIndex := len(lines) - 2; lineIndex >= 0; lineIndex-- {
		for crateIndex, k := 1, 1; crateIndex < len(lines[lineIndex]); crateIndex += 4 {
			crate := string(lines[lineIndex][crateIndex])
			if crate != " " {
				crates := stack[k]
				crates = append(crates, crate)
				stack[k] = crates
			}
			k++
		}
	}
}

func readNextMove(line string) (count int, from int, to int) {
	l := strings.Split(line, " ")
	count, err := strconv.Atoi(l[1])
	if err != nil {
		log.Fatal(err)
	}
	from, err = strconv.Atoi(l[3])
	if err != nil {
		log.Fatal(err)
	}
	to, err = strconv.Atoi(l[5])
	if err != nil {
		log.Fatal(err)
	}
	return count, from, to
}

func moveCrates(stackMap map[int][]string, count int, from int, to int) {
	for i := 0; i < count; i++ {
		toStack := stackMap[to]
		fromStack := stackMap[from]

		toStack = append(toStack, fromStack[len(fromStack)-1])
		fromStack = fromStack[:len(fromStack)-1]

		stackMap[from] = fromStack
		stackMap[to] = toStack
	}
}

func moveCratesPartTwo(stackMap map[int][]string, count int, from int, to int) {
	toStack := stackMap[to]
	fromStack := stackMap[from]

	fl := len(fromStack)
	toStack = append(toStack, fromStack[fl-count:fl]...)
	fromStack = fromStack[:fl-count]

	stackMap[from] = fromStack
	stackMap[to] = toStack
}
