package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	max := 0
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
			if currentSum > max {
				max = currentSum
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

	fmt.Println("CURRENT MAX:", max)
}
