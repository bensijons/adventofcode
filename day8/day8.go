package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day8/test.txt")
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grid := TreeGrid{}
	numberOfVisibleTrees := 0

	for i := 0; i < len(lines); i++ {
		if i == 0 {
			numberOfVisibleTrees += len(lines[0])
			continue
		}
		grid.previous = &lines[i-1]
		grid.current = &lines[i]

		if i == len(lines)-1 {
			grid.next = nil
		} else {
			grid.next = &lines[i+1]
		}
		fmt.Println(i)
		numberOfVisibleTrees += countInnerVisibleTrees(grid)
	}
	fmt.Println(numberOfVisibleTrees)
}

type TreeGrid struct {
	previous *string
	current  *string
	next     *string
}

func countInnerVisibleTrees(grid TreeGrid) int {
	if grid.previous == nil || grid.next == nil {
		return len(*grid.current)
	}

	count := 0

	for i := 0; i < len(*grid.current)-1; i++ {
		if i == 0 || i == len(*grid.current)-1 {
			count++
			continue
		}
		current := *grid.current
		previous := *grid.previous
		next := *grid.next
		visibleUp := current[i] > previous[i]
		visibleDown := current[i] > next[i]
		visibleLeft := current[i] > current[i-1]
		visibleRight := current[i] > current[i+1]
		fmt.Println(i)
		fmt.Println("visibleDown", visibleDown, visibleUp, visibleLeft, visibleRight)

		if visibleUp || visibleDown || visibleLeft || visibleRight {
			count++
		}
	}

	return count
}
