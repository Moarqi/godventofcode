package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/Moarqi/godventofcode/util"
)

func solveFirstPart(lineChannel chan string, isTest bool) {
	signalStartIndex := 0

	line := <-lineChannel
	windowMap := make(map[rune]int)

	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		windowMap[char] += 1
		if i < 3 {
			continue
		}

		if i > 3 {
			lastChar := rune(line[i-4])
			if windowMap[lastChar] > 1 {
				windowMap[lastChar] -= 1
			} else {
				delete(windowMap, rune(line[i-4]))
			}
		}

		if len(windowMap) == 4 {
			fmt.Println("yay")
			signalStartIndex = i + 1
			break
		}
	}

	if isTest {
		if signalStartIndex != 7 {
			fmt.Println(signalStartIndex)
			panic("wrong start index")
		}
	}

	fmt.Println(signalStartIndex)
	fmt.Println("first part solved")
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	messageStartIndex := 0

	line := <-lineChannel
	windowMap := make(map[rune]int)

	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		windowMap[char] += 1
		if i < 13 {
			continue
		}

		if i > 13 {
			lastChar := rune(line[i-14])
			if windowMap[lastChar] > 1 {
				windowMap[lastChar] -= 1
			} else {
				delete(windowMap, rune(line[i-14]))
			}
		}

		if len(windowMap) == 14 {
			fmt.Println("yay")
			messageStartIndex = i + 1
			break
		}
	}

	if isTest {
		if messageStartIndex != 19 {
			fmt.Println(messageStartIndex)
			panic("wrong start index")
		}
	}

	fmt.Println(messageStartIndex)
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
