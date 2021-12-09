package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Coord struct {
	X, Y int
}

func Run() {
	input := Input()
	lows, lowCoords := findLows(input)
	sum := 0
	for _, low := range lows {
		sum += low + 1
	}
	fmt.Printf("part1: sum of lows is: %d\n", sum)
	basins := findBasins(input, lowCoords)
	basinSizes := []int{}
	for _, basin := range basins {
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Ints(basinSizes)
	sum = basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
	fmt.Printf("part2: product of 3 largest basins is: %d\n", sum)
}

func findBasins(d [][]int, c []Coord) map[Coord][]Coord {
	basins := make(map[Coord][]Coord)
	for _, coord := range c {
		basins[coord] = []Coord{coord} // the low point is in the basin by default
		d[coord.X][coord.Y] = 9        // flip to a 9 as already seen
		if coord.X < len(d)-1 {
			AddPointToBasin(basins, d, coord, Coord{coord.X + 1, coord.Y})
		}
		if coord.X > 0 {
			AddPointToBasin(basins, d, coord, Coord{coord.X - 1, coord.Y})
		}
		if coord.Y < len(d)-1 {
			AddPointToBasin(basins, d, coord, Coord{coord.X, coord.Y + 1})
		}
		if coord.Y > 0 {
			AddPointToBasin(basins, d, coord, Coord{coord.X, coord.Y - 1})
		}
	}
	return basins
}

func AddPointToBasin(basins map[Coord][]Coord, d [][]int, basinLow Coord, lookingAt Coord) {
	if d[lookingAt.X][lookingAt.Y] == 9 {
		return
	}
	basins[basinLow] = append(basins[basinLow], lookingAt)
	d[lookingAt.X][lookingAt.Y] = 9
	if lookingAt.X < len(d)-1 {
		AddPointToBasin(basins, d, basinLow, Coord{lookingAt.X + 1, lookingAt.Y})
	}
	if lookingAt.X > 0 {
		AddPointToBasin(basins, d, basinLow, Coord{lookingAt.X - 1, lookingAt.Y})
	}
	if lookingAt.Y < len(d)-1 {
		AddPointToBasin(basins, d, basinLow, Coord{lookingAt.X, lookingAt.Y + 1})
	}
	if lookingAt.Y > 0 {
		AddPointToBasin(basins, d, basinLow, Coord{lookingAt.X, lookingAt.Y - 1})
	}
}

func findLows(d [][]int) ([]int, []Coord) {
	lows := []int{}
	lowCoords := []Coord{}
	for x, line := range d {
		for y, value := range line {
			if x == 0 {
				if y == 0 {
					if value < d[x+1][y] &&
						value < d[x][y+1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				} else if y == len(line)-1 {
					if value < d[x+1][y] &&
						value < d[x][y-1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				} else {
					if value < d[x+1][y] &&
						value < d[x][y+1] &&
						value < d[x][y-1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				}
			} else if x == len(d)-1 {
				if y == 0 {
					if value < d[x-1][y] &&
						value < d[x][y+1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				} else if y == len(line)-1 {
					if value < d[x-1][y] &&
						value < d[x][y-1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				} else {
					if value < d[x-1][y] &&
						value < d[x][y+1] &&
						value < d[x][y-1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				}
			} else {
				if y == 0 {
					if value < d[x-1][y] &&
						value < d[x+1][y] &&
						value < d[x][y+1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				} else if y == len(line)-1 {
					if value < d[x-1][y] &&
						value < d[x+1][y] &&
						value < d[x][y-1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				} else {
					if value < d[x-1][y] &&
						value < d[x+1][y] &&
						value < d[x][y+1] &&
						value < d[x][y-1] {
						lows = append(lows, value)
						lowCoords = append(lowCoords, Coord{x, y})
					}
				}
			}
		}
	}
	return lows, lowCoords
}

func Input() [][]int {
	lines := parseFile("/Users/wbean/AoC2021/day9/input.txt")
	input := make([][]int, len(lines))
	for x, line := range lines {
		input[x] = make([]int, len(line))
		for y, char := range line {
			input[x][y] = int(char) - 48
		}
	}
	return input
}

func parseFile(f string) []string {

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
