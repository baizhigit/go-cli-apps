package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World counter")

	data, _ := os.ReadFile("./words.txt")
	fmt.Println(string(data))
}
