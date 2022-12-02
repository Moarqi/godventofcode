package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getStartingSimulation(inputChannel chan string) []int {
	startingPopulation := make([]int, 9)

	for line := range inputChannel {
		splitStrings := strings.Split(line, ",")

		for _, popAsString := range splitStrings {
			pop, err := strconv.Atoi(popAsString)
			if err != nil {
				panic(err)
			}

			startingPopulation[pop] += 1
		}
	}

	return startingPopulation
}

func startSimulation(population []int, epochs int) int {
	for i := 0; i < epochs; i++ {
		newEntities := population[0]
		population = population[1:]
		population = append(population, newEntities)
		population[6] += newEntities
	}

	totalPopulation := 0
	for _, e := range population {
		totalPopulation += e
	}

	return totalPopulation
}

func main() {
	epochs := 256

	start := time.Now()

	lineChannel := make(chan string)
	go readInput("/home/markus/dev/godventofcode/06/data/input", lineChannel)

	startingPopulation := getStartingSimulation(lineChannel)
	finalCount := startSimulation(startingPopulation, epochs)

	fmt.Println(finalCount)

	elapsed := time.Since(start)
	fmt.Printf("simulation took %s\n", elapsed)
}
