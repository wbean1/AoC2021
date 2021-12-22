package day22

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/wbean1/AoC/utils"
)

type Coord struct {
	x, y, z int
}

type Operation struct {
	op       bool
	min, max Coord
}

func Run() {
	ops := Input("c:\\Users\\william\\AoC2021\\day22\\input_part1.txt")
	gameBoard := initGameBoard(101)
	for _, operation := range ops {
		PerformOperation(gameBoard, operation, 50)
	}
	fmt.Printf("part1: cubes on: %d\n", Count(gameBoard))

	ops = Input("c:\\Users\\william\\AoC2021\\day22\\input_part2.txt")
	gameBoard = initGameBoard(200001)
	for i, operation := range ops {
		fmt.Printf("performing op #%d\n", i)
		PerformOperation(gameBoard, operation, 100000)
	}
	fmt.Println("starting count...")
	fmt.Printf("part2: cubes on: %d\n", CountUint64(gameBoard))
}

func PerformOperation(b [][][]bool, op Operation, indexAdd int) [][][]bool {
	for i := op.min.x + indexAdd; i <= op.max.x+indexAdd; i++ {
		for j := op.min.y + indexAdd; j <= op.max.y+indexAdd; j++ {
			for k := op.min.z + indexAdd; k <= op.max.z+indexAdd; k++ {
				b[i][j][k] = op.op
			}
		}
	}
	return b
}

func initGameBoard(size int) [][][]bool {
	gameBoard := make([][][]bool, size)
	for x := range gameBoard {
		gameBoard[x] = make([][]bool, size)
		for y := range gameBoard[x] {
			gameBoard[x][y] = make([]bool, size)
		}
	}
	return gameBoard
}

func CountUint64(b [][][]bool) uint64 {
	count := uint64(0)
	for x := range b {
		for y := range b[x] {
			for _, val := range b[x][y] {
				if val {
					count++
				}
			}
		}
	}
	return count
}

func Count(b [][][]bool) int {
	count := 0
	for x := range b {
		for y := range b[x] {
			for _, val := range b[x][y] {
				if val {
					count++
				}
			}
		}
	}
	return count
}

func Input(file string) []Operation {
	operations := []Operation{}
	i := utils.ParseFileToStrings(file)
	for _, line := range i {
		op := Operation{}
		opStr := line[0:3]
		if opStr == "off" {
			op.op = false
		} else {
			op.op = true
		}
		re := regexp.MustCompile(`=(-?\d+..-?\d+)`)
		found := re.FindAllSubmatch([]byte(line), -1)
		var min, max Coord
		for i, subline := range found {
			re := regexp.MustCompile(`(-?\d+)`)
			nums := re.FindAllSubmatch(subline[1], -1)
			minInt, _ := strconv.Atoi(string(nums[0][0]))
			maxInt, _ := strconv.Atoi(string(nums[1][0]))
			if i == 0 {
				min.x = minInt
				max.x = maxInt
			} else if i == 1 {
				min.y = minInt
				max.y = maxInt
			} else {
				min.z = minInt
				max.z = maxInt
			}
		}
		op.min = min
		op.max = max
		operations = append(operations, op)
	}
	return operations
}
