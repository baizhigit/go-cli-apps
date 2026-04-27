package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./words.txt")
	if err != nil {
		os.Exit(1)
	}

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)

	return len(words)
}
