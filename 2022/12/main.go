package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/Moarqi/godventofcode/util"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Node struct {
	i, j, height, distance int
	prev                   *Node
}

func parseGrid(lineChannel chan string) ([][]*Node, *Node, *Node) {
	var start, goal *Node
	grid := make([][]*Node, 0)
	rowIdx := 0

	for line := range lineChannel {
		lineSplit := []rune(line)
		row := make([]*Node, len(lineSplit))
		if len(lineSplit) < 2 {
			break
		}

		for j, val := range lineSplit {
			row[j] = &Node{i: rowIdx, j: j, height: int(val - 'a'), distance: MaxInt, prev: nil}
			if val == 'S' {
				start = row[j]
				(*start).distance = 0
				(*start).height = 0
			} else if val == 'E' {
				goal = row[j]
				(*goal).height = int('z' - 'a')
			}
		}

		grid = append(grid, row)
		rowIdx += 1
	}

	return grid, start, goal
}

func initUnvisited(grid [][]*Node, start *Node) map[*Node]bool {
	unvisited := make(map[*Node]bool)

	for i := range grid {
		for j := range grid[i] {
			unvisited[grid[i][j]] = true
		}
	}

	return unvisited
}

func getNextUnvisited(unvisited map[*Node]bool) *Node {
	minDist := MaxInt
	var minDistNode *Node

	for node := range unvisited {
		if node.distance < minDist {
			minDist = node.distance
			minDistNode = node
		}
	}

	return minDistNode
}

func getReachableNeightbors(grid [][]*Node, node *Node) []*Node {
	neighbors := make([]*Node, 0)

	maxI := len(grid) - 1
	maxJ := len(grid[0]) - 1

	// left
	if node.i > 0 {
		neighbor := grid[node.i-1][node.j]
		if reachable(node, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	if node.j > 0 {
		neighbor := grid[node.i][node.j-1]
		if reachable(node, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	if node.i < maxI {
		neighbor := grid[node.i+1][node.j]
		if reachable(node, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	if node.j < maxJ {
		neighbor := grid[node.i][node.j+1]
		if reachable(node, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func reachable(node, neighbor *Node) bool {
	return node.height+1 >= neighbor.height
}

func solveGrid(grid [][]*Node, start, goal *Node) ([]Node, bool) {
	path := make([]Node, 0)

	visited := make(map[*Node]bool)
	unvisited := initUnvisited(grid, start)
	node := start // copy?

	for true {
		neighbors := getReachableNeightbors(grid, node)

		// if len(neighbors) < 1 {
		// 	return path, false
		// }

		for _, neighbor := range neighbors {
			if _, ok := visited[neighbor]; !ok {
				neighbor.distance = node.distance + 1
				neighbor.prev = node
			}
		}

		delete(unvisited, node)
		visited[node] = true
		if node == goal {
			break
		}

		node = getNextUnvisited(unvisited)
		if node == nil {
			return path, false
		}
	}

	for node != start {
		path = append(path, *node)
		node = node.prev
	}

	return path, true
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	grid, start, goal := parseGrid(lineChannel)
	path, _ := solveGrid(grid, start, goal)

	pathLength := len(path)

	for _, node := range path {
		fmt.Println(node.i, node.j, node.height)
	}

	if isTest {
		if pathLength != 31 {
			panic("wrong min path length")
		}
	}

	fmt.Println(pathLength)
	fmt.Println("first part solved")
}

func resetGrid(grid [][]*Node, start *Node) {
	for _, row := range grid {
		for _, node := range row {
			node.distance = MaxInt
		}
	}

	start.distance = 0
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	grid, _, goal := parseGrid(lineChannel)

	minPathLength := MaxInt

	for _, row := range grid {
		for _, nodePtr := range row {
			if nodePtr.height == 0 {
				resetGrid(grid, nodePtr)

				if nodePtr.height == 0 {
					path, ok := solveGrid(grid, nodePtr, goal)
					if ok && len(path) < minPathLength {
						minPathLength = len(path)
						fmt.Println(minPathLength)
					}
				}
			}
		}
	}

	if isTest {
		if minPathLength != 29 {
			panic("wrong min path length")
		}
	}

	fmt.Println(minPathLength)
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
