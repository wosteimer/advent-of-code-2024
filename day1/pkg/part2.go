package day1

import (
	"strconv"
	"strings"
)

func count(l []int, v int) int {
	result := 0
	for _, v2 := range l {
		if v2 == v {
			result += 1
		}
	}
	return result
}

func Part2(s string) int {
	lines := strings.Split(s, "\n")
	l1 := []int{}
	l2 := []int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		va := strings.Split(line, "   ")
		v1, err := strconv.Atoi(va[0])
		if err != nil {
			panic(err)
		}
		v2, err := strconv.Atoi(va[1])
		if err != nil {
			panic(err)
		}
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}
	sum := 0
	for _, v := range l1 {
		c := count(l2, v)
		sum += v * c
	}
	return sum
}
