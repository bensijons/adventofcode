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

		folders = ParseCommands(folders, line, sizeMap)
	}
	CalculateTotalSize(sizeMap)
}

func ParseCommands(folders []string, line string, sizeMap map[string]int) []string {
	w := strings.Split(line, " ")

	if w[0] == "$" {
		if w[1] == "cd" && w[2] != ".." {
			folder := w[2]
			parentFolder := getParentFolderPath(folders)
			folders = append(folders, getFolderPath(parentFolder, folder))
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
	return folders
}

func getParentFolderPath(folders []string) string {
	if len(folders) == 0 {
		return ""
	}

	return folders[len(folders)-1]
}

func getFolderPath(parentFolderPath string, folder string) string {
	if parentFolderPath == "" {
		return folder
	}
	return fmt.Sprintf("%s%s/", parentFolderPath, folder)
}

func CalculateTotalSize(sizeMap map[string]int) {
	sum := int(0)
	for k, folderSize := range sizeMap {
		if folderSize <= 100000 {
			sum += folderSize
		}
		fmt.Println("Folder", k, ", size", folderSize)
	}
	fmt.Println("sum", sum)
}
