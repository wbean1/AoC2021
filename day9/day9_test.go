package day9

import (
	"fmt"
	"sort"
	"testing"
)

func TestSumOfLows(t *testing.T) {
	expected := 15
	input := Input("/Users/wbean/AoC2021/day9/day9_test_input.txt")
	lows, _ := findLows(input)
	got := 0
	for _, low := range lows {
		got += low + 1
	}
	if got != expected {
		t.Errorf("wrong SumOfLows.  got: %d, expected: %d", got, expected)
	}
}

func TestProductOf3LargestBasins(t *testing.T) {
	expected := 1134
	input := Input("/Users/wbean/AoC2021/day9/day9_test_input.txt")
	_, lowCoords := findLows(input)
	basins := findBasins(input, lowCoords)
	fmt.Println(basins)
	basinSizes := []int{}
	for _, basin := range basins {
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Ints(basinSizes)
	got := basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
	if got != expected {
		t.Errorf("wrong ProductOf3LargestBasins.  got: %d, expected: %d", got, expected)
	}
}
