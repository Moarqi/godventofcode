package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// func main() {
// 	body, err := iotuil.ReadFile("file.txt")
// 	if err != nil {
// 		log.Fatalf("unable to read file: %v", err)
// 	}
// 	fmt.Println(string(body))
// }

func readInput(filePath string) []int {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	inputs := make([]int, 0, 1024)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		strValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		inputs = append(inputs, strValue)
	}

	return inputs
}

func countIncreases(input []int) int {
	count := 0

	prevValue := input[0]

	for i := 0; i < len(input); i++ {
		if input[i]-prevValue > 0 {
			count += 1
		}

		prevValue = input[i]
	}

	return count
}

func countSlidingWindowIncreases(input []int) int {
	count := 0

	prevWindowSum := input[0] + input[1] + input[2]

	for i := 2; i < len(input); i++ {
		windowSum := input[i] + input[i-1] + input[i-2]
		if windowSum-prevWindowSum > 0 {
			count += 1
		}

		prevWindowSum = windowSum
	}

	return count
}

func main() {
	inputValues := readInput("data/day1")

	fmt.Println(countIncreases(inputValues))
	fmt.Println(countSlidingWindowIncreases(inputValues))
}
