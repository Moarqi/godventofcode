package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Moarqi/godventofcode/util"
)

func main() {
	lineChannel := make(chan string)
	go util.ReadInput("/home/markus/dev/godventofcode/2022/01/input.txt", lineChannel)

	caloriesPerElf := make([]int, 0)
	currentCalories := 0

	for line := range lineChannel {
		if len(line) < 1 {
			caloriesPerElf = append(caloriesPerElf, currentCalories)
			currentCalories = 0

			continue
		}

		calories, err := strconv.Atoi(line)

		if err != nil {
			panic(err)
		}

		currentCalories += calories
	}

	if currentCalories > 0 {
		caloriesPerElf = append(caloriesPerElf, currentCalories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))

	fmt.Println(caloriesPerElf[0])
	fmt.Println(caloriesPerElf[1])
	fmt.Println(caloriesPerElf[2])

	fmt.Println(caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2])
}
