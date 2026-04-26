package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := countWords(data)

	fmt.Println(wordCount)
}

func countWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	wordDetected := false
	wordCount := 0

	for _, v := range data {
		if v == ' ' {
			wordCount++
		} else {
			wordDetected = true
		}
	}

	if !wordDetected {
		return 0
	}

	wordCount++

	return wordCount
}
