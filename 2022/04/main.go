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

type Section struct {
	min, max int
}

func splitSections(sectionsString string) (Section, Section) {
	sectionSplit := strings.Split(sectionsString, ",")

	sections := make([]Section, 0)

	for _, section := range sectionSplit {
		sectionRange := strings.Split(section, "-")
		minSection, err := strconv.Atoi(sectionRange[0])
		if err != nil {
			panic(err)
		}
		maxSection, err := strconv.Atoi(sectionRange[1])
		if err != nil {
			panic(err)
		}

		sections = append(sections, Section{min: minSection, max: maxSection})
	}

	return sections[0], sections[1]
}

func isFullSectionOverlap(sectionA Section, sectionB Section) bool {
	if sectionA.min <= sectionB.min && sectionA.max >= sectionB.max {
		return true
	}

	if sectionB.min <= sectionA.min && sectionB.max >= sectionA.max {
		return true
	}

	return false
}

func isPartlySectionOverlap(sectionA Section, sectionB Section) bool {
	if sectionA.max < sectionB.min || sectionB.max < sectionA.min {
		return false
	}

	return true
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	fullOverlaps := 0

	for line := range lineChannel {
		sectionA, sectionB := splitSections(line)
		if isFullSectionOverlap(sectionA, sectionB) {
			fullOverlaps += 1
		}
	}

	if isTest {
		if fullOverlaps != 2 {
			fmt.Println(fullOverlaps)
			panic("wrong count")
		}
	}

	fmt.Println(fullOverlaps)
	fmt.Println("first part solved")
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	partlyOverlaps := 0

	for line := range lineChannel {
		sectionA, sectionB := splitSections(line)
		if isPartlySectionOverlap(sectionA, sectionB) {
			partlyOverlaps += 1
		}
	}

	if isTest {
		if partlyOverlaps != 4 {
			fmt.Println(partlyOverlaps)
			panic("wrong count")
		}
	}

	fmt.Println(partlyOverlaps)
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
		go util.ReadInput(path.Dir(filename)+"/test.txt", lineChannel)
	} else {
		go util.ReadInput(path.Dir(filename)+"/input.txt", lineChannel)
	}

	if len(os.Args) > 1 && os.Args[1] == "1" {
		solveFirstPart(lineChannel, isTest)
	} else if len(os.Args) > 1 && os.Args[1] == "2" {
		solveSecondPart(lineChannel, isTest)
	} else {
		fmt.Println("part not specified")
	}
}
