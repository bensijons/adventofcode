package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	visibleTrees := make(map[string]bool)
	findVisibleTreesLeftAndTop(visibleTrees, lines)
	findVisibleTreesRightAndDown(visibleTrees, lines)

	fmt.Println("Day 8 part 1 answer is:", len(visibleTrees))
}

func findVisibleTreesLeftAndTop(visibleTrees map[string]bool, lines []string) {
	largestTop := make(map[int]int)
	largestLeft := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j, c := range line {
			current, _ := strconv.Atoi(string(c))
			if i == 0 || i == len(lines)-1 {
				largestTop[j] = current
				setVisible(visibleTrees, i, j)
				continue
			}
			if j == 0 || j == len(line)-1 {
				largestLeft = current
				setVisible(visibleTrees, i, j)
				continue
			}
			if current > largestLeft {
				largestLeft = current
				setVisible(visibleTrees, i, j)
			}
			if current > largestTop[j] {
				largestTop[j] = current
				setVisible(visibleTrees, i, j)
			}
		}
	}
}

func findVisibleTreesRightAndDown(visibleTrees map[string]bool, lines []string) {
	largestDown := make(map[int]int)
	largestRight := 0
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		for j := len(line) - 1; j >= 0; j-- {
			current, _ := strconv.Atoi(string(line[j]))
			if i == 0 || i == len(lines)-1 {
				largestDown[j] = current
				setVisible(visibleTrees, i, j)
				continue
			}
			if j == 0 || j == len(line)-1 {
				largestRight = current
				setVisible(visibleTrees, i, j)
				continue
			}
			if current > largestRight {
				largestRight = current
				setVisible(visibleTrees, i, j)
			}
			if current > largestDown[j] {
				largestDown[j] = current
				setVisible(visibleTrees, i, j)
			}
		}
	}
}

func setVisible(visibleTrees map[string]bool, i, j int) {
	key := fmt.Sprintf("%d-%d", i, j)
	visibleTrees[key] = true
}
