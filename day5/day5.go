package main

import "fmt"

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2

func main() {
	stackMap := map[int][]string{}
	stackMap[1] = []string{"Z", "N", "D"}
	stackMap[2] = []string{"M", "C"}
	stackMap[3] = []string{"P"}
	move(stackMap, 1, 2, 1)
	for i := 1; i <= len(stackMap); i++ {
		sl := stackMap[i]
		fmt.Println(sl[len(sl)-1])
	}

	fmt.Println("Day 5 part 1 answer is")
}

func move(stackMap map[int][]string, count int, from int, to int) {
	toStack := stackMap[to]
	fromStack := stackMap[from]

	toStack = append(toStack, fromStack[len(fromStack)-1])
	fromStack = fromStack[:len(fromStack)-1]

	stackMap[from] = fromStack
	stackMap[to] = toStack
}
