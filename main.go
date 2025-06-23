package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	_ = data

	wordCount := 0

	for _, value := range data {
		if value == ' ' {
			wordCount++
		}
	}

	fmt.Println("test")

	wordCount++

	fmt.Println(wordCount)
}
