package day5

import (
	"slices"
	"strconv"
	"strings"
)

func Part2(s string) int {
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
		update := parse_update(us)
		if !is_valid(update, rules) {
			new_update := []int{}
			for _, v := range update {
				for j := 0; j < len(new_update)+1; j++ {
					temp := slices.Concat(new_update[:j], []int{v}, new_update[j:])
					if is_valid(temp, rules) {
						new_update = temp
						break
					}
				}
			}
			result += new_update[len(update)/2]
		}

	}
	return result
}

func is_valid(update []int, rules map[int][]int) bool {
	for i, v := range update {
		rule, ok := rules[v]
		if !ok {
			continue
		}
		for _, r := range rule {
			if slices.Contains(update[:i], r) {
				return false
			}
		}
	}
	return true
}
