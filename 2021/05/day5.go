package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	UNKNOWN Direction = iota
	NORTH
	EAST
	SOUTH
	WEST
	NE
	SE
	SW
	NW
)

type Vec2 struct {
	x int
	y int
}

type Line struct {
	from      Vec2
	to        Vec2
	direction Direction
}

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

func convertStringCoordinatesToVec2(coords string) Vec2 {
	elements := strings.Split(coords, ",")
	if len(elements) > 2 {
		panic("too many values in coordinate representation!")
	}

	intCoords := make([]int, 0)

	for _, el := range elements {
		intValue, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}

		intCoords = append(intCoords, intValue)
	}

	return Vec2{
		intCoords[0],
		intCoords[1],
	}
}

func convertCoordinateStrings(input []string) Line {
	from := convertStringCoordinatesToVec2(input[0])
	to := convertStringCoordinatesToVec2(input[1])

	line := Line{
		from,
		to,
		UNKNOWN,
	}

	if from.x == to.x {
		if from.y < to.y {
			line.direction = SOUTH
		} else if from.y > to.y {
			line.direction = NORTH
		}
	} else if from.y == to.y {
		if from.x < to.x {
			line.direction = EAST
		} else if from.x > to.x {
			line.direction = WEST
		}
	} else {
		if from.x < to.x {
			if from.y < to.y {
				line.direction = SE
			} else if from.y > to.y {
				line.direction = NE
			}
		} else if from.x > to.x {
			if from.y < to.y {
				line.direction = SW
			} else if from.y > to.y {
				line.direction = NW
			}
		}
	}

	if line.direction == UNKNOWN {
		fmt.Println(from)
		fmt.Println(to)
		panic("unknown direction!")
	}
	return line
}

func handleInputLine(inputLine string, grid *[][]int, useDiagonals bool) {
	rawCoordinates := strings.Split(inputLine, " -> ")
	if len(rawCoordinates) != 2 {
		fmt.Println(rawCoordinates)
		panic("coordinate split failed")
	}

	line := convertCoordinateStrings(rawCoordinates)

	if !useDiagonals &&
		line.direction != NORTH &&
		line.direction != EAST &&
		line.direction != SOUTH &&
		line.direction != WEST {
		return
	}

	vecPointer := Vec2{
		line.from.x,
		line.from.y,
	}

	for {
		for len(*grid)-1 < vecPointer.y {
			*grid = append(*grid, make([]int, 0))
		}
		for len((*grid)[vecPointer.y])-1 < vecPointer.x {
			(*grid)[vecPointer.y] = append((*grid)[vecPointer.y], 0)
		}

		(*grid)[vecPointer.y][vecPointer.x] += 1

		if vecPointer.x == line.to.x && vecPointer.y == line.to.y {
			break
		}

		incrementPointer(&vecPointer, line.direction)
	}
}

func incrementPointer(vec *Vec2, direction Direction) {
	if direction == NORTH {
		vec.y -= 1
	} else if direction == SOUTH {
		vec.y += 1
	} else if direction == EAST {
		vec.x += 1
	} else if direction == WEST {
		vec.x -= 1
	} else if direction == NE {
		vec.x += 1
		vec.y -= 1
	} else if direction == SE {
		vec.x += 1
		vec.y += 1
	} else if direction == SW {
		vec.x -= 1
		vec.y += 1
	} else if direction == NW {
		vec.x -= 1
		vec.y -= 1
	} else {
		panic("increment failed, direciton unhandled!")
	}
}

func printGrid(grid *[][]int) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			if (*grid)[i][j] > 0 {
				fmt.Print((*grid)[i][j])
			} else {
				fmt.Print(",")
			}
		}
		fmt.Print("\n")
	}
}

func countGridIntersections(grid *[][]int) int {
	intersections := 0
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			if (*grid)[i][j] > 1 {
				intersections += 1
			}
		}
	}

	return intersections
}

func main() {
	lineChannel := make(chan string)
	go readInput("/home/markus/dev/godventofcode/05/data/input", lineChannel)

	grid := make([][]int, 0)
	for line := range lineChannel {
		handleInputLine(line, &grid, true)
	}

	// printGrid(&grid)
	intersections := countGridIntersections(&grid)
	fmt.Println(intersections)
}
