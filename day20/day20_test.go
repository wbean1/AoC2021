package day20

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := 35
	input := Input("c:\\Users\\william\\AoC2021\\day20\\day20_test_input.txt")
	paddedInput := PadImage(PadImage(input)))
	algo := GetEnhancementAlgorithm("..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#")
	output := ApplyEnhancementAlgorithm(paddedInput, algo)
	secondOutput := ApplyEnhancementAlgorithm(output, algo)
	got := Count(secondOutput)
	if got != expected {
		t.Errorf("wrong number of pixels lit.  got: %d, expected: %d", got, expected)
	}
}

func TestGet3By3Integer(t *testing.T) {
	expected := 34
	input := Input("c:\\Users\\william\\AoC2021\\day20\\day20_test_input_small.txt")
	got := Get3By3Integer(input, 2, 2)
	if got != expected {
		t.Errorf("wrong 3x3 integer.  got: %d, expected: %d", got, expected)
	}
}
