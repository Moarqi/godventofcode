package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadInput(filePath string, output chan string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	defer close(output)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		output <- strings.TrimSpace(scanner.Text())
		if err != nil {
			panic(err)
		}
	}
}
