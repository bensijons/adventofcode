package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	FindFolderToDelete(sizeMap, 70000000)
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
	for _, folderSize := range sizeMap {
		if folderSize <= 100000 {
			sum += folderSize
		}
	}
	fmt.Println("sum", sum)
}

type Folder struct {
	Name string
	Size int
}

func FindFolderToDelete(sizeMap map[string]int, totalSpace int) {
	unusedSpace := totalSpace - sizeMap["/"]
	desiredSpace := 30000000
	spaceToFreeUp := desiredSpace - unusedSpace

	folders := []Folder{}
	for folderName, folderSize := range sizeMap {
		folders = append(folders, Folder{Name: folderName, Size: folderSize})
	}

	sort.Slice(folders, func(i, j int) bool {
		return folders[i].Size < folders[j].Size
	})

	for _, folder := range folders {
		if folder.Size > spaceToFreeUp {
			fmt.Println("The folder to delete is", folder)
			break
		}
	}
}
