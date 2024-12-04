package day3

import (
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	if r != 161 {
		t.Fatalf("result: %d is diferent from 161", r)
	}
}

func TestPart2(t *testing.T) {
	r := Part2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	if r != 48 {
		t.Fatalf("result: %d is diferent from 48", r)
	}
}
