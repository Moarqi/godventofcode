package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/Moarqi/godventofcode/util"
)

func parseTreeMap(lineChannel chan string) [][]int {
	treeMap := make([][]int, 0)

	rowIndex := 0
	for line := range lineChannel {
		lineSplit := strings.Split(line, "")
		treeMap = append(treeMap, make([]int, len(lineSplit)))

		for i, tree := range lineSplit {
			treeHeight, err := strconv.Atoi(tree)
			if err != nil {
				panic(err)
			}

			treeMap[rowIndex][i] = treeHeight
		}
		rowIndex += 1
	}

	return treeMap
}

func getMaxScenicScore(treeMap [][]int) int {
	leftMaxIdx := make([]int, 0)
	rightMaxIdx := make([]int, 0)

	botMaxIdx := make([][]int, len(treeMap[0]))
	topMaxIdx := make([][]int, len(treeMap[0]))

	leftScoreMatrix := make([][]int, len(treeMap))
	rightScoreMatrix := make([][]int, len(treeMap))
	topScoreMatrix := make([][]int, len(treeMap))
	botScoreMatrix := make([][]int, len(treeMap))

	for i := 0; i < len(treeMap); i += 1 {
		leftScoreMatrix[i] = make([]int, len(treeMap[i]))
		topScoreMatrix[i] = make([]int, len(treeMap[i]))
		leftMaxIdx = leftMaxIdx[:0]

		if i == 0 {
			continue
		}

		for j := 1; j < len(treeMap[i]); j += 1 {
			value := treeMap[i][j]
			leftValue := treeMap[i][j-1]
			topValue := treeMap[i-1][j]

			if value > leftValue {
				leftMostIdx := -1
				for _j := len(leftMaxIdx) - 1; _j >= 0; _j -= 1 {
					idx := leftMaxIdx[_j]
					leftMax := treeMap[i][idx]

					if value > leftMax {
						leftMostIdx = idx
					} else {
						break
					}
				}

				if leftMostIdx >= 0 {
					leftScoreMatrix[i][j] += leftScoreMatrix[i][leftMostIdx]
					leftScoreMatrix[i][j] += j - leftMostIdx
				} else {
					leftScoreMatrix[i][j] = leftScoreMatrix[i][j-1] + 1
				}
			} else {
				leftScoreMatrix[i][j] = 1
				leftMaxIdx = append(leftMaxIdx, j-1)
			}

			if value > topValue {
				topMostIdx := -1
				for _i := len(topMaxIdx[j]) - 1; _i >= 0; _i -= 1 {
					idx := topMaxIdx[j][_i]
					topMax := treeMap[idx][j]
					if value > topMax {
						topMostIdx = idx
					} else {
						break
					}
				}
				if topMostIdx >= 0 {
					topScoreMatrix[i][j] += topScoreMatrix[topMostIdx][j]
					topScoreMatrix[i][j] += i - topMostIdx
				} else {
					topScoreMatrix[i][j] = topScoreMatrix[i-1][j] + 1
				}
			} else {
				topScoreMatrix[i][j] = 1
				topMaxIdx[j] = append(topMaxIdx[j], i-1)
			}
		}
	}

	for i := len(treeMap) - 1; i >= 0; i -= 1 {
		rightScoreMatrix[i] = make([]int, len(treeMap[i]))
		botScoreMatrix[i] = make([]int, len(treeMap[i]))
		rightMaxIdx = rightMaxIdx[:0]
		if i == len(treeMap)-1 {
			continue
		}

		for j := len(treeMap[i]) - 2; j >= 0; j -= 1 {
			value := treeMap[i][j]
			rightValue := treeMap[i][j+1]
			botValue := treeMap[i+1][j]

			if value > rightValue {
				rightMostIdx := -1

				for _j := len(rightMaxIdx) - 1; _j >= 0; _j -= 1 {
					idx := rightMaxIdx[_j]
					rightMax := treeMap[i][idx]

					if value > rightMax {
						rightMostIdx = idx
					} else {
						break
					}
				}

				if rightMostIdx >= 0 {
					rightScoreMatrix[i][j] += rightScoreMatrix[i][rightMostIdx]
					rightScoreMatrix[i][j] += rightMostIdx - j
				} else {
					rightScoreMatrix[i][j] = rightScoreMatrix[i][j+1] + 1
				}
			} else {
				rightScoreMatrix[i][j] = 1
				rightMaxIdx = append(rightMaxIdx, j+1)
			}

			if value > botValue {
				botMostIdx := -1

				for _i := len(botMaxIdx[j]) - 1; _i >= 0; _i -= 1 {
					idx := botMaxIdx[j][_i]
					botMax := treeMap[idx][j]

					if value > botMax {
						botMostIdx = idx
					} else {
						break
					}
				}

				if botMostIdx >= 0 {
					botScoreMatrix[i][j] += botScoreMatrix[botMostIdx][j]
					botScoreMatrix[i][j] += botMostIdx - i
				} else {
					botScoreMatrix[i][j] = botScoreMatrix[i+1][j] + 1
				}
			} else {
				botScoreMatrix[i][j] = 1
				botMaxIdx[j] = append(botMaxIdx[j], i+1)
			}
		}
	}

	maxValue := 0
	for i := 1; i < len(treeMap)-1; i += 1 {
		for j := 1; j < len(treeMap[i])-1; j += 1 {
			value := leftScoreMatrix[i][j] * topScoreMatrix[i][j] * rightScoreMatrix[i][j] * botScoreMatrix[i][j]
			if value > maxValue {
				// fmt.Println(value, i, j)
				maxValue = value
			}
		}
	}

	return maxValue
}

func getMaxScenicScoreNaive(treeMap [][]int) int {
	maxScenicScore := 0

	for i := 0; i < len(treeMap); i += 1 {
		for j := 0; j < len(treeMap[i]); j += 1 {
			topScore, leftScore, rightScore, botScore, scenicScore := 0, 0, 0, 0, 0
			treeValue := treeMap[i][j]
			for _i := i + 1; _i < len(treeMap); _i += 1 {
				botScore += 1
				if treeMap[_i][j] >= treeValue {
					break
				}
			}

			for _i := i - 1; _i >= 0; _i -= 1 {
				topScore += 1
				if treeMap[_i][j] >= treeValue {
					break
				}
			}

			for _j := j + 1; _j < len(treeMap[i]); _j += 1 {
				rightScore += 1
				if treeMap[i][_j] >= treeValue {
					break
				}
			}

			for _j := j - 1; _j >= 0; _j -= 1 {
				leftScore += 1
				if treeMap[i][_j] >= treeValue {
					break
				}
			}

			scenicScore = topScore * botScore * rightScore * leftScore

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore
}

func getVisibleTrees(treeMap [][]int) int {
	visibleTrees := 2*len(treeMap) + 2*len(treeMap[0]) - 4
	visibleMap := make(map[string]bool)
	colTopMax := make([]int, len(treeMap[0]))
	copy(colTopMax, treeMap[0])
	colBotMax := make([]int, len(treeMap[len(treeMap)-1]))
	copy(colBotMax, treeMap[len(treeMap)-1])

	for i := 1; i < len(treeMap)-1; i += 1 {
		leftMax := treeMap[i][0]
		rightMax := treeMap[i][len(treeMap[i])-1]

		for j := 1; j < len(treeMap[i])-1; j += 1 {
			index := fmt.Sprintf("%d,%d", i, j)
			value := treeMap[i][j]

			if value > colTopMax[j] {
				colTopMax[j] = value
				fmt.Printf("%d at (%d, %d) is visible from the top\n", value, i, j)

				visibleMap[index] = true
			}

			if value > leftMax {
				leftMax = value
				fmt.Printf("%d at (%d, %d) is visible from the left\n", value, i, j)

				visibleMap[index] = true
			}
		}

		for j := len(treeMap[i]) - 2; j > 0; j -= 1 {
			index := fmt.Sprintf("%d,%d", i, j)

			// i can break here cause i have seen the trees up until here from left
			// fuck you no, i might only have seen them from above
			// if visibleMap[index] {
			// 	break
			// }

			if treeMap[i][j] > rightMax {
				rightMax = treeMap[i][j]
				fmt.Printf("%d at (%d, %d) is visible from the right\n", treeMap[i][j], i, j)

				visibleMap[index] = true
			}

		}
	}

	for i := len(treeMap) - 2; i > 0; i -= 1 {
		for j := 1; j < len(treeMap[i])-1; j += 1 {
			index := fmt.Sprintf("%d,%d", i, j)
			value := treeMap[i][j]

			// if visibleMap[index] {
			// 	continue
			// }

			if value > colBotMax[j] {
				fmt.Println(colBotMax)
				colBotMax[j] = value
				fmt.Printf("%d at (%d, %d) is visible from the bot\n", value, i, j)

				visibleMap[index] = true
				// continue
			}
		}
	}

	return visibleTrees + len(visibleMap)
}

func printMap(treeMap [][]int) {
	for _, row := range treeMap {
		fmt.Println(row)
	}
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	treeMap := parseTreeMap(lineChannel)
	// treeMap = treeMap[:4]

	visibleTreeCount := getVisibleTrees(treeMap)
	printMap(treeMap)

	if isTest {
		if visibleTreeCount != 21 {
			fmt.Println(visibleTreeCount)
			panic("wrong tree count")
		}
	}

	fmt.Println(visibleTreeCount)
	fmt.Println("first part solved")
}

var treeMap [][]int

func test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMaxScenicScore(treeMap)
	}
}

func testNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMaxScenicScoreNaive(treeMap)
	}
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	treeMap = parseTreeMap(lineChannel)
	// treeMap = treeMap[:4]

	fmt.Println(testing.Benchmark(testNaive))
	fmt.Println(testing.Benchmark(test))

	// fmt.Println(scenicScoreAlt)

	// if isTest {
	// 	if scenicScore != 8 {
	// 		fmt.Println(scenicScore)
	// 		panic("wrong scenic score")
	// 	}
	// }

	// fmt.Println(scenicScore)
	// fmt.Println("second part solved")
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
