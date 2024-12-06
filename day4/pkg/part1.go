package day4

import (
	"strings"
)

func Part1(s string) int {
	result := strings.Count(s, "XMAS")
	result += strings.Count(s, "SAMX")
	verticals := get_verticals(s)
	for _, l := range verticals {
		result += strings.Count(l, "XMAS")
		result += strings.Count(l, "SAMX")
	}
	diagonals := get_diagonals(s)
	for _, l := range diagonals {
		result += strings.Count(l, "XMAS")
		result += strings.Count(l, "SAMX")
	}
	return result
}

func get_verticals(s string) []string {
	lines := strings.Split(s, "\n")
	result := []string{}
	for i := 0; i < len(lines[0]); i++ {
		buf := []byte{}
		for j := 0; j < len(lines); j++ {
			buf = append(buf, lines[j][i])
		}
		result = append(result, string(buf))
	}
	return result
}

func get_diagonals(s string) []string {
	matrix := strings.Split(s, "\n")
	rows := len(matrix)
	cols := len(matrix[0])
	diagonals := []string{}
	for col := 0; col < cols; col++ {
		diagonal := []byte{}
		for i, j := 0, col; i < rows && j < cols; i, j = i+1, j+1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, string(diagonal))
	}
	for row := 1; row < rows; row++ {
		diagonal := []byte{}
		for i, j := row, 0; i < rows && j < cols; i, j = i+1, j+1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, string(diagonal))
	}
	for col := cols - 1; col >= 0; col-- {
		diagonal := []byte{}
		for i, j := 0, col; i < rows && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, string(diagonal))
	}
	for row := 1; row < rows; row++ {
		diagonal := []byte{}
		for i, j := row, cols-1; i < rows && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, string(diagonal))
	}
	return diagonals
}
