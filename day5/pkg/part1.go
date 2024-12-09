package day5

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(s string) int {
	splited := strings.Split(s, "\n\n")
	rules := map[int][]int{}
	result := 0
	for _, rule := range strings.Split(splited[0], "\n") {
		rs := strings.Split(rule, "|")
		v0, _ := strconv.Atoi(rs[0])
		v1, _ := strconv.Atoi(rs[1])
		rules[v0] = append(rules[v0], v1)
	}
	for _, us := range strings.Split(splited[1], "\n") {
		is_valid := true
		update := map_to_int(us)
		for i, v := range update {
			rule, ok := rules[v]
			if !ok {
				continue
			}
			for _, r := range rule {
				if slices.Contains(update[:i], r) {
					is_valid = false
				}
			}
		}
		if is_valid {
			result += update[len(update)/2]
		}

	}
	return result
}

func map_to_int(s string) []int {
	result := []int{}
	for _, vs := range strings.Split(s, ",") {
		v, _ := strconv.Atoi(vs)
		result = append(result, v)
	}
	return result
}
