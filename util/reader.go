package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadInput(filePath string, output chan string, trimWhitespace bool) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	defer close(output)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		if trimWhitespace {
			text = strings.TrimSpace(text)
		}

		output <- text
		if err != nil {
			panic(err)
		}
	}
}
