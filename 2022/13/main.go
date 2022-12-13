package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"

	"github.com/Moarqi/godventofcode/util"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func parseList(runes []rune, initLevel int) []any {
	list := make([]any, 0)
	parseIntMode := false
	intStringBuffer := ""
	level := initLevel
	for i, r := range runes {
		if level > 1 && level > initLevel {
			if r == ']' {
				level -= 1
			}
			if r == '[' {
				level += 1
			}
			continue
		}
		if parseIntMode {
			if r == ',' {
				parseIntMode = false
				arg, err := strconv.Atoi(intStringBuffer)
				intStringBuffer = ""
				if err != nil {
					panic(err)
				}
				list = append(list, arg)
				continue
			} else if r == ']' {
				parseIntMode = false
				arg, err := strconv.Atoi(intStringBuffer)
				intStringBuffer = ""
				if err != nil {
					panic(err)
				}
				list = append(list, arg)
			} else {
				intStringBuffer += string(r)
				continue
			}
		}

		if r == '[' {
			if level > 0 {
				list = append(list, parseList(runes[i+1:], level+1))
			}
			level += 1
		} else if r == ']' {
			return list
		} else if r == ',' {
			continue
		} else {
			parseIntMode = true
			intStringBuffer += string(r)
		}
	}

	return list
}

func compare(left, right []any) (bool, bool) {
	if len(left) == 0 && len(right) == 0 {
		return false, false
	}
	if len(left) == 0 && len(right) > 0 {
		return true, true
	}
	if len(right) == 0 && len(left) > 0 {
		return false, true
	}

	leftItem := left[0]
	rightItem := right[0]

	rightArray, rightIsArray := rightItem.([]any)
	leftArray, leftIsArray := leftItem.([]any)
	rightInt, rightIsInt := rightItem.(int)
	leftInt, leftIsInt := leftItem.(int)

	if leftIsArray && rightIsArray {
		// fmt.Println("recurse call")
		if res, final := compare(leftArray, rightArray); final {
			return res, final
		}
	} else if leftIsArray {
		rightIntArray := []any{rightInt}
		if res, final := compare(leftArray, rightIntArray); final {
			return res, final
		}
	} else if rightIsArray {
		leftIntArray := []any{leftInt}
		if res, final := compare(leftIntArray, rightArray); final {
			return res, final
		}
	}

	if leftIsInt && rightIsInt {
		if leftInt < rightInt {
			return true, true
		} else if leftInt > rightInt {
			return false, true
		}
	}

	return compare(left[1:], right[1:])
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	readLeft := true
	var left, right []any

	i := 0
	correctOrderSum := 0

	for line := range lineChannel {
		if len(line) < 1 {
			continue
		}

		if readLeft {
			left = parseList([]rune(line), 0)
			readLeft = false
			continue
		}

		right = parseList([]rune(line), 0)
		readLeft = true
		i += 1

		if res, final := compare(left, right); final {
			if res {
				correctOrderSum += i
			}
		} else {
			panic("whaa")
		}
	}

	if isTest {
		if correctOrderSum != 13 {
			fmt.Println(correctOrderSum)
			panic("wrong indices sum")
		}
	}

	fmt.Println(correctOrderSum)
	fmt.Println("first part solved")
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	signals := make([]*[]any, 0)

	for line := range lineChannel {
		if len(line) < 1 {
			continue
		}
		signal := parseList([]rune(line), 0)
		signals = append(signals, &signal)
	}

	divider2 := []any{[]any{2}}
	divider6 := []any{[]any{6}}
	signals = append(signals, &divider2)
	signals = append(signals, &divider6)

	sort.Slice(signals, func(i, j int) bool {
		res, _ := compare(*signals[i], *signals[j])
		return res
	})

	decoderKey := 1
	for i, signal := range signals {
		if signal == &divider2 {
			decoderKey *= i + 1
		}
		if signal == &divider6 {
			decoderKey *= i + 1
		}
	}

	if isTest {
		if decoderKey != 140 {
			fmt.Println(decoderKey)
			panic("wrong decoder key")
		}
	}

	fmt.Println(decoderKey)
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
