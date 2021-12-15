package day11

import (
	"testing"
)

func Test100Flashes(t *testing.T) {
	expected := 1656
	state := Input("/Users/wbean/AoC2021/day11/day11_test_input.txt")
	got := 0
	for i := 1; i <= 100; i++ {
		got += state.Step()
	}
	if got != expected {
		t.Errorf("wrong number of flashes.  got: %d, expected: %d", got, expected)
	}
}
