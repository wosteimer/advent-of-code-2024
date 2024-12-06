package day3

import (
	"strconv"
)

func Part2(s string) int {
	return state_1(s, 0)
}

func state_1(s string, n int) int {
	result := 0
	for {
		if n+3 >= len(s) {
			return result
		}
		if s[n:n+3] == "mul" {
			n += 3
			if s[n] != '(' {
				n += 1
				continue
			}
			a := ""
			b := ""
			for {
				n += 1
				if n >= len(s) {
					return result
				}
				_, err := strconv.Atoi(string(s[n]))
				if err != nil {
					break
				}
				a += string(s[n])
			}
			if a == "" || s[n] != ',' {
				n += 1
				continue
			}
			for {
				n += 1
				if n >= len(s) {
					return result
				}
				_, err := strconv.Atoi(string(s[n]))
				if err != nil {
					break
				}
				b += string(s[n])

			}
			if b == "" || s[n] != ')' {
				n += 1
				continue
			}
			an, _ := strconv.Atoi(a)
			bn, _ := strconv.Atoi(b)
			result += an * bn
			continue
		}
		if n+7 >= len(s) {
			return result
		}
		if s[n:n+7] == "don't()" {
			n += 7
			return state_2(s, n) + result
		}
		n += 1
	}
}

func state_2(s string, n int) int {
	for i := n; i < len(s)-4; i++ {
		if s[i:i+4] == "do()" {
			return state_1(s, i+4)
		}
	}
	return 0
}
