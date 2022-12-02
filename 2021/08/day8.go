package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/Moarqi/godventofcode/util"
)

var segmentMap = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func makeInputMap(inputSplit []string) map[string]bool {
	inputMap := make(map[string]bool)

	for _, v := range inputSplit {
		inputMap[v] = true
	}

	return inputMap
}

func getDecodeMap(inputPattern []string) map[string]string {
	sort.Slice(inputPattern, func(i, j int) bool {
		return len(inputPattern[i]) < len(inputPattern[j])
	})

	encoding := make(map[string]string)
	potentialDecodings := make(map[string][]string)

	for i, v := range inputPattern {
		inputLength := len(v)
		inputSplit := strings.Split(v, "")
		inputMap := makeInputMap(inputSplit)

		if inputLength == 2 { // one
			potentialDecodings["c"] = append(potentialDecodings["c"], inputSplit...)
			potentialDecodings["f"] = append(potentialDecodings["f"], inputSplit...)
		} else if inputLength == 3 { // seven
			for _, pv := range potentialDecodings["c"] {
				delete(inputMap, pv)
			}

			if len(inputMap) > 1 {
				panic("expected one entry for a decoding!")
			}

			for k := range inputMap {
				encoding["a"] = k
			}
		} else if inputLength == 4 { // four
			for _, pv := range potentialDecodings["c"] {
				delete(inputMap, pv)
			}

			potentailDB := make([]string, 0)
			for k := range inputMap {
				potentailDB = append(potentailDB, k)
			}

			potentialDecodings["b"] = append(potentialDecodings["b"], potentailDB...)
			potentialDecodings["d"] = append(potentialDecodings["d"], potentailDB...)
		} else if inputLength == 5 { // two, three and five
			// TODO: rather hardcode instead of elif and check all three potential inputs
			inputs := make([]map[string]bool, 3)
			inputs[0] = inputMap
			inputs[1] = makeInputMap(strings.Split(inputPattern[i+1], ""))
			inputs[2] = makeInputMap(strings.Split(inputPattern[i+2], ""))

			solved := false
			solvedFive := false
			solvedThree := false

			for solved == false {
				for j, input := range inputs {
					for _, decoded := range encoding {
						delete(inputs[j], decoded)
					}

					if len(input) == 0 {
						continue
					}

					// fmt.Println(input)

					if solvedFive && len(input) == 1 {
						// this must be two
						for k := range input {
							encoding["e"] = k
						}
						fmt.Println("solved two!")
						solved = true
						break
					} else if solvedFive {
						fmt.Println(inputs, input)
						fmt.Println(encoding)
						panic("five solved but input invalid!")
					}

					// if inputMap has potential b and d -> f can be resolved
					// if three was solved, the only remaining key is for f
					if len(input) == 1 {
						// this must be five
						for k := range input {
							encoding["f"] = k
						}

						if encoding["f"] == potentialDecodings["f"][0] {
							encoding["c"] = potentialDecodings["f"][1]
						} else {
							encoding["c"] = potentialDecodings["f"][0]
						}

						delete(potentialDecodings, "c")
						delete(potentialDecodings, "f")

						fmt.Println("solved five!")
						solvedFive = true
					}

					// if inputMap has potential c and f -> g, d and b can be resolved
					if !solvedThree && input[potentialDecodings["c"][0]] && input[potentialDecodings["c"][1]] {
						// this must be 3
						if input[potentialDecodings["d"][0]] {
							encoding["d"] = potentialDecodings["d"][0]
							encoding["b"] = potentialDecodings["d"][1]
						}

						if input[potentialDecodings["d"][1]] {
							encoding["d"] = potentialDecodings["d"][1]
							encoding["b"] = potentialDecodings["d"][0]
						}

						delete(potentialDecodings, "d")
						delete(potentialDecodings, "b")
						delete(input, encoding["b"])
						delete(input, encoding["d"])

						for k := range input {
							if k == potentialDecodings["c"][0] || k == potentialDecodings["c"][1] {
								continue
							}

							encoding["g"] = k
						}

						fmt.Println("solved three!")
						solvedThree = true
					}
				}
			}
			break
		} else {
			break
		}
	}

	// invert decoding
	decoding := make(map[string]string, 7)
	for k, v := range encoding {
		decoding[v] = k
	}

	return decoding
}

// / this returns the decoded outputs as elements in an array
func decodeOutput(outputStrings []string, decoding map[string]string) int {
	digits := len(outputStrings)

	outputValue := 0

	for i, enabledSegments := range outputStrings {
		enabledSegmentsDecoded := make([]string, 0)
		for _, segment := range enabledSegments {
			enabledSegmentsDecoded = append(enabledSegmentsDecoded, decoding[string(segment)])
		}

		digit := getDigit(enabledSegmentsDecoded)

		outputValue += digit * int(math.Pow10(digits-1-i))

		// n := len(enabledSegments)
		// // for now, only count valid
		// if n == 2 || n == 4 || n == 3 || n == 7 {
		// 	decoded[i] = 1
		// } else {
		// 	decoded[i] = -1
		// }
	}

	return outputValue
}

func getDigit(enabledSegmentsDecoded []string) int {
	sort.Strings(enabledSegmentsDecoded)
	return segmentMap[strings.Join(enabledSegmentsDecoded, "")]
}

// func countUniqueDigits(inputLine string) int {
// 	validCount := 0

// 	splitStrings := strings.Split(inputLine, "|")

// 	// signal := splitStrings[0]
// 	output := splitStrings[1]

// 	// signals := strings.Split(signal, " ")
// 	outputs := strings.Split(strings.TrimSpace(output), " ")

// 	decodedOutput := decodeOutput(outputs)

// 	for i := range decodedOutput {
// 		if decodedOutput[i] > 0 {
// 			validCount += 1
// 		}
// 	}

// 	return validCount
// }

// func getUniqueOutputs(inputLines chan string) int {
// 	totalValid := 0
// 	for line := range inputLines {
// 		totalValid += countUniqueDigits(line)
// 	}

// 	return totalValid
// }

func decodeSignal(inputLines chan string) {
	totalOutputValue := 0

	for line := range inputLines {
		splitStrings := strings.Split(line, "|")

		signal := splitStrings[0]
		output := splitStrings[1]

		signals := strings.Split(strings.TrimSpace(signal), " ")
		outputs := strings.Split(strings.TrimSpace(output), " ")
		_ = outputs

		decodeMap := getDecodeMap(signals)
		outputValue := decodeOutput(outputs, decodeMap)

		fmt.Println(decodeMap, outputValue)
		totalOutputValue += outputValue
	}

	fmt.Println(totalOutputValue)
}

func main() {
	start := time.Now()

	lineChannel := make(chan string)
	go util.ReadInput("/home/markus/dev/godventofcode/2021/08/data/input", lineChannel)

	decodeSignal(lineChannel)
	// validCount := getUniqueOutputs(lineChannel)
	// fmt.Printf("%d are valid\n", validCount)

	elapsed := time.Since(start)
	fmt.Printf("calculation took %s\n", elapsed)
}
