package day15

import (
	"fmt"

	"github.com/wbean1/AoC/utils"
)

func Run() {
	input := Input()
	minLength := findMinPathLength(input)
	fmt.Printf("part1: min path length is: %d\n", minLength)
	newInput := expandMap(input)
	minLength2 := findMinPathLength(newInput)
	fmt.Printf("part2: min path length is: %d\n", minLength2)
}

func Input() [][]int {
	return utils.ParseFileToTwoDIntArray("c:\\Users\\William\\AoC2021\\day15\\input.txt")
}

func expandMap(d [][]int) [][]int {
	m := make([][]int, len(d)*5)
	for x, _ := range m {
		m[x] = make([]int, len(d[0])*5)
	}

	for x := 0; x <= len(m)-1; x++ {
		for y := 0; y <= len(m[x])-1; y++ {
			dX := x % len(d)
			dY := y % len(d[0])
			xBox := x / len(d)
			yBox := y / len(d)
			boxAdd := xBox + yBox
			m[x][y] = d[dX][dY] + boxAdd
			if m[x][y] > 9 {
				m[x][y] = m[x][y] % 9
			}
		}
	}
	fmt.Println(m)
	return m
}

func findMinPathLength(d [][]int) int {
	m := make([][]int, len(d))
	for x, _ := range d {
		m[x] = make([]int, len(d[x]))
	}

	for x := 0; x <= len(d)-1; x++ {
		for y := 0; y <= len(d[x])-1; y++ {
			if x == 0 && y == 0 {
				m[x][y] = 0
			} else if x == 0 && y > 0 {
				m[x][y] = m[x][y-1] + d[x][y]
			} else if x > 0 && y == 0 {
				m[x][y] = m[x-1][y] + d[x][y]
			} else if x > 0 && y > 0 {
				m[x][y] = min(m[x-1][y], m[x][y-1]) + d[x][y]
			}
		}
	}
	for i := 0; i < 100; i++ { // find corrections... real ugly omg
		for x := 0; x <= len(d)-1; x++ {
			for y := 0; y <= len(d[x])-1; y++ {
				if x == 0 && y == 0 {
					m[x][y] = 0
				} else if x == 0 && y > 0 {
					m[x][y] = m[x][y-1] + d[x][y]
				} else if x > 0 && y == 0 {
					m[x][y] = m[x-1][y] + d[x][y]
				} else if x > 0 && y > 0 && x != len(d)-1 && y != len(d[x])-1 {
					fromLeftOrUp := min(m[x-1][y], m[x][y-1]) + d[x][y]
					fromRightOrDown := min(m[x+1][y], m[x][y+1]) + d[x][y]
					m[x][y] = min(fromLeftOrUp, fromRightOrDown)
				} else if x > 0 && y > 0 && x != len(d)-1 && y == len(d[x])-1 {
					fromLeftOrUp := min(m[x-1][y], m[x][y-1]) + d[x][y]
					fromRight := m[x+1][y] + d[x][y]
					m[x][y] = min(fromLeftOrUp, fromRight)
				} else if x > 0 && y > 0 && x == len(d)-1 && y != len(d[x])-1 {
					fromLeftOrUp := min(m[x-1][y], m[x][y-1]) + d[x][y]
					fromDown := m[x][y+1] + d[x][y]
					m[x][y] = min(fromLeftOrUp, fromDown)
				} else if x == len(d)-1 && y == len(d[x])-1 {
					m[x][y] = min(m[x-1][y], m[x][y-1]) + d[x][y]
				}
			}
		}
	}
	return m[len(d)-1][len(d[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
