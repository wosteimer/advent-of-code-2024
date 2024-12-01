package main

import (
	day1 "advent-of-code-2024/day-1/pkg"
	"fmt"
	"os"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(buf)
	result := day1.Part1(input)
	fmt.Println(result)
}
