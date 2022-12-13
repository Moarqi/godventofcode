package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/Moarqi/godventofcode/util"
)

// key is joined path
// parent is key without dir name
// size is sum of files
type Dir struct {
	key    string
	parent string
	size   int
}

func getDirMap(lineChannel chan string) map[string]Dir {
	dirMap := make(map[string]Dir)
	rootNode := Dir{
		key:    "/",
		parent: "",
		size:   0,
	}
	dirMap[rootNode.key] = rootNode
	currentKey := rootNode.key

	lsMode := false

	for line := range lineChannel {
		argSplit := strings.Split(line, " ")

		if strings.ContainsRune(line, '$') {
			lsMode = false
			if argSplit[1] == "cd" {
				if argSplit[2] == ".." {
					currentKey = dirMap[currentKey].parent
					continue
				}
				if argSplit[2] == "/" {
					currentKey = rootNode.key
					continue
				}

				newNode := Dir{
					key:    currentKey + argSplit[len(argSplit)-1] + "/",
					parent: currentKey,
					size:   0,
				}
				dirMap[newNode.key] = newNode
				currentKey = newNode.key
			} else if argSplit[1] == "ls" {
				fmt.Println("turn on ls mode")
				lsMode = true
				continue
			}
		} else if lsMode {
			fmt.Println(argSplit)
			if argSplit[0] == "dir" {
				continue
			}

			fileSize, err := strconv.Atoi(argSplit[0])

			if err != nil {
				panic(err)
			}

			node := dirMap[currentKey]
			// TODO: check if already seen

			for true {
				fmt.Println("add file " + argSplit[1] + " to " + node.key)
				node.size += fileSize
				dirMap[node.key] = node

				if len(node.parent) < 1 {
					break
				}

				node = dirMap[node.parent]
			}
		}
	}

	return dirMap
}

func solveFirstPart(lineChannel chan string, isTest bool) {
	dirSizeSum := 0

	dirMap := getDirMap(lineChannel)

	for _, value := range dirMap {
		if value.size <= 100000 {
			dirSizeSum += value.size
		}
	}

	if isTest {
		if dirSizeSum != 95437 {
			fmt.Println(dirSizeSum)
			panic("wrong start index")
		}
	}

	fmt.Println(dirSizeSum)
	fmt.Println("first part solved")
}

func solveSecondPart(lineChannel chan string, isTest bool) {
	deletedSpace := 0
	totalSpace := 70000000
	requiredSpace := 30000000

	dirMap := getDirMap(lineChannel)
	freeSpace := totalSpace - dirMap["/"].size
	requiredSpace -= freeSpace

	possibleCanidates := make([]int, 0)
	for _, value := range dirMap {
		if value.size > requiredSpace {
			possibleCanidates = append(possibleCanidates, value.size)
		}
	}

	sort.Ints(possibleCanidates)
	fmt.Println(possibleCanidates)
	deletedSpace = possibleCanidates[0]

	if isTest {
		if deletedSpace != 24933642 {
			fmt.Println(deletedSpace)
			panic("wrong start index")
		}
	}

	fmt.Println(deletedSpace)
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
