package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gameboard struct {
	fields  [][]int
	crossed [][]bool
	score   int
	won     bool
}

func readInput(filePath string, output chan string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	defer close(output)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		output <- scanner.Text()
		if err != nil {
			panic(err)
		}
	}
}

func calcScore(board *gameboard, newNumber int) {
	uncrossedSum := 0
	for i := 0; i < len(board.fields); i++ {
		for j := 0; j < len(board.fields[i]); j++ {
			if !board.crossed[i][j] {
				uncrossedSum += board.fields[i][j]
			}
		}
	}

	board.score = uncrossedSum * newNumber
}

func cross(board *gameboard, newNumber int) bool {
	if board.won {
		return false
	}

	for i := 0; i < len(board.fields); i++ {
		for j := 0; j < len(board.fields[i]); j++ {
			if board.fields[i][j] == newNumber {
				board.crossed[i][j] = true
				if checkCrosses(board, i, j) {
					calcScore(board, newNumber)
					board.won = true
					return true
				}
			}
		}
	}

	return false
}

func checkCrossedRow(board *gameboard, row int) bool {
	for j := 0; j < len(board.fields[row]); j++ {
		if !board.crossed[row][j] {
			return false
		}
	}
	return true
}

func checkCrossedCol(board *gameboard, col int) bool {
	for i := 0; i < len(board.fields); i++ {
		if !board.crossed[i][col] {
			return false
		}
	}
	return true
}

func checkCrossedDia(board *gameboard) bool {
	passed := true
	for i := 0; i < len(board.fields); i++ {
		if !board.crossed[i][i] {
			passed = false
			break
		}
	}
	if passed {
		return true
	} else {
		passed = true
	}

	for i := 0; i < len(board.fields); i++ {
		if !board.crossed[len(board.fields)-1-i][i] {
			passed = false
			break
		}
	}

	return passed
}

func checkCrosses(board *gameboard, row int, col int) bool {
	return checkCrossedCol(board, col) || checkCrossedRow(board, row)
}

func firstWin() {
	lineChannel := make(chan string)
	go readInput("/home/markus/dev/godventofcode/04/data/input", lineChannel)

	crossedValues := make([]int, 0, 1024)
	initialLine := <-lineChannel
	crossValueSplits := strings.Split(initialLine, ",")

	for i := 0; i < len(crossValueSplits); i++ {
		crossedValue, err := strconv.Atoi(crossValueSplits[i])

		if err != nil {
			panic(err)
		}

		crossedValues = append(crossedValues, crossedValue)
	}

	gameBoards := make([]gameboard, 0, 100)

	// gameBoardLineIdx := 0
	gameBoardIndex := 0

	for line := range lineChannel {
		if line == "" {
			continue
		}

		fields := make([][]int, 5)
		crossed := make([][]bool, 5)
		newBoard := gameboard{fields: fields, crossed: crossed}

		gameBoardLineIdx := 0
		for line != "" {
			splits := strings.Split(line, " ")
			lineArray := make([]int, 5)
			lineCrossedArray := make([]bool, 5)

			lineIdx := 0
			for i := 0; i < len(splits); i++ {
				if splits[i] == "" {
					continue
				}

				intValue, err := strconv.Atoi(splits[i])

				if err != nil {
					panic(err)
				}

				lineArray[lineIdx] = intValue
				lineIdx += 1
			}

			newBoard.fields[gameBoardLineIdx] = lineArray
			newBoard.crossed[gameBoardLineIdx] = lineCrossedArray
			gameBoardLineIdx += 1

			line = <-lineChannel
		}

		gameBoards = append(gameBoards, newBoard)
		gameBoardIndex += 1
	}

	for i := 0; i < len(crossedValues); i++ {
		val := crossedValues[i]
		fmt.Println(val)
		won := false
		for gbIdx := 0; gbIdx < len(gameBoards); gbIdx++ {
			if cross(&gameBoards[gbIdx], val) {
				fmt.Println(gameBoards[gbIdx].score)
				fmt.Println(gameBoards[gbIdx])
				won = true
				break
			}
		}
		if won {
			break
		}
	}
}

func lastWin() {
	lineChannel := make(chan string)
	go readInput("/home/markus/dev/godventofcode/04/data/input", lineChannel)

	crossedValues := make([]int, 0, 1024)
	initialLine := <-lineChannel
	crossValueSplits := strings.Split(initialLine, ",")

	for i := 0; i < len(crossValueSplits); i++ {
		crossedValue, err := strconv.Atoi(crossValueSplits[i])

		if err != nil {
			panic(err)
		}

		crossedValues = append(crossedValues, crossedValue)
	}

	gameBoards := make([]gameboard, 0, 100)

	// gameBoardLineIdx := 0
	gameBoardIndex := 0

	for line := range lineChannel {
		if line == "" {
			continue
		}

		fields := make([][]int, 5)
		crossed := make([][]bool, 5)
		newBoard := gameboard{fields: fields, crossed: crossed}

		gameBoardLineIdx := 0
		for line != "" {
			splits := strings.Split(line, " ")
			lineArray := make([]int, 5)
			lineCrossedArray := make([]bool, 5)

			lineIdx := 0
			for i := 0; i < len(splits); i++ {
				if splits[i] == "" {
					continue
				}

				intValue, err := strconv.Atoi(splits[i])

				if err != nil {
					panic(err)
				}

				lineArray[lineIdx] = intValue
				lineIdx += 1
			}

			newBoard.fields[gameBoardLineIdx] = lineArray
			newBoard.crossed[gameBoardLineIdx] = lineCrossedArray
			gameBoardLineIdx += 1

			line = <-lineChannel
		}

		gameBoards = append(gameBoards, newBoard)
		gameBoardIndex += 1
	}

	wonIdx := make([]int, 0, len(gameBoards))

	for i := 0; i < len(crossedValues); i++ {
		val := crossedValues[i]
		for gbIdx := 0; gbIdx < len(gameBoards); gbIdx++ {
			if cross(&gameBoards[gbIdx], val) {
				wonIdx = append(wonIdx, gbIdx)
			}
		}

		if len(wonIdx) == len(gameBoards) {
			lastWonIdx := wonIdx[len(wonIdx)-1]
			fmt.Println(lastWonIdx)
			fmt.Println(gameBoards[lastWonIdx].score)
			break
		}
	}
}

func main() {
	lastWin()
}
