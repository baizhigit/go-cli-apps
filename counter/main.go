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
	wordCount := 0

	for _, v := range data {
		if v == ' ' {
			wordCount++
		}
	}

	wordCount++

	return wordCount
}
