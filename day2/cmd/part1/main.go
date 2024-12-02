package main

import (
	day2 "advent-of-code-2024/day2/pkg"
	"fmt"
	"os"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(buf)
	result := day2.Part1(input)
	fmt.Println(result)
}
