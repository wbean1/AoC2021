package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func Run() {
	input := Input()
	newCoords := FoldVertical(input, 655)
	fmt.Printf("part1: there are %d dots visible\n", len(newCoords))
	newCoords = FoldHorizontal(newCoords, 447)
	newCoords = FoldVertical(newCoords, 327)
	newCoords = FoldHorizontal(newCoords, 223)
	newCoords = FoldVertical(newCoords, 163)
	newCoords = FoldHorizontal(newCoords, 111)
	newCoords = FoldVertical(newCoords, 81)
	newCoords = FoldHorizontal(newCoords, 55)
	newCoords = FoldVertical(newCoords, 40)
	newCoords = FoldHorizontal(newCoords, 27)
	newCoords = FoldHorizontal(newCoords, 13)
	newCoords = FoldHorizontal(newCoords, 6)
	d := make([][]string, 40)
	for i, _ := range d {
		d[i] = make([]string, 8)
	}
	for _, coord := range newCoords {
		d[coord.X][coord.Y] = "X"
	}
	for x, _ := range d {
		for y, str := range d[x] {
			if str != "X" {
				d[x][y] = "."
			}
		}
	}
	for _, arr := range d {
		fmt.Println(arr)
	}
}

func FoldVertical(c []Coord, xFold int) []Coord {
	newCoords := []Coord{}
	for _, coord := range c {
		if coord.X < xFold {
			newCoords = append(newCoords, coord)
		} else if coord.X == xFold {
			log.Fatalf("coord's X matches xFold: %d", xFold)
		} else {
			newCoords = append(newCoords, Coord{xFold - (coord.X - xFold), coord.Y})
		}
	}
	return Dedupe(newCoords)
}

func FoldHorizontal(c []Coord, yFold int) []Coord {
	newCoords := []Coord{}
	for _, coord := range c {
		if coord.Y < yFold {
			newCoords = append(newCoords, coord)
		} else if coord.Y == yFold {
			log.Fatalf("coord's Y matches yFold: %d", yFold)
		} else {
			newCoords = append(newCoords, Coord{coord.X, yFold - (coord.Y - yFold)})
		}
	}
	return Dedupe(newCoords)
}

func Dedupe(c []Coord) []Coord {
	coordMap := make(map[Coord]bool)
	for _, coord := range c {
		coordMap[coord] = true
	}
	newCoords := []Coord{}
	for coord, _ := range coordMap {
		newCoords = append(newCoords, coord)
	}
	return newCoords
}

func Input() []Coord {
	return parseFileToCoords("c:\\Users\\william\\AoC2021\\day13\\input.txt")
}

func parseFileToCoords(f string) []Coord {
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
	coords := []Coord{}
	for _, line := range result {
		parts := strings.Split(line, ",")
		part1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		part2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		coords = append(coords, Coord{part1, part2})
	}
	return coords
}
