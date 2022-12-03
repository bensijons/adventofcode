package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	max := [3]int{}
	currentSum := 0

	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if currentSum > max[2] {
				max = [3]int{max[1], max[2], currentSum}
			} else if currentSum > max[1] {
				max = [3]int{max[1], currentSum, max[2]}
			} else if currentSum > max[2] {
				max = [3]int{currentSum, max[1], max[2]}
			}
			currentSum = 0
		} else {
			i, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			currentSum += i
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 answer:", max[2])

	p2 := 0
	for _, m := range max {
		p2 += m
	}
	fmt.Println("Part 2 answer:", p2)
}
