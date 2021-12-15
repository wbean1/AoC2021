package day11

import (
	"fmt"

	"github.com/wbean1/AoC/utils"
)

type OctoState [][]int

func Run() {
	state := Input("/Users/wbean/AoC2021/day11/input.txt")
	flashCount := 0
	for i := 1; i <= 100; i++ {
		flashCount += state.Step()
	}
	fmt.Printf("part1: flashed %d times\n", flashCount)
	state = Input("/Users/wbean/AoC2021/day11/input.txt")
	step := state.stepTilAllFlash()
	fmt.Printf("part2: all flashed on step %d\n", step)
}

func (s *OctoState) stepTilAllFlash() int {
	for i := 1; 1 > 0; i++ {
		s.Step()
		if s.isAllZero() {
			return i
		}
	}
	return 0
}

func (s *OctoState) isAllZero() bool {
	for x := 0; x < len(*s); x++ {
		for y := 0; y < len((*s)[x]); y++ {
			if (*s)[x][y] != 0 {
				return false
			}
		}
	}
	return true
}

func (s *OctoState) Step() int {
	flashCount := 0
	alreadyFlashed := make([][]bool, len(*s))
	for x := 0; x < len(*s); x++ {
		alreadyFlashed[x] = make([]bool, len((*s)[x]))
		for y := 0; y < len((*s)[x]); y++ {
			(*s)[x][y]++
		}
	}
	for x := 0; x < len(*s); x++ {
		for y := 0; y < len((*s)[x]); y++ {
			if (*s)[x][y] > 9 && !alreadyFlashed[x][y] {
				flashCount += s.Flash(x, y, alreadyFlashed)
			}
		}
	}
	for x := 0; x < len(*s); x++ {
		for y := 0; y < len((*s)[x]); y++ {
			if (*s)[x][y] > 9 {
				(*s)[x][y] = 0
			}
		}
	}
	return flashCount
}

func (s *OctoState) Flash(x, y int, alreadyFlashed [][]bool) int {
	// if x,y == 10, flash it and everything around it
	flashCount := 0
	if (*s)[x][y] > 9 && !alreadyFlashed[x][y] {
		alreadyFlashed[x][y] = true
		flashCount++
		if x == 0 {
			if y == 0 {
				(*s)[x+1][y]++
				flashCount += s.Flash(x+1, y, alreadyFlashed)
				(*s)[x+1][y+1]++
				flashCount += s.Flash(x+1, y+1, alreadyFlashed)
				(*s)[x][y+1]++
				flashCount += s.Flash(x, y+1, alreadyFlashed)
			} else if y > 0 && y < len((*s)[x])-1 {
				(*s)[x+1][y]++
				flashCount += s.Flash(x+1, y, alreadyFlashed)
				(*s)[x+1][y-1]++
				flashCount += s.Flash(x+1, y-1, alreadyFlashed)
				(*s)[x+1][y+1]++
				flashCount += s.Flash(x+1, y+1, alreadyFlashed)
				(*s)[x][y-1]++
				flashCount += s.Flash(x, y-1, alreadyFlashed)
				(*s)[x][y+1]++
				flashCount += s.Flash(x, y+1, alreadyFlashed)
			} else {
				(*s)[x+1][y]++
				flashCount += s.Flash(x+1, y, alreadyFlashed)
				(*s)[x+1][y-1]++
				flashCount += s.Flash(x+1, y-1, alreadyFlashed)
				(*s)[x][y-1]++
				flashCount += s.Flash(x, y-1, alreadyFlashed)
			}
		} else if x > 0 && x < len(*s)-1 {
			if y == 0 {
				(*s)[x+1][y]++
				flashCount += s.Flash(x+1, y, alreadyFlashed)
				(*s)[x+1][y+1]++
				flashCount += s.Flash(x+1, y+1, alreadyFlashed)
				(*s)[x][y+1]++
				flashCount += s.Flash(x, y+1, alreadyFlashed)
				(*s)[x-1][y]++
				flashCount += s.Flash(x-1, y, alreadyFlashed)
				(*s)[x-1][y+1]++
				flashCount += s.Flash(x-1, y+1, alreadyFlashed)
			} else if y > 0 && y < len((*s)[x])-1 {
				(*s)[x+1][y]++
				flashCount += s.Flash(x+1, y, alreadyFlashed)
				(*s)[x+1][y-1]++
				flashCount += s.Flash(x+1, y-1, alreadyFlashed)
				(*s)[x+1][y+1]++
				flashCount += s.Flash(x+1, y+1, alreadyFlashed)
				(*s)[x][y-1]++
				flashCount += s.Flash(x, y-1, alreadyFlashed)
				(*s)[x][y+1]++
				flashCount += s.Flash(x, y+1, alreadyFlashed)
				(*s)[x-1][y]++
				flashCount += s.Flash(x-1, y, alreadyFlashed)
				(*s)[x-1][y-1]++
				flashCount += s.Flash(x-1, y-1, alreadyFlashed)
				(*s)[x-1][y+1]++
				flashCount += s.Flash(x-1, y+1, alreadyFlashed)
			} else {
				(*s)[x+1][y]++
				flashCount += s.Flash(x+1, y, alreadyFlashed)
				(*s)[x+1][y-1]++
				flashCount += s.Flash(x+1, y-1, alreadyFlashed)
				(*s)[x][y-1]++
				flashCount += s.Flash(x, y-1, alreadyFlashed)
				(*s)[x-1][y]++
				flashCount += s.Flash(x-1, y, alreadyFlashed)
				(*s)[x-1][y-1]++
				flashCount += s.Flash(x-1, y-1, alreadyFlashed)
			}
		} else {
			if y == 0 {
				(*s)[x-1][y]++
				flashCount += s.Flash(x-1, y, alreadyFlashed)
				(*s)[x-1][y+1]++
				flashCount += s.Flash(x-1, y+1, alreadyFlashed)
				(*s)[x][y+1]++
				flashCount += s.Flash(x, y+1, alreadyFlashed)
			} else if y > 0 && y < len((*s)[x])-1 {
				(*s)[x-1][y]++
				flashCount += s.Flash(x-1, y, alreadyFlashed)
				(*s)[x-1][y-1]++
				flashCount += s.Flash(x-1, y-1, alreadyFlashed)
				(*s)[x-1][y+1]++
				flashCount += s.Flash(x-1, y+1, alreadyFlashed)
				(*s)[x][y-1]++
				flashCount += s.Flash(x, y-1, alreadyFlashed)
				(*s)[x][y+1]++
				flashCount += s.Flash(x, y+1, alreadyFlashed)
			} else {
				(*s)[x-1][y]++
				flashCount += s.Flash(x-1, y, alreadyFlashed)
				(*s)[x-1][y-1]++
				flashCount += s.Flash(x-1, y-1, alreadyFlashed)
				(*s)[x][y-1]++
				flashCount += s.Flash(x, y-1, alreadyFlashed)
			}
		}
	}
	return flashCount
}

func Input(file string) OctoState {
	return utils.ParseFileToTwoDIntArray(file)
}
