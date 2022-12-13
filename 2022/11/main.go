package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/Moarqi/godventofcode/util"
)

type Monkey struct {
	id               int
	items            []int
	update           func(int) int
	test             func(int) bool
	destinationTrue  int
	destinationFalse int
	inspections      int
}

func parseMonkeys(lineChannel chan string) []Monkey {
	monkey := Monkey{inspections: 0}
	monkeys := make([]Monkey, 0)
	factors := make([]int, 0)

	for line := range lineChannel {
		line := strings.TrimSpace(line)
		lineSplit := strings.Split(line, " ")
		if strings.HasPrefix(line, "M") {
			id, err := strconv.Atoi(strings.Replace(lineSplit[1], ":", "", -1))
			if err != nil {
				panic(err)
			}

			monkey.id = id
		} else if strings.HasPrefix(line, "S") {
			line = strings.Replace(line, "Starting items: ", "", -1)
			itemSplit := strings.Split(line, ", ")
			for _, itemString := range itemSplit {
				item, err := strconv.Atoi(itemString)
				if err != nil {
					panic(err)
				}

				monkey.items = append(monkey.items, item)
			}
		} else if strings.HasPrefix(line, "O") {
			argString := lineSplit[len(lineSplit)-1]
			operation := lineSplit[len(lineSplit)-2]
			if argString == "old" {
				if operation == "+" {
					monkey.update = func(x int) int { return x + x }
				} else if operation == "*" {
					monkey.update = func(x int) int { return x * x }
				} else {
					panic("unknown update")
				}

				continue
			}

			argInt, err := strconv.Atoi(argString)
			if err != nil {
				panic(err)
			}

			if operation == "+" {
				monkey.update = func(x int) int { return x + argInt }
			} else if operation == "*" {
				monkey.update = func(x int) int { return x * argInt }
			} else {
				panic("unknown update")
			}
		} else if strings.HasPrefix(line, "T") {
			argString := lineSplit[len(lineSplit)-1]
			operation := lineSplit[len(lineSplit)-3]

			argInt, err := strconv.Atoi(argString)
			if err != nil {
				panic(err)
			}

			if operation != "divisible" {
				panic("unknown test function")
			}

			monkey.test = func(x int) bool {
				return x%argInt == 0
			}
			factors = append(factors, argInt)
		} else if strings.HasPrefix(line, "If true") {
			destinationString := lineSplit[len(lineSplit)-1]

			destination, err := strconv.Atoi(destinationString)
			if err != nil {
				panic(err)
			}

			monkey.destinationTrue = destination
		} else if strings.HasPrefix(line, "If false") {
			destinationString := lineSplit[len(lineSplit)-1]

			destination, err := strconv.Atoi(destinationString)
			if err != nil {
				panic(err)
			}

			monkey.destinationFalse = destination
		} else {
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
		}
	}

	monkeys = append(monkeys, monkey)
	fmt.Println(factors)
	fac := 1
	for _, factor := range factors {
		fac *= factor
	}
	fmt.Println(fac)

	return monkeys
}

func simulateRound(monkeys *[]Monkey, divideWorryLevel bool) {
	for id, monkey := range *monkeys {
		for i := range monkey.items {
			level := monkey.update(monkey.items[i])
			if divideWorryLevel {
				level = int(math.Floor(float64(level) / 3))
			} else {
				level = level % 9699690
			}

			var destination int
			if monkey.test(level) {
				destination = monkey.destinationTrue
			} else {
				destination = monkey.destinationFalse
			}

			monkey.inspections += 1
			(*monkeys)[destination].items = append((*monkeys)[destination].items, level)
		}

		monkey.items = monkey.items[:0]
		(*monkeys)[id] = monkey
	}
}

func multiplyMostActiveInspections(monkeys []Monkey) int {
	inspectionCount := make([]int, 0)
	for _, monkey := range monkeys {
		inspectionCount = append(inspectionCount, monkey.inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCount)))

	return inspectionCount[0] * inspectionCount[1]
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	monkeys := parseMonkeys(lineChannel)
	for r := 0; r < 20; r += 1 {
		simulateRound(&monkeys, true)
	}

	monkeyBusiness := multiplyMostActiveInspections(monkeys)
	fmt.Println(monkeys)

	if isTest {
		if monkeyBusiness != 10605 {
			panic("wrong business level")
		}
	}

	fmt.Println(monkeyBusiness)
	fmt.Println("first part solved")
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	monkeys := parseMonkeys(lineChannel)
	for r := 0; r < 10000; r += 1 {
		simulateRound(&monkeys, false)
	}

	monkeyBusiness := multiplyMostActiveInspections(monkeys)
	fmt.Println(monkeys)

	if isTest {
		if monkeyBusiness != 2713310158 {
			panic("wrong business level")
		}
	}

	fmt.Println(monkeyBusiness)

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
