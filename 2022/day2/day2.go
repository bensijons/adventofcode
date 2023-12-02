package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	score := 0
	scorePt2 := 0

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		res := strings.Split(text, " ")

		outcome := calculateOutcome(res[0], res[1])
		score += outcome.Score

		// Part 2
		myMove := figureOutMySelection(res[0], res[1])
		outcome2 := calculateOutcome(res[0], myMove)
		scorePt2 += outcome2.Score
	}
	fmt.Println("Part 1 Score is:", score)
	fmt.Println("Part 2 Score is:", scorePt2)
}

type Outcome struct {
	Result string
	Score  int
}

func calculateOutcome(theirSelection string, ourSelection string) Outcome {
	switch theirSelection {
	case TheySelectedRock:
		if ourSelection == Rock {
			return Draw(ourSelection)
		} else if ourSelection == Paper {
			return Win(ourSelection)
		} else if ourSelection == Scissors {
			return Loss(ourSelection)
		}
	case TheySelectedPaper:
		if ourSelection == Rock {
			return Loss(ourSelection)
		} else if ourSelection == Paper {
			return Draw(ourSelection)
		} else if ourSelection == Scissors {
			return Win(ourSelection)
		}
	case TheySelectedScissors:
		if ourSelection == Rock {
			return Win(ourSelection)
		} else if ourSelection == Paper {
			return Loss(ourSelection)
		} else if ourSelection == Scissors {
			return Draw(ourSelection)
		}
	}
	return Outcome{}
}

func Draw(selection string) Outcome {
	return Outcome{Result: "DRAW", Score: 3 + ScoreForSelection(selection)}
}

func Win(selection string) Outcome {
	return Outcome{Result: "WIN", Score: 6 + ScoreForSelection(selection)}
}

func Loss(selection string) Outcome {
	return Outcome{Result: "LOSS", Score: ScoreForSelection(selection)}
}

func ScoreForSelection(selection string) int {
	switch selection {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}

const Rock = "X"
const Paper = "Y"
const Scissors = "Z"
const TheySelectedRock = "A"
const TheySelectedPaper = "B"
const TheySelectedScissors = "C"

const LOSE = "X"
const DRAW = "Y"
const WIN = "Z"

func figureOutMySelection(theirSelection string, desiredOutcome string) string {
	switch theirSelection {
	case TheySelectedRock:
		if desiredOutcome == DRAW {
			return Rock
		} else if desiredOutcome == WIN {
			return Paper
		} else if desiredOutcome == LOSE {
			return Scissors
		}
	case TheySelectedPaper:
		if desiredOutcome == LOSE {
			return Rock
		} else if desiredOutcome == DRAW {
			return Paper
		} else if desiredOutcome == WIN {
			return Scissors
		}
	case TheySelectedScissors:
		if desiredOutcome == WIN {
			return Rock
		} else if desiredOutcome == LOSE {
			return Paper
		} else if desiredOutcome == DRAW {
			return Scissors
		}
	}
	return ""
}
