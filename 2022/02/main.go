package main

import (
	"fmt"
	"strings"

	"github.com/Moarqi/godventofcode/util"
)

func roundWon(ownDraft string, opponentDraft string) bool {
	if ownDraft == "X" && opponentDraft == "C" {
		return true
	}

	if ownDraft == "Y" && opponentDraft == "A" {
		return true
	}

	if ownDraft == "Z" && opponentDraft == "B" {
		return true
	}

	return false
}

func choiceScore(draft string) int {
	if draft == "X" {
		return 1
	}

	if draft == "Y" {
		return 2
	}

	return 3
}

func isDraft(ownDraft string, opponentDraft string) bool {
	if ownDraft == "X" && opponentDraft == "A" {
		return true
	}
	if ownDraft == "Y" && opponentDraft == "B" {
		return true
	}
	if ownDraft == "Z" && opponentDraft == "C" {
		return true
	}
	return false
}

func solveFirstPart(lineChannel chan string) {
	// first column: opponents draft
	// A Rock
	// B Paper
	// C Scissors

	// second column my draft
	// X Rock
	// Y Paper
	// Z Scissors

	// shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
	// + (0 if you lost, 3 if the round was a draw, and 6 if you won).
	totalScore := 0

	for line := range lineChannel {
		drafts := strings.Split(line, " ")
		opponent := drafts[0]
		own := drafts[1]

		outcomeScore := 0

		if isDraft(own, opponent) {
			outcomeScore = 3
		} else if roundWon(own, opponent) {
			outcomeScore = 6
		}

		fmt.Println(outcomeScore, choiceScore(own))

		totalScore += outcomeScore + choiceScore(own)
	}

	fmt.Println(totalScore)
}

func getChoice(opponent, desiredOutcome string) string {
	if desiredOutcome == "Y" {
		if opponent == "A" {
			return "X"
		}
		if opponent == "B" {
			return "Y"
		}
		if opponent == "C" {
			return "Z"
		}
	}

	if desiredOutcome == "X" {
		if opponent == "A" {
			return "Z"
		}
		if opponent == "B" {
			return "X"
		}
		if opponent == "C" {
			return "Y"
		}
	}

	if desiredOutcome == "Z" {
		if opponent == "A" {
			return "Y"
		}
		if opponent == "B" {
			return "Z"
		}
		if opponent == "C" {
			return "X"
		}
	}

	panic("what")
}

func solveSecondPart(lineChannel chan string) {
	// first column: opponents draft
	// A Rock
	// B Paper
	// C Scissors

	// second column: desired outcome
	// X lose
	// Y draw
	// Z win
	totalScore := 0

	for line := range lineChannel {
		drafts := strings.Split(line, " ")
		opponent := drafts[0]
		desiredOutcome := drafts[1]

		outcomeScore := 0

		own := getChoice(opponent, desiredOutcome)

		if desiredOutcome == "Y" {
			outcomeScore = 3
		} else if desiredOutcome == "Z" {
			outcomeScore = 6
		}

		totalScore += outcomeScore + choiceScore(own)
	}

	fmt.Println(totalScore)
}

func main() {
	lineChannel := make(chan string)
	go util.ReadInput("/home/markus/dev/godventofcode/2022/02/input.txt", lineChannel, true)

	// solveFirstPart(lineChannel)
	solveSecondPart(lineChannel)
}
