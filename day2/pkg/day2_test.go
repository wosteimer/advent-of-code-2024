package day2

import (
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1("7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9")
	if r != 2 {
		t.Fatalf("result: %d is diferent from 2", r)
	}
}

func TestPart2(t *testing.T) {
}
