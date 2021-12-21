package day20

import (
	"fmt"

	"github.com/wbean1/AoC/utils"
)

func Run() {
	input := Input("c:\\Users\\william\\AoC2021\\day20\\input.txt")
	paddedInput := PadImage(PadImage(input))
	algo := GetEnhancementAlgorithm("###....#.#.#.#.##...#####....##.#..##...##...##..##.######..##.##......#...#.####..#.....##...##.######..##....#...###.##.####.##.####....###..#...##....##...##.###...####.#.##..#..#....#####.#..#...#.#..##..#..##.#.##.#.##...#######.####.#.#..#......##.#...##.#..##....#.###.##.#..####.#......#..#..##.....#.####..#####..###.######..#......####.###.##....##.#.#####..##.#####....#.###..###.#..#..##.##..#.##..###....##.###..#.##.#..########....###.####..##..###..#.#..######..#.##.####.##...#####....#........#.")
	output := ApplyEnhancementAlgorithm(paddedInput, algo, 0)
	secondOutput := ApplyEnhancementAlgorithm(output, algo, 511)
	count := Count(secondOutput)
	fmt.Printf("part1: pixel count is %d\n", count)
	output = input
	for i := 1; i <= 50; i++ {
		output = PadImage(PadImage(PadImage(output)))
	}
	var def int
	for i := 1; i <= 50; i++ {
		if i%2 == 1 {
			def = 0
		} else {
			def = 511
		}
		output = ApplyEnhancementAlgorithm(output, algo, def)
	}
	count = Count(output)
	fmt.Printf("part2: pixel count is %d\n", count)

}

func Get3By3Integer(b [][]bool, x, y int, def int) int {
	maxX := len(b) - 1
	maxY := len(b[0]) - 1
	if x < 2 || y < 2 {
		return def
	} else if x > maxX-1 || y > maxY-1 {
		return def
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

func ApplyEnhancementAlgorithm(b [][]bool, algo []bool, def int) [][]bool {
	output := [][]bool{}
	for x, line := range b {
		newLine := []bool{}
		for y, _ := range line {
			index := Get3By3Integer(b, x, y, def)
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

func GetEnhancementAlgorithm(algoStr string) []bool {
	algo := convertToBools(algoStr)
	return algo
}
