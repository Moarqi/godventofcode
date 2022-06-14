package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func handleLine(line string, bits *[]int, nBits int) int {
	stringBits := strings.Split(line, "")

	if len(stringBits) != nBits {
		panic("bit length mismatch!")
	}

	for i := 0; i < len(*bits); i++ {
		value, err := strconv.Atoi(stringBits[i])
		if err != nil {
			panic(err)
		}

		if value > 0 {
			(*bits)[i] += 1
		}
	}

	value, err := strconv.ParseInt(line, 2, 16)

	if err != nil {
		panic(err)
	}

	return int(value)
}

func getPowerConsumption(nLines int, onesPerBit []int, nBits int) int {
	commonThreshold := nLines / 2

	binaryString := ""

	for i := 0; i < len(onesPerBit); i++ {
		if onesPerBit[i] >= commonThreshold {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
	}

	intValue, err := strconv.ParseUint(binaryString, 2, nBits)

	if err != nil {
		panic(err)
	}

	flippedValue := ^intValue
	flippedBinary := strconv.FormatUint(flippedValue, 2)
	flippedBinary = flippedBinary[len(flippedBinary)-nBits:]

	flippedShortValue, err := strconv.ParseUint(flippedBinary, 2, nBits)

	if err != nil {
		panic(err)
	}

	fmt.Println(binaryString, intValue, flippedBinary, flippedShortValue)

	return int(intValue * flippedShortValue)
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func getRatingIndex(nLines int, onesPerBit []int, nBits int, _integerValues []int, dominant bool) int {
	integerValues := make([]int, len(_integerValues))
	copy(integerValues, _integerValues)

	availableIndices := make([]int, len(integerValues))
	for i := 0; i < len(integerValues); i++ {
		availableIndices[i] = i
	}

	for len(availableIndices) > 1 {
		for i := 0; i < nBits; i++ {
			nValues := len(availableIndices)
			if nValues == 1 {
				break
			}
			minBitValue := int(math.Pow(2, float64(nBits-1-i)))
			nLowerThanMinBitValue := 0

			for j := 0; j < len(availableIndices); j++ {
				if integerValues[availableIndices[j]] < minBitValue {
					nLowerThanMinBitValue += 1
				}
			}

			nValuesHalf := nValues / 2

			condition := nLowerThanMinBitValue > nValuesHalf
			if !dominant {
				condition = !condition
			}
			// zero dominant
			if condition {
				for k := 0; k < len(availableIndices); k++ {
					if integerValues[availableIndices[k]] >= minBitValue {
						availableIndices = RemoveIndex(availableIndices, k)
						k--
					}
				}
			} else {
				for k := 0; k < len(availableIndices); k++ {
					if integerValues[availableIndices[k]] < minBitValue {
						availableIndices = RemoveIndex(availableIndices, k)
						k--
					} else {
						integerValues[availableIndices[k]] -= minBitValue
					}
				}
			}

		}
	}

	return availableIndices[0]
}

func getLifeSupportRating(nLines int, onesPerBit []int, nBits int, integerValues []int, dominant bool) int {
	integerValuesOriginal := make([]int, len(integerValues))
	copy(integerValuesOriginal, integerValues)

	oxygenGeneratorRatingIdx := getRatingIndex(nLines, onesPerBit, nBits, integerValues, true)
	co2ScrubberRatingIdx := getRatingIndex(nLines, onesPerBit, nBits, integerValues, false)

	return integerValuesOriginal[oxygenGeneratorRatingIdx] * integerValuesOriginal[co2ScrubberRatingIdx]
}

func main() {
	nBits := 12
	lineChannel := make(chan string)
	go readInput("/home/markus/dev/godventofcode/03/data/input", lineChannel)

	nLines := 0

	onesPerBit := make([]int, nBits)
	integerValues := make([]int, 0, 1024)

	for line := range lineChannel {
		nLines += 1
		intValue := handleLine(line, &onesPerBit, nBits)
		integerValues = append(integerValues, intValue)
	}

	powerConsumption := getPowerConsumption(nLines, onesPerBit, nBits)
	fmt.Println(powerConsumption)
	lifeSupportRate := getLifeSupportRating(nLines, onesPerBit, nBits, integerValues, true)
	fmt.Println(lifeSupportRate)
}
