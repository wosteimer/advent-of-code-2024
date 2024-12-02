package day2

import (
	"strconv"
	"strings"
)

func Part1(s string) int {
	lines := strings.Split(s, "\n")
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		record := []int{}
		for _, v := range strings.Split(line, " ") {
			vi, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			record = append(record, vi)
		}
		if is_safe(record) {
			result += 1
		}
	}
	return result
}

func is_safe(record []int) bool {
	is_increasing := record[0] < record[1]
	for i := 0; i < len(record)-1; i++ {
		a := record[i]
		b := record[i+1]
		if is_increasing {
			if a >= b {
				return false
			}
		} else {
			if a <= b {
				return false
			}
		}
		d := max(a-b, b-a)
		if d < 1 || d > 3 {
			return false
		}
	}
	return true
}
