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

func TestPart2(t *testing.T) {
	r := Part2("MMMSXXMASM\n" +
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
	if r != 9 {
		t.Fatalf("result: %d is diferent from 9", r)
	}
}
