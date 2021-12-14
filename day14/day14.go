package day14

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/wbean1/AoC/utils"
)

func Run() {
	inserts := Input()
	polymer := "BSONBHNSSCFPSFOPHKPK"
	part1Polymer := runPolymerSteps([]byte(polymer), inserts, 10)
	maxElementCount, minElementCount := findMaxAndMinLetterCount(part1Polymer)
	fmt.Printf("part1: max(%d) - min(%d) = %d\n", maxElementCount, minElementCount, maxElementCount-minElementCount)
	part2Polymer := runPolymerSteps([]byte(polymer), inserts, 40)
	maxElementCount, minElementCount = findMaxAndMinLetterCount(part2Polymer)
	fmt.Printf("part2: max(%d) - min(%d) = %d\n", maxElementCount, minElementCount, maxElementCount-minElementCount)
}

func runPolymerSteps(polymer []byte, inserts map[string]byte, numSteps int) string {
	for i := 1; i <= numSteps; i++ {
		fmt.Printf("starting step %d with polymer len: %d\n", i, len(polymer))
		var newPolymer bytes.Buffer
		for x, currentByte := range polymer[:len(polymer)-1] {
			nextByte := polymer[x+1]
			if val, ok := inserts[string([]byte{currentByte, nextByte})]; ok {
				// there is an insert for this
				newPolymer.WriteByte(currentByte)
				newPolymer.WriteByte(val)
			} else {
				// there is not an insert for this
				newPolymer.WriteByte(currentByte)
			}
		}
		newPolymer.WriteByte(polymer[len(polymer)-1])
		polymer = newPolymer.Bytes()
	}
	return string(polymer)
}

func findMaxAndMinLetterCount(str string) (uint64, uint64) {
	charMap := make(map[rune]uint64)
	for _, char := range str {
		charMap[char]++
	}
	max := uint64(0)
	min := uint64(18446744073709551615)
	for _, count := range charMap {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return max, min
}

func Input() map[string]byte {
	input := utils.ParseFileToStrings("c:\\Users\\William\\AoC2021\\day14\\input.txt")
	inserts := make(map[string]byte)
	for _, str := range input {
		parts := strings.Split(str, " -> ")
		if len(parts) != 2 {
			log.Fatalf("%s doesn't split right", str)
		}
		inserts[parts[0]] = []byte(parts[1])[0]
	}
	return inserts
}
