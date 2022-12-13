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
	NOOP = iota
	ADDX
)

type Instruction struct {
	instructionType int
	argument        int
	cycles          int
}

func multiplySignalStrengths(signalStrengths []int) int {
	fmt.Println(signalStrengths[19])
	fmt.Println(signalStrengths[59])
	fmt.Println(signalStrengths[99])
	fmt.Println(signalStrengths[139])
	fmt.Println(signalStrengths[179])
	fmt.Println(signalStrengths[219])
	return signalStrengths[19]*20 + signalStrengths[59]*60 + signalStrengths[99]*100 + signalStrengths[139]*140 + signalStrengths[179]*180 + signalStrengths[219]*220
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	registerHistory := make([]int, 240)

	performCycles(lineChannel, &registerHistory)

	finalSignalStrength := multiplySignalStrengths(registerHistory)

	if isTest {
		if finalSignalStrength != 13140 {
			panic("wrong signal strength")
		}
	}

	fmt.Println(finalSignalStrength)
	fmt.Println("first part solved")
}

func drawPixels(registerHistory []int) {
	for i, X := range registerHistory {
		diff := (i % 40) - X
		if diff < 0 {
			diff = -diff
		}

		if diff < 2 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		if (i+1)%40 == 0 {
			fmt.Printf("\n")
		}
	}
}

func parseInstruction(line string) Instruction {
	lineSplit := strings.Split(line, " ")
	instruction := Instruction{
		argument: 0,
		cycles:   1,
	}

	if lineSplit[0] == "noop" {
		instruction.instructionType = NOOP
		return instruction
	} else if lineSplit[0] == "addx" {
		arg, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			panic(err)
		}

		instruction.cycles = 2
		instruction.argument = arg
		instruction.instructionType = ADDX
		return instruction
	} else {
		panic("unknown instruction!")
	}
}

func performInstruction(register, cycleCount *int, registerHistory *[]int, instruction Instruction) {
	if *cycleCount > 239 {
		return
	}

	if instruction.cycles <= 0 {
		if instruction.instructionType == ADDX {
			*register += instruction.argument
		}

		return
	}

	(*registerHistory)[*cycleCount] = *register
	*cycleCount += 1
	instruction.cycles -= 1

	performInstruction(register, cycleCount, registerHistory, instruction)
}

func performCycles(lineChannel chan string, registerHistory *[]int) {
	cycleCount := 0
	register := 1

	for line := range lineChannel {
		instruction := parseInstruction(line)
		performInstruction(&register, &cycleCount, registerHistory, instruction)
		if cycleCount > 239 {
			break
		}
	}
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	registerHistory := make([]int, 240)

	performCycles(lineChannel, &registerHistory)
	drawPixels(registerHistory)

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
