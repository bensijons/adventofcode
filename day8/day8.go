package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day8/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	visibleTrees := make(map[string]bool)
	scenicScore := make(map[string]int)
	findVisibleTreesLeftAndTop(visibleTrees, lines, scenicScore)
	findVisibleTreesRightAndDown(visibleTrees, lines, scenicScore)

	fmt.Println("Day 8 part 1 answer is:", len(visibleTrees))

	fmt.Println(scenicScore)
}

func findVisibleTreesLeftAndTop(visibleTrees map[string]bool, lines []string, scenicScore map[string]int) {
	largestTop := make(map[int]int)
	largestLeft := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j, c := range line {
			key := fmt.Sprintf("%d-%d", i, j)
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
				calculateScenicScore(scenicScore, key, j)
				largestLeft = current
				setVisible(visibleTrees, i, j)
			}
			if current > largestTop[j] {
				calculateScenicScore(scenicScore, key, i)
				largestTop[j] = current
				setVisible(visibleTrees, i, j)
			}
		}
	}
}

func findVisibleTreesRightAndDown(visibleTrees map[string]bool, lines []string, scenicScore map[string]int) {
	largestDown := make(map[int]int)
	largestRight := 0
	downCounter := 0
	for i := len(lines) - 1; i >= 0; i-- {
		downCounter++
		rightCounter := 0
		line := lines[i]

		for j := len(line) - 1; j >= 0; j-- {
			rightCounter++
			key := fmt.Sprintf("%d-%d", i, j)
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
				calculateScenicScore(scenicScore, key, rightCounter)
				largestRight = current
				setVisible(visibleTrees, i, j)
			}
			if current > largestDown[j] {
				calculateScenicScore(scenicScore, key, downCounter)
				largestDown[j] = current
				setVisible(visibleTrees, i, j)
			}
		}
	}
}

func calculateScenicScore(scenicScore map[string]int, key string, multiplier int) {
	val, ok := scenicScore[key]
	if ok {
		scenicScore[key] = val * multiplier
	} else {
		scenicScore[key] = multiplier
	}
}

func setVisible(visibleTrees map[string]bool, i, j int) {
	key := fmt.Sprintf("%d-%d", i, j)
	visibleTrees[key] = true
}
