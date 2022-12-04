package main

import (
	"fmt"

	"github.com/Moarqi/godventofcode/util"
)

func solveFirstPart(lineChannel chan string, isTest bool) {
	prioritySum := 0

	for line := range lineChannel {
		chars := []rune(line)
		firstCompartmentRunes := chars[:len(chars)/2]
		secondCompartmentRunes := chars[len(chars)/2:]

		runeMap := make(map[rune]bool)
		runeIntersectionMap := make(map[rune]bool)
		for _, item := range firstCompartmentRunes {
			runeMap[item] = true
		}

		for _, item := range secondCompartmentRunes {
			if runeMap[item] {
				runeIntersectionMap[item] = true
			}
		}

		for duplicate := range runeIntersectionMap {
			if duplicate > 96 {
				prioritySum += int(duplicate) - 96
			} else {
				prioritySum += int(duplicate) - 64 + 26
			}
		}
	}

	if isTest {
		if prioritySum != 157 {
			fmt.Println(prioritySum)
			panic("wrong sum")
		}
	}

	fmt.Println("first part solved")
	fmt.Println(prioritySum)
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	prioritySum := 0

	lineBuffer := make([]string, 0, 3)
	fmt.Println()

	for line := range lineChannel {
		if len(lineBuffer) < 2 {
			lineBuffer = append(lineBuffer, line)
			continue
		}
		if len(lineBuffer) < 3 {
			lineBuffer = append(lineBuffer, line)
		}

		firstRucksack := []rune(lineBuffer[0])
		secondRucksack := []rune(lineBuffer[1])
		thirdRucksack := []rune(lineBuffer[2])

		firstRucksackMap := make(map[rune]bool)
		secondRucksackMap := make(map[rune]bool)
		runeIntersectionMap := make(map[rune]bool)

		for _, item := range firstRucksack {
			firstRucksackMap[item] = true
		}

		for _, item := range secondRucksack {
			if firstRucksackMap[item] {
				secondRucksackMap[item] = true
			}
		}

		for _, item := range thirdRucksack {
			if secondRucksackMap[item] {
				runeIntersectionMap[item] = true
			}
		}

		for duplicate := range runeIntersectionMap {
			if duplicate > 96 {
				prioritySum += int(duplicate) - 96
			} else {
				prioritySum += int(duplicate) - 64 + 26
			}
		}

		lineBuffer = lineBuffer[:0]
	}

	if isTest {
		if prioritySum != 70 {
			fmt.Println(prioritySum)
			panic("wrong sum")
		}
	}

	fmt.Println("second part solved")
	fmt.Println(prioritySum)
}

func main() {
	isTest := false
	lineChannel := make(chan string)
	if isTest {
		go util.ReadInput("/home/markus/dev/godventofcode/2022/03/test.txt", lineChannel)
	} else {
		go util.ReadInput("/home/markus/dev/godventofcode/2022/03/input.txt", lineChannel)
	}

	solveFirstPart(lineChannel, isTest)
	// solveSecondPart(lineChannel, isTest)
}
