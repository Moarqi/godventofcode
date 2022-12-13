package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/Moarqi/godventofcode/util"
)

const (
	start = iota
	empty
	visited
	head
	tail
)

const (
	up = iota
	right
	down
	left
)

var dirMap = map[string]int{
	"U": up,
	"D": down,
	"R": right,
	"L": left,
}

type Command struct {
	direction, amount int
}

type Coord struct {
	i, j int
}

type MapState struct {
	start Coord
	rope  []Coord
}

func generateMap(ropLength int) ([][]int, MapState) {
	/**
	......
	......
	......
	......
	H.....
	*/
	gridMap := make([][]int, 1000)
	for i := range gridMap {
		gridMap[i] = make([]int, 1000)
		for j := range gridMap[i] {
			gridMap[i][j] = empty
		}
	}

	// gridMap[4][0] = head
	rope := make([]Coord, ropLength)
	for i := range rope {
		rope[i].i = 400
		rope[i].j = 400
	}

	state := MapState{
		start: Coord{400, 400},
		rope:  rope,
	}

	return gridMap, state
}

func printMap(gridMap *[][]int, gridState *MapState, reset bool) {
	head := gridState.rope[0]
	for i, row := range *gridMap {
		fmt.Printf("\r")
		for j, field := range row {
			ropeOccupied := false
			for ropeId := len(gridState.rope) - 1; ropeId > 0; ropeId -= 1 {
				ropeElement := gridState.rope[ropeId]
				if i == ropeElement.i && j == ropeElement.j {
					fmt.Printf("%d", ropeId)
					ropeOccupied = true
					break
				}
			}

			if i == head.i && j == head.j {
				if ropeOccupied {
					fmt.Printf("\b")
				}
				fmt.Printf("H")
			} else if i == gridState.start.i && j == gridState.start.j {
				if ropeOccupied {
					fmt.Printf("\b")
				}
				fmt.Printf("s")
			} else if field == visited {
				if ropeOccupied {
					fmt.Printf("\b")
				}
				fmt.Printf("#")
			} else if !ropeOccupied {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	if reset {
		for range *gridMap {
			fmt.Printf("\033[1A")
		}
	}
}

func parseCommand(line string) Command {
	inputSplit := strings.Split(line, " ")

	command := Command{}
	dirAsString := inputSplit[0]

	amount, err := strconv.Atoi(inputSplit[1])
	if err != nil {
		panic(err)
	}

	command.amount = amount
	command.direction = dirMap[dirAsString]

	return command
}

func follow(head, tail *Coord) {
	iDiff := head.i - tail.i
	jDiff := head.j - tail.j

	if jDiff < 0 {
		jDiff = -jDiff
	}
	if iDiff < 0 {
		iDiff = -iDiff
	}
	if jDiff+iDiff > 1 {
		if iDiff+jDiff > 3 {
			if head.j-tail.j > 0 {
				tail.j += 1
			} else {
				tail.j -= 1
			}
			if head.i-tail.i > 0 {
				tail.i += 1
			} else {
				tail.i -= 1
			}
		} else if jDiff+iDiff > 2 {
			if iDiff < jDiff {
				tail.i = head.i
				if head.j-tail.j > 0 {
					tail.j += 1
				} else {
					tail.j -= 1
				}
			} else {
				tail.j = head.j
				if head.i-tail.i > 0 {
					tail.i += 1
				} else {
					tail.i -= 1
				}
			}
		}

		if iDiff == 0 {
			if head.j-tail.j > 0 {
				tail.j += 1
			} else {
				tail.j -= 1
			}
		} else if jDiff == 0 {
			if head.i-tail.i > 0 {
				tail.i += 1
			} else {
				tail.i -= 1
			}
		}
	}

}

func executeCommand(gridMapPtr *[][]int, mapStatePtr *MapState, command Command) {
	if command.amount < 1 {
		return
	}
	// printMap(gridMapPtr, mapStatePtr, true)
	head := &mapStatePtr.rope[0]
	if command.direction == right {
		head.j += 1
	} else if command.direction == left {
		head.j -= 1
	} else if command.direction == up {
		head.i -= 1
	} else if command.direction == down {
		head.i += 1
	}

	for ropeId := 1; ropeId < len(mapStatePtr.rope); ropeId += 1 {
		head := &mapStatePtr.rope[ropeId-1]
		tail := &mapStatePtr.rope[ropeId]

		if ropeId == len(mapStatePtr.rope)-1 {
			(*gridMapPtr)[tail.i][tail.j] = visited
		}

		follow(head, tail)

		if ropeId == len(mapStatePtr.rope)-1 {
			(*gridMapPtr)[tail.i][tail.j] = visited
		}
	}

	command.amount -= 1
	// fmt.Scanln()
	// fmt.Printf("\033[1A")
	// time.Sleep(300 * time.Millisecond)
	executeCommand(gridMapPtr, mapStatePtr, command)
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	visitedField := 0
	gridMap, mapState := generateMap(2)

	for line := range lineChannel {
		command := parseCommand(line)
		executeCommand(&gridMap, &mapState, command)
	}
	printMap(&gridMap, &mapState, false)

	for i := range gridMap {
		for j := range gridMap[i] {
			if gridMap[i][j] == visited {
				visitedField += 1
			}
		}
	}

	if isTest {
		if visitedField != 13 {
			fmt.Println(visitedField)
			panic("wrong count")
		}
	}

	fmt.Println(visitedField)
	fmt.Println("first part solved")
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	visitedField := 0
	gridMap, mapState := generateMap(10)

	for line := range lineChannel {
		command := parseCommand(line)
		executeCommand(&gridMap, &mapState, command)
	}
	// printMap(&gridMap, &mapState, false)

	for i := range gridMap {
		for j := range gridMap[i] {
			if gridMap[i][j] == visited {
				visitedField += 1
			}
		}
	}

	if isTest {
		if visitedField != 36 {
			fmt.Println(visitedField)
			panic("wrong count")
		}
	}

	fmt.Println(visitedField)
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
