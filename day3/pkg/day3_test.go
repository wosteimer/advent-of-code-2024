package day3

import (
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	if r != 161 {
		t.Fatalf("result: %d is diferent from 2", r)
	}
}

// func TestPart2(t *testing.T) {
// 	r := Part2("7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9")
// 	if r != 4 {
// 		t.Fatalf("result: %d is diferent from 4", r)
// 	}
// }
