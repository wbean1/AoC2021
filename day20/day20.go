package day20

import (
	"fmt"

	"github.com/wbean1/AoC/utils"
)

func Run() {
	input := Input("/Users/wbean/AoC2021/day20/input.txt")
	paddedInput := PadImage(input)
	algo := GetEnhancementAlgorithm()
	output := ApplyEnhancementAlgorithm(paddedInput, algo)
	paddedOutput := PadImage(output)
	secondOutput := ApplyEnhancementAlgorithm(paddedOutput, algo)
	count := Count(secondOutput)
	fmt.Printf("part1: pixel count is %d\n", count)
}

func Get3By3Integer(b [][]bool, x, y int) int {
	maxX := len(b) - 1
	maxY := len(b[0]) - 1
	if x < 3 || y < 3 {
		return 0
	} else if x > maxX-2 || y > maxY-2 {
		return 0
	} else {
		myBools := []bool{}
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				myBools = append(myBools, b[i][j])
			}
		}
		myInt := 0
		for _, bit := range myBools {
			myInt = myInt << 1
			if bit {
				myInt++
			}
		}
		return myInt
	}
}

func ApplyEnhancementAlgorithm(b [][]bool, algo []bool) [][]bool {
	output := [][]bool{}
	for x, line := range b {
		newLine := []bool{}
		for y, _ := range line {
			index := Get3By3Integer(b, x, y)
			newLine = append(newLine, algo[index])
		}
		output = append(output, newLine)
	}
	return output
}

func Count(b [][]bool) int {
	count := 0
	for _, line := range b {
		for _, ok := range line {
			if ok {
				count++
			}
		}
	}
	return count
}

func PadImage(b [][]bool) [][]bool {
	new := [][]bool{}
	for _, line := range b {
		line = append(line, []bool{false, false, false, false}...)
		line = append([]bool{false, false, false, false}, line...)
		new = append(new, line)
	}
	blankLine := []bool{}
	for i := 0; i < len(new[0]); i++ {
		blankLine = append(blankLine, false)
	}
	for i := 0; i < 4; i++ {
		new = append(new, [][]bool{blankLine, blankLine, blankLine, blankLine}...)
		new = append([][]bool{blankLine, blankLine, blankLine, blankLine}, new...)
	}
	return new
}

func Input(file string) [][]bool {
	lines := utils.ParseFileToStrings(file)
	image := [][]bool{}
	for _, line := range lines {
		image = append(image, convertToBools(line))
	}
	return image
}

func convertToBools(s string) []bool {
	line := []bool{}
	for _, char := range s {
		if char == rune('#') {
			line = append(line, true)
		} else {
			line = append(line, false)
		}
	}
	return line
}

func GetEnhancementAlgorithm() []bool {
	algoStr := "###....#.#.#.#.##...#####....##.#..##...##...##..##.######..##.##......#...#.####..#.....##...##.######..##....#...###.##.####.##.####....###..#...##....##...##.###...####.#.##..#..#....#####.#..#...#.#..##..#..##.#.##.#.##...#######.####.#.#..#......##.#...##.#..##....#.###.##.#..####.#......#..#..##.....#.####..#####..###.######..#......####.###.##....##.#.#####..##.#####....#.###..###.#..#..##.##..#.##..###....##.###..#.##.#..########....###.####..##..###..#.#..######..#.##.####.##...#####....#........#."
	algo := convertToBools(algoStr)
	return algo
}
