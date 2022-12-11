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
	file, err := os.Open("./day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sizeMap := make(map[string]int)
	folders := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		w := strings.Split(line, " ")

		if w[0] == "$" {
			if w[1] == "cd" && w[2] != ".." {
				folder := w[2]
				folders = append(folders, folder)
			} else if w[1] == "cd" && w[2] == ".." {
				folders = folders[:len(folders)-1]
			}
		} else if w[0] != "dir" {
			size, err := strconv.Atoi(w[0])
			if err != nil {
				log.Fatal(err)
			}
			for _, folder := range folders {
				sizeMap[folder] = sizeMap[folder] + size
			}
		}
	}
	CalculateTotalSize(sizeMap)
}

func CalculateTotalSize(sizeMap map[string]int) {
	sum := int(0)
	for _, folderSize := range sizeMap {
		if folderSize <= 100000 {
			sum += folderSize
		}
		fmt.Println("Folder", k, ", size", folderSize)
	}
	fmt.Println("sum", sum)
}
