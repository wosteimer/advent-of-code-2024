package main

import (
	day1 "advent-of-code-2024/day1/pkg"
	"fmt"
	"os"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(buf)
	result := day1.Part2(input)
	fmt.Println(result)
}
