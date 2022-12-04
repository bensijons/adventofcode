package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Your total score is the sum of your scores for each round.
	// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
	// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

	score := 0

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
	}
	fmt.Println("Score is:", score)
}

type Outcome struct {
	Result string
	Score  int
}

func calculateOutcome(theirSelection string, ourSelection string) Outcome {
	switch theirSelection {
	case "A":
		if ourSelection == "X" {
			return Draw(ourSelection)
		} else if ourSelection == "Y" {
			return Win(ourSelection)
		} else if ourSelection == "Z" {
			return Loss(ourSelection)
		}
	case "B":
		if ourSelection == "X" {
			return Loss(ourSelection)
		} else if ourSelection == "Y" {
			return Draw(ourSelection)
		} else if ourSelection == "Z" {
			return Win(ourSelection)
		}
	case "C":
		if ourSelection == "X" {
			return Win(ourSelection)
		} else if ourSelection == "Y" {
			return Loss(ourSelection)
		} else if ourSelection == "Z" {
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
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}
