package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// TODO: move to utils
func readInput(filePath string, output chan string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	defer close(output)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		output <- scanner.Text()
		if err != nil {
			panic(err)
		}
	}
}

func getDecodeMap(inputPattern []string) map[string]string {
	sort.Slice(inputPattern, func(i, j int) bool {
		return len(inputPattern[i]) < len(inputPattern[j])
	})

	decoding := make(map[string]string)
	potentialDecodings := make(map[string][]string)

	for i, v := range inputPattern {
		inputLength := len(v)
		inputSplit := strings.Split(v, "")
		inputMap := make(map[string]bool)

		for _, v := range inputSplit {
			inputMap[v] = true
		}

		if inputLength == 2 { // one
			potentialDecodings["c"] = append(potentialDecodings["c"], inputSplit...)
			potentialDecodings["f"] = append(potentialDecodings["f"], inputSplit...)
		} else if inputLength == 3 { // seven
			for _, pv := range potentialDecodings["c"] {
				delete(inputMap, pv)
			}

			// todo this should be one valued!
			for k, _ := range inputMap {
				decoding["a"] = k
			}
		} else if inputLength == 4 { // four
			for _, pv := range potentialDecodings["c"] {
				delete(inputMap, pv)
			}

			potentialDecodings["b"] = append(potentialDecodings["b"], inputSplit...)
			potentialDecodings["d"] = append(potentialDecodings["d"], inputSplit...)
		} else if inputLength == 5 { // two, three and five
			// TODO: rather hardcode instead of elif and check all three potential inputs

			// simulatenously check next 2 inputs
			secondInput := inputPattern[i+1]
			thirdInput := inputPattern[i+2]

			delete(inputMap, decoding["a"])

			// if inputMap has potential c and f -> g, d and b can be resolved

			// if inputMap has potential b and d -> f can be resolved
			// use the remaining input to resolve
		}
	}
}

/// this returns the decoded outputs as elements in an array
func decodeOutput(outputStrings []string) []int {
	decoded := make([]int, 4)

	for i, v := range outputStrings {
		n := len(v)
		// for now, only count valid
		if n == 2 || n == 4 || n == 3 || n == 7 {
			decoded[i] = 1
		} else {
			decoded[i] = -1
		}
	}

	return decoded
}

func countUniqueDigits(inputLine string) int {
	validCount := 0

	splitStrings := strings.Split(inputLine, "|")

	// signal := splitStrings[0]
	output := splitStrings[1]

	// signals := strings.Split(signal, " ")
	outputs := strings.Split(strings.TrimSpace(output), " ")

	decodedOutput := decodeOutput(outputs)

	for i := range decodedOutput {
		if decodedOutput[i] > 0 {
			validCount += 1
		}
	}

	return validCount
}

func getUniqueOutputs(inputLines chan string) int {
	totalValid := 0
	for line := range inputLines {
		totalValid += countUniqueDigits(line)
	}

	return totalValid
}

func main() {
	start := time.Now()

	lineChannel := make(chan string)
	go readInput("/home/markus/dev/godventofcode/08/data/input", lineChannel)

	validCount := getUniqueOutputs(lineChannel)
	fmt.Printf("%d are valid\n", validCount)

	elapsed := time.Since(start)
	fmt.Printf("calculation took %s\n", elapsed)
}
