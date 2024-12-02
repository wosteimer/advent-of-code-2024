package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(s string) int {
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
	sort.Slice(l1, func(i, j int) bool {
		return l1[i] > l1[j]
	})
	sort.Slice(l2, func(i, j int) bool {
		return l2[i] > l2[j]
	})
	sum := 0
	for i := 0; i < len(l1); i++ {
		r := l1[i] - l2[i]
		if r < 0 {
			r *= -1
		}
		sum += r
	}
	return sum
}
