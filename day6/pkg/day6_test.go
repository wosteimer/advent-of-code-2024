package day6

import (
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1("....#.....\n" +
		".........#\n" +
		"..........\n" +
		"..#.......\n" +
		".......#..\n" +
		"..........\n" +
		".#..^.....\n" +
		"........#.\n" +
		"#.........\n" +
		"......#...",
	)
	if r != 41 {
		t.Fatalf("result: %d is diferent from 41", r)
	}
}

// func TestPart2(t *testing.T) {
// 	r := Part2("")
// 	if r != 123 {
// 		t.Fatalf("result: %d is diferent from 123", r)
// 	}
// }
