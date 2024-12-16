package main

import (
	day5 "advent-of-code-2024/day5/pkg"
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
	result := day5.Part2(input)
	fmt.Println(result)
}
