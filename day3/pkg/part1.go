package day3

import (
	"regexp"
	"strconv"
)

func Part1(s string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	m := r.FindAllStringSubmatch(s, -1)
	result := 0
	for _, match := range m {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		result += a * b
	}
	return result
}
