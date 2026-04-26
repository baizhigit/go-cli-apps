package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	_ = data

	wordCount := 0

	for _, v := range data {
		if v == ' ' {
			wordCount++
		}
	}

	wordCount++

	fmt.Println(wordCount)
}
