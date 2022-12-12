package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day8/input.txt")
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prev
	}
}

func ScanTree(prev string, current string, next string) {

}
