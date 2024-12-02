package day2

import (
	"strconv"
	"strings"
)

func Part2(s string) int {
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
		if is_safe_2(record) {
			result += 1
		}
	}
	return result
}

func is_safe_2(record []int) bool {
	if is_safe(record) {
		return true
	}
	for i := 0; i < len(record); i++ {
		if is_safe(remove(record, i)) {
			return true
		}
	}
	return false
}

func remove(r []int, to_remove int) []int {
	result := []int{}
	for i, v := range r {
		if to_remove == i {
			continue
		}
		result = append(result, v)
	}
	return result
}
