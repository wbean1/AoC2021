package utils

import (
	"bufio"
	"log"
	"os"
)

func ParseFileToStrings(f string) []string {
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

func ParseFileToTwoDIntArray(f string) [][]int {
	result := ParseFileToStrings(f)
	m := make([][]int, len(result))
	for x, line := range result {
		m[x] = make([]int, len(line))
		for y, char := range line {
			m[x][y] = int(char) - 48
		}
	}
	return m
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
