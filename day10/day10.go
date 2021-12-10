package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Stack []rune

const opens = "({[<"

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str rune) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(' '), false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func Run() {
	input := Input()
	incorrects := []rune{}
	incompletes := []string{}
	for _, str := range input {
		incorrect, hasIncorrect := findIncorrect(str)
		if hasIncorrect {
			incorrects = append(incorrects, incorrect)
		} else {
			incompletes = append(incompletes, findIncomplete(str))
		}
	}
	score := scoreIncorrects(incorrects)
	fmt.Println(incorrects)
	fmt.Printf("part1: syntax error score: %d\n", score)
	score = scoreIncompletes(incompletes)
	fmt.Println(incompletes)
	fmt.Printf("part2: incompletes middle score: %d\n", score)
}

func findIncomplete(str string) string {
	correspondingClose := map[rune]rune{
		rune('('): rune(')'),
		rune('['): rune(']'),
		rune('{'): rune('}'),
		rune('<'): rune('>'),
	}
	var openStack Stack
	for _, char := range str {
		if strings.Contains(opens, string(char)) {
			openStack.Push(char)
		} else {
			_, ok := openStack.Pop()
			if !ok {
				log.Fatal("tried to pop on empty stack, oh noes")
			}
			// just going to assume its correct here...
		}
	}
	closesNeeded := []rune{}
	for _, char := range openStack {
		closesNeeded = append(closesNeeded, correspondingClose[char])
	}
	return string(closesNeeded)
}

func scoreIncompletes(incompletes []string) int {
	incompletesScores := []int{}
	for _, incomplete := range incompletes {
		incompletesScores = append(incompletesScores, scoreIncomplete(incomplete))
	}
	sort.Ints(incompletesScores)
	return incompletesScores[len(incompletesScores)/2]
}

func scoreIncomplete(incomplete string) int {
	charScores := map[rune]int{
		rune(')'): 1,
		rune(']'): 2,
		rune('}'): 3,
		rune('>'): 4,
	}
	score := 0
	for _, char := range incomplete {
		score *= 5
		score += charScores[char]
	}
	return score
}

func findIncorrect(str string) (rune, bool) {
	correspondingClose := map[rune]rune{
		rune('('): rune(')'),
		rune('['): rune(']'),
		rune('{'): rune('}'),
		rune('<'): rune('>'),
	}
	var openStack Stack
	for _, char := range str {
		if strings.Contains(opens, string(char)) {
			openStack.Push(char)
		} else {
			need, ok := openStack.Pop()
			if !ok {
				log.Fatal("tried to pop on empty stack, oh noes")
			}
			need = correspondingClose[need]
			if char != need {
				return char, true
			}
		}
	}
	return ' ', false
}

func scoreIncorrects(chars []rune) int {
	score := 0
	charScores := map[rune]int{
		rune(')'): 3,
		rune(']'): 57,
		rune('}'): 1197,
		rune('>'): 25137,
	}
	for _, char := range chars {
		score += charScores[char]
	}
	return score
}

func Input() []string {
	return parseFile("/Users/wbean/AoC2021/day10/input.txt")
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
