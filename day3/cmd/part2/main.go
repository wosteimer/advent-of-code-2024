package main

import (
	day3 "advent-of-code-2024/day3/pkg"
	"fmt"
	"os"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(buf)
	result := day3.Part2(input)
	fmt.Println(result)
}
