package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"

	"github.com/Moarqi/godventofcode/util"
)

func readStacks(lineChannel chan string) [][]rune {
	stacks := make([][]rune, 0, 0)

	for line := range lineChannel {
		if len(line) == 0 {
			break
		}
		lineAsRunes := []rune(line)

		for i := range lineAsRunes {
			if i%4 != 0 {
				continue
			}

			if lineAsRunes[i] == '[' {
				crate := lineAsRunes[i+1]

				stackIndex := i / 4

				if len(stacks) <= stackIndex {
					for j := len(stacks); j <= stackIndex; j += 1 {
						stacks = append(stacks, make([]rune, 0))
					}
				}

				stacks[stackIndex] = append(stacks[stackIndex], crate)
			}
		}
	}

	// topmost element in stack: stack[0]
	return stacks
}

func parseMove(instruction string) (int, int, int) {
	var from, to, count int

	splitRegex := regexp.MustCompile("(\\s?[a-z]+\\s?)")

	instructionList := splitRegex.Split(instruction, -1)
	count, err := strconv.Atoi(instructionList[1])
	if err != nil {
		panic(err)
	}
	from, err = strconv.Atoi(instructionList[2])
	if err != nil {
		panic(err)
	}
	to, err = strconv.Atoi(instructionList[3])
	if err != nil {
		panic(err)
	}

	return from - 1, to - 1, count
}

func performMove(stacks [][]rune, from, to, count int) {
	for c := 0; c < count; c += 1 {
		topCrate := stacks[from][0]
		stacks[from] = stacks[from][1:]

		stacks[to] = append(stacks[to], 0)
		for j := len(stacks[to]) - 1; j > 0; j -= 1 {
			stacks[to][j] = stacks[to][j-1]
		}
		stacks[to][0] = topCrate
	}
}

func performMoveStacked(stacks [][]rune, from, to, count int) {
	topCrates := stacks[from][0:count]
	stacks[from] = stacks[from][count:]

	for j := 0; j < count; j += 1 {
		stacks[to] = append(stacks[to], 0)
	}

	for i := 0; i < count; i += 1 {
		for j := len(stacks[to]) - 1; j > 0+i; j -= 1 {
			stacks[to][j] = stacks[to][j-1]
		}

		stacks[to][i] = topCrates[i]
	}
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	var topCrates string
	stacks := readStacks(lineChannel)

	for line := range lineChannel {
		from, to, count := parseMove(line)
		fmt.Println(from, to, count)
		performMove(stacks, from, to, count)
		fmt.Println(stacks)
	}

	for _, stack := range stacks {
		topCrates += string(stack[0])
	}

	if isTest {
		if topCrates != "CMZ" {
			fmt.Println(topCrates)
			panic("wrong top crates")
		}
	}

	fmt.Println(topCrates)
	fmt.Println("first part solved")
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	var topCrates string
	stacks := readStacks(lineChannel)

	for line := range lineChannel {
		from, to, count := parseMove(line)
		performMoveStacked(stacks, from, to, count)
	}

	for _, stack := range stacks {
		topCrates += string(stack[0])
	}

	if isTest {
		if topCrates != "MCD" {
			fmt.Println(topCrates)
			panic("wrong top crates")
		}
	}

	fmt.Println(topCrates)
	fmt.Println("second part solved")
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot retrieve current filename")
	}

	isTest := false
	if len(os.Args) > 2 && os.Args[2] == "--test" {
		isTest = true
	}

	lineChannel := make(chan string)
	if isTest {
		go util.ReadInput(path.Dir(filename)+"/test.txt", lineChannel, false)
	} else {
		go util.ReadInput(path.Dir(filename)+"/input.txt", lineChannel, false)
	}

	if len(os.Args) > 1 && os.Args[1] == "1" {
		solveFirstPart(lineChannel, isTest)
	} else if len(os.Args) > 1 && os.Args[1] == "2" {
		solveSecondPart(lineChannel, isTest)
	} else {
		fmt.Println("part not specified")
	}
}
