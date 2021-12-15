package day10

import (
	"testing"
)

func TestIncorrects(t *testing.T) {
	expected := 26397
	input := Input("/Users/wbean/AoC2021/day10/day10_test_input.txt")
	incorrects, _ := findIncorrectsAndIncompletes(input)
	got := scoreIncorrects(incorrects)
	if got != expected {
		t.Errorf("wrong incorrects score.  got: %d, expected: %d", got, expected)
	}
}

func TestIncompletes(t *testing.T) {
	expected := 288957
	input := Input("/Users/wbean/AoC2021/day10/day10_test_input.txt")
	_, incompletes := findIncorrectsAndIncompletes(input)
	got := scoreIncompletes(incompletes)
	if got != expected {
		t.Errorf("wrong incompletes score.  got: %d, expected: %d", got, expected)
	}
}
