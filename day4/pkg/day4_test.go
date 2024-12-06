package day4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	r := Part1("MMMSXXMASM\n" +
		"MSAMXMSMSA\n" +
		"AMXSXMAAMM\n" +
		"MSAMASMSMX\n" +
		"XMASAMXAMM\n" +
		"XXAMMXXAMA\n" +
		"SMSMSASXSS\n" +
		"SAXAMASAAA\n" +
		"MAMMMXMMMM\n" +
		"MXMXAXMASX",
	)
	if r != 18 {
		t.Fatalf("result: %d is diferent from 18", r)
	}
}

// func TestPart2(t *testing.T) {
// 	r := Part2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
// 	if r != 48 {
// 		t.Fatalf("result: %d is diferent from 48", r)
// 	}
// }
