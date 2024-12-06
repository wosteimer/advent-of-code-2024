package day4

import "strings"

func Part2(s string) int {
	lines := strings.Split(s, "\n")
	rows := len(lines)
	cols := len(lines[0])
	result := 0
	for i := 0; i < rows-2; i++ {
		for j := 0; j < cols-2; j++ {
			m := [3][3]byte{}
			m[0][0] = lines[i][j]
			m[0][2] = lines[i][j+2]
			m[1][1] = lines[i+1][j+1]
			m[2][0] = lines[i+2][j]
			m[2][2] = lines[i+2][j+2]
			for i := 0; i <= 4; i++ {
				if is_x_mas(m) {
					result += 1
					break
				}
				m = rotate(m)
			}
		}
	}
	return result
}

func is_x_mas(m [3][3]byte) bool {
	return m[0][0] == 'M' &&
		m[0][2] == 'S' &&
		m[1][1] == 'A' &&
		m[2][0] == 'M' &&
		m[2][2] == 'S'
}

func rotate(m [3][3]byte) [3][3]byte {
	n := 3
	rotated := [3][3]byte{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-i-1] = m[i][j]
		}
	}
	return rotated
}
