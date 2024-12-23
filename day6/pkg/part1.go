package day6

import (
	"strings"
)

type Direction uint

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Vec2i struct {
	x, y int
}

type Guard struct {
	direction Direction
	pos       Vec2i
}

func Part1(s string) int {
	result := 0
	stop := false
	guard, scene := load(s)
	for {
		if stop {
			break
		}
		scene[guard.pos.y][guard.pos.x] = 2
		switch guard.direction {
		case UP:
			if guard.pos.y == 0 {
				stop = true
				break
			}
			v := scene[guard.pos.y-1][guard.pos.x]
			if v == 1 {
				guard.direction = (guard.direction + 1) % 4
				break
			}
			if v != 2 {
				result += 1
			}
			guard.pos.y -= 1
		case RIGHT:
			if guard.pos.x == len(scene[0])-1 {
				stop = true
				break
			}
			v := scene[guard.pos.y][guard.pos.x+1]
			if v == 1 {
				guard.direction = (guard.direction + 1) % 4
				break
			}
			if v != 2 {
				result += 1
			}
			guard.pos.x += 1
		case DOWN:
			if guard.pos.y == len(scene)-1 {
				stop = true
				break
			}
			v := scene[guard.pos.y+1][guard.pos.x]
			if v == 1 {
				guard.direction = (guard.direction + 1) % 4
				break
			}
			if v != 2 {
				result += 1
			}
			guard.pos.y += 1
		case LEFT:
			if guard.pos.x == 0 {
				stop = true
				break
			}
			v := scene[guard.pos.y][guard.pos.x-1]
			if v == 1 {
				guard.direction = (guard.direction + 1) % 4
				break
			}
			if v != 2 {
				result += 1
			}
			guard.pos.x -= 1
		}
	}
	return result + 1
}

func load(s string) (Guard, [][]int) {
	guard := Guard{}
	scene := [][]int{}
	for i, line := range strings.Split(s, "\n") {
		scene = append(scene, []int{})
		for j, v := range line {
			t := 0
			if v == '^' {
				guard = Guard{0, Vec2i{j, i}}
			}
			if v == '#' {
				t = 1
			}
			scene[i] = append(scene[i], t)
		}
	}
	return guard, scene
}
