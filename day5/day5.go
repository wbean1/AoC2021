package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const file = "/Users/wbean/AoC/day5/input.txt"

func Input(withDiag bool) [][]int {
	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
	}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x1, y1, x2, y2 := parseLine(scanner.Text())
		if x1 == x2 {
			if y1 <= y2 {
				for y := y1; y <= y2; y++ {
					matrix[x1][y] += 1
				}
			} else {
				for y := y2; y <= y1; y++ {
					matrix[x1][y] += 1
				}
			}
		} else if y1 == y2 {
			if x1 <= x2 {
				for x := x1; x <= x2; x++ {
					matrix[x][y1] += 1
				}
			} else {
				for x := x2; x <= x1; x++ {
					matrix[x][y1] += 1
				}
			}
		} else {
			if !withDiag {
				// log.Printf("skipping line (%d, %d)->(%d, %d) for now\n", x1, y1, x2, y2)
			} else {
				if x1 < x2 && y1 < y2 {
					y := y1
					for x := x1; x <= x2; x++ {
						matrix[x][y] += 1
						y++
					}
				} else if x1 < x2 && y1 > y2 {
					y := y1
					for x := x1; x <= x2; x++ {
						matrix[x][y] += 1
						y--
					}
				} else if x1 > x2 && y1 < y2 {
					y := y1
					for x := x1; x >= x2; x-- {
						matrix[x][y] += 1
						y++
					}
				} else if x1 > x2 && y1 > y2 {
					y := y1
					for x := x1; x >= x2; x-- {
						matrix[x][y] += 1
						y--
					}
				} else {
					log.Fatal("should not be here")
				}
			}
		}
	}
	return matrix
}

func parseLine(line string) (int, int, int, int) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " -> ")
	if len(parts) != 2 {
		log.Fatal("did not get input parts")
	}
	first := strings.Split(parts[0], ",")
	second := strings.Split(parts[1], ",")
	x1, err := strconv.Atoi(first[0])
	if err != nil {
		log.Fatal(err)
	}
	y1, err := strconv.Atoi(first[1])
	if err != nil {
		log.Fatal(err)
	}
	x2, err := strconv.Atoi(second[0])
	if err != nil {
		log.Fatal(err)
	}
	y2, err := strconv.Atoi(second[1])
	if err != nil {
		log.Fatal(err)
	}
	return x1, y1, x2, y2
}

func Run() {
	input := Input(false)
	count := CountCoords(input, 2)
	fmt.Printf("part1: number of coords overlap: %d\n", count)
	input = Input(true)
	count = CountCoords(input, 2)
	fmt.Printf("part2: number of coords overlap: %d\n", count)
}

func CountCoords(input [][]int, limit int) int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if input[x][y] >= limit {
				count += 1
			}
		}
	}
	return count
}
