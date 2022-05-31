package main

import (
	"bufio"
	"fmt"
	"log"
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

func handleLine(line string, X *int, Y *int) {
	splits := strings.Split(line, " ")

	cmd := splits[0]
	arg, err := strconv.Atoi(splits[1])
	if err != nil {
		panic(err)
	}

	switch cmd {
	case "forward":
		*X += arg
	case "up":
		*Y += arg
	case "down":
		*Y -= arg
	}
}

func handleLineAim(line string, X *int, Y *int, Aim *int) {
	splits := strings.Split(line, " ")

	cmd := splits[0]
	arg, err := strconv.Atoi(splits[1])
	if err != nil {
		panic(err)
	}

	switch cmd {
	case "forward":
		*X += arg
		*Y -= *Aim * arg
	case "up":
		*Aim -= arg
	case "down":
		*Aim += arg
	}
}

func main() {
	lineChannel := make(chan string)
	go readInput("./02/data/input", lineChannel)

	X := 0
	Y := 0
	Aim := 0

	for line := range lineChannel {
		handleLineAim(line, &X, &Y, &Aim)
	}

	fmt.Println(X, Y)

	if Y < 0 {
		Y = -Y
	}

	fmt.Println(X * Y)
}
