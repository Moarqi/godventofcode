package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
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

func getCounts(inputChannel chan string) map[int]int {
	counts := make(map[int]int)

	for line := range inputChannel {
		splitStrings := strings.Split(line, ",")

		for _, popAsString := range splitStrings {
			pos, err := strconv.Atoi(popAsString)
			if err != nil {
				panic(err)
			}

			counts[pos] += 1
		}
	}

	return counts
}

func fib(n int) int {
	if n == 0 {
		return 0
	}

	return n + fib(n-1)
}

func getFibCost(counts map[int]int, position int) int {
	totalCost := 0

	for k, v := range counts {
		if (k - position) > 0 {
			totalCost += v * fib(k-position)
		} else {
			totalCost += v * fib(-1*(k-position))
		}
	}
	return totalCost
}

func getCost(counts map[int]int, position int) int {
	totalCost := 0
	for k, v := range counts {
		if (k - position) > 0 {
			totalCost += v * (k - position)
		} else {
			totalCost += -1 * v * (k - position)
		}
	}
	return totalCost
}

func getMedian(keys []int) int {
	length := len(keys)
	if length == 1 {
		return 0
	}

	if length%2 == 0 {
		fmt.Println("warning, even number median!")
	}

	return (length + 1) / 2
}

func splitAtMedian(keys []int) ([]int, []int) {
	median := getMedian(keys)
	fmt.Printf("split at index %d\n", median)
	return keys[0 : median+1], keys[median:]
}

func splitAndCalcCost(counts map[int]int, keys []int) {
	if len(keys) == 0 {
		// fmt.Printf("%d seems to be it!\n", keys[0])
		return
	}
	if len(keys) == 1 {
		fmt.Printf("%d seems to be it!\n", keys[0])
		fmt.Printf("%d fuel cost", getFibCost(counts, keys[0]))
		return
	}

	left, right := splitAtMedian(keys)
	leftCost := getFibCost(counts, left[getMedian(left)])
	rightCost := getFibCost(counts, right[getMedian(right)])

	fmt.Println(leftCost)
	fmt.Println(rightCost)

	if leftCost < rightCost {
		splitAndCalcCost(counts, left)
	} else {
		splitAndCalcCost(counts, right)
	}
}

// TODO: split in middle, calculate cost in median, go to side with lower cost -> recursive
func getMinimum(counts map[int]int) {
	keys := make([]int, 0)

	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	maxVal := keys[len(keys)-1]

	keys = make([]int, maxVal+1)

	for i := 0; i <= maxVal; i++ {
		keys[i] = i
	}

	splitAndCalcCost(counts, keys)
}

func getMinBad(counts map[int]int) {
	keys := make([]int, 0)

	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	maxVal := keys[len(keys)-1]

	keys = make([]int, maxVal+1)

	for i := 0; i <= maxVal; i++ {
		keys[i] = i
	}

	minCost := math.MaxInt32
	minIndex := 0

	for k := range keys {
		cost := getFibCost(counts, k)
		if cost < minCost {
			minCost = cost
			minIndex = k
		}
	}

	fmt.Println(minIndex)
	fmt.Println(minCost)
}

func main() {
	start := time.Now()

	lineChannel := make(chan string)
	go readInput("/home/markus/dev/godventofcode/07/data/input", lineChannel)

	probs := getCounts(lineChannel)
	getMinimum(probs)

	// fmt.Println(finalCount)

	elapsed := time.Since(start)
	fmt.Printf("simulation took %s\n", elapsed)
}
