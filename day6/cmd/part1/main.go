package main

import (
	day6 "advent-of-code-2024/day6/pkg"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.TrimSuffix(string(buf), "\n")
	result := day6.Part1(input)
	fmt.Println(result)
}
