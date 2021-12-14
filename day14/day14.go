package day14

import (
	"fmt"
	"log"
	"strings"

	"github.com/wbean1/AoC/utils"
)

func Run() {
	inserts := Input()
	polymer := "BSONBHNSSCFPSFOPHKPK"
	pairMapCount := makePairCountMap(polymer)
	part1Polymer := runPolymerSteps(pairMapCount, inserts, 10)
	maxElementCount, minElementCount := findMaxAndMinLetterCount(part1Polymer, []byte("K")[0])
	fmt.Printf("part1: max(%d) - min(%d) = %d\n", maxElementCount, minElementCount, maxElementCount-minElementCount)
	part2Polymer := runPolymerSteps(pairMapCount, inserts, 40)
	maxElementCount, minElementCount = findMaxAndMinLetterCount(part2Polymer, []byte("K")[0])
	fmt.Printf("part2: max(%d) - min(%d) = %d\n", maxElementCount, minElementCount, maxElementCount-minElementCount)
}

func makePairCountMap(str string) map[string]uint64 {
	m := make(map[string]uint64)
	for x, _ := range str[:len(str)-1] {
		pair := string([]byte{str[x], str[x+1]})
		m[pair]++
	}
	return m
}

func runPolymerSteps(pairMapCount map[string]uint64, inserts map[string]byte, numSteps int) map[string]uint64 {
	for i := 1; i <= numSteps; i++ {
		fmt.Printf("starting step %d\n", i)
		fmt.Println(pairMapCount)
		pairMapCount = runPolymerStep(pairMapCount, inserts)
	}
	return pairMapCount
}

func runPolymerStep(pairMapCount map[string]uint64, inserts map[string]byte) map[string]uint64 {
	newPairMapCount := make(map[string]uint64)
	for str, count := range pairMapCount {
		if val, ok := inserts[str]; ok {
			new1 := []byte{str[0], val}
			new2 := []byte{val, str[1]}
			newPairMapCount[string(new1)] += count
			newPairMapCount[string(new2)] += count
		} else {
			newPairMapCount[str] += count
		}
	}
	return newPairMapCount
}

func findMaxAndMinLetterCount(pairMapCount map[string]uint64, lastLetter byte) (uint64, uint64) {
	charMap := make(map[byte]uint64)

	for str, count := range pairMapCount {
		charMap[str[0]] += count
	}
	charMap[lastLetter]++

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
