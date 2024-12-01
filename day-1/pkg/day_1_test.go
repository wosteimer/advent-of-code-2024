package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1("3   4\n4   3\n2   5\n1   3\n3   9\n3   3")
	if r != 11 {
		t.Fatalf("")
	}
}

func TestPart2(t *testing.T) {
	r := Part2("3   4\n4   3\n2   5\n1   3\n3   9\n3   3")
	if r != 31 {
		t.Fatalf("")
	}
}
