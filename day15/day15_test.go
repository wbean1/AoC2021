package day15

import (
	"testing"

	"github.com/wbean1/AoC/utils"
)

func TestFindMinPathLength(t *testing.T) {
	expected := 40
	i := utils.ParseFileToTwoDIntArray("c:\\Users\\William\\AoC2021\\day15\\day15_test_input.txt")
	got := findMinPathLength(i)
	if got != expected {
		t.Errorf("wrong minPathLength.  got: %d, expected: %d", got, expected)
	}
}

func TestExpandedMinPathLength(t *testing.T) {
	expected := 315
	i := utils.ParseFileToTwoDIntArray("c:\\Users\\William\\AoC2021\\day15\\day15_test_input.txt")
	got := findMinPathLength(expandMap(i))
	if got != expected {
		t.Errorf("wrong minPathLength.  got: %d, expected: %d", got, expected)
	}
}
