package day18

import (
	"fmt"
	"log"
	"strconv"

	"github.com/wbean1/AoC/utils"
)

type SFN struct {
	X, Y, Parent *SFN
	Value        int
	Depth        int
}

func Add(sfn1, sfn2 SFN) SFN {
	new := SFN{X: &sfn1, Y: &sfn2, Depth: 0}
	for new.NeedsExploded(0) || new.NeedsSplit() {
		new = toSFN(new.String(), 0, nil)
		if new.NeedsExploded(0) {
			fmt.Printf("before explode: %s\n", new.String())
			new.Explode()
			fmt.Printf("after explode: %s\n", new.String())
		} else {
			fmt.Printf("before split: %s\n", new.String())
			new.Split()
			fmt.Printf("after split: %s\n", new.String())
		}
	}
	return new
}

func (sfn *SFN) findSFNtoSplit() (*SFN, bool) {
	if sfn.Value > 9 && sfn.X == nil && sfn.Y == nil {
		return sfn, true
	} else {
		var foundSfnLeft, foundSfnRight *SFN
		var leftFound, rightFound bool
		if sfn.X != nil {
			foundSfnLeft, leftFound = sfn.X.findSFNtoSplit()
		}
		if sfn.Y != nil {
			foundSfnRight, rightFound = sfn.Y.findSFNtoSplit()
		}
		if leftFound {
			return foundSfnLeft, leftFound
		} else if rightFound {
			return foundSfnRight, rightFound
		} else {
			return nil, false
		}
	}
}

func (sfn *SFN) Split() {
	// find where to split...
	childToSplit, found := sfn.findSFNtoSplit()
	fmt.Printf("found sfn to split: %s\n", childToSplit)
	if !found {
		log.Fatal("didn't find sfn to split!")
	}
	valueToSplit := childToSplit.Value
	var leftValue, rightValue int
	if valueToSplit%2 == 0 {
		leftValue = valueToSplit / 2
		rightValue = valueToSplit / 2
	} else {
		leftValue = valueToSplit / 2
		rightValue = (valueToSplit / 2) + 1
	}
	childToSplit.X = &SFN{Value: leftValue, Parent: childToSplit}
	childToSplit.Y = &SFN{Value: rightValue, Parent: childToSplit}
	if childToSplit.Parent != nil {
		childToSplit.Depth = childToSplit.Parent.Depth + 1
	} else {
		childToSplit.Depth = 0
	}
}

func (sfn *SFN) Explode() {
	// find the SFN with depth == 4
	depth4sfn, found := sfn.Find(4)
	fmt.Printf("depth4sfn is %s\n", depth4sfn.String())
	if !found {
		log.Fatalf("didn't find depth4 on explode()")
	}
	leftVal := depth4sfn.X.Value
	rightVal := depth4sfn.Y.Value
	depth4sfn.X = nil
	depth4sfn.Y = nil
	depth4sfn.Value = 0
	depth4sfn.AddToLeft(leftVal)
	depth4sfn.AddToRight(rightVal)
}

func (sfn *SFN) String() string {
	if sfn.X != nil && sfn.Y != nil {
		return fmt.Sprintf("[%s,%s]", sfn.X.String(), sfn.Y.String())
	} else {
		return fmt.Sprintf("%d", sfn.Value)
	}
}

func (sfn *SFN) AddToLeft(leftVal int) {
	movedLeft := false
	child := sfn
	var current *SFN
	if sfn.Parent != nil {
		current = sfn.Parent
	} else {
		return
	}

	for {
		fmt.Printf("looking at: %s\n", current.String())
		if !movedLeft {
			if current.X != nil && current.X != child {
				movedLeft = true
				current = current.X
				fmt.Println("moved left!")
			} else if current.Parent != nil {
				child = current
				current = current.Parent
				fmt.Println("moved up!")
			} else {
				return
			}
		} else if current.Y != nil && *current.Y != *child {
			current = current.Y
			fmt.Println("looking right!")
		} else if current.X != nil {
			current = current.X
			fmt.Println("looking left!")
		} else {
			fmt.Printf("found value: %d\n", current.Value)
			current.Value += leftVal

			return
		}
	}
}

func (sfn *SFN) AddToRight(rightVal int) {
	movedRight := false
	child := sfn
	var current *SFN
	if sfn.Parent != nil {
		current = sfn.Parent
	} else {
		return
	}
	for {
		if !movedRight {
			if current.Y != nil && *current.Y != *child {
				movedRight = true
				current = current.Y
			} else if current.Parent != nil {
				child = current
				current = current.Parent
			} else {
				return
			}
		} else if current.X != nil && *current.X != *child {
			current = current.X
		} else if current.Y != nil {
			current = current.Y
		} else {
			current.Value += rightVal
			return
		}
	}
}

func (sfn *SFN) Find(depth int) (*SFN, bool) {
	if sfn.Depth == depth && sfn.X != nil && sfn.Y != nil {
		return sfn, true
	} else {
		var leftSFN, rightSFN *SFN
		var foundLeft, foundRight bool
		if sfn.X != nil {
			leftSFN, foundLeft = sfn.X.Find(depth)
		}
		if sfn.Y != nil {
			rightSFN, foundRight = sfn.Y.Find(depth)
		}
		if foundLeft {
			return leftSFN, true
		}
		if foundRight {
			return rightSFN, true
		}
	}
	return nil, false
}

func (sfn *SFN) NeedsExploded(depth int) bool {
	sfn.Depth = depth
	if depth > 4 {
		return true
	} else {
		var leftNeeds, rightNeeds bool
		if sfn.X != nil {
			leftNeeds = sfn.X.NeedsExploded(depth + 1)
		}
		if sfn.Y != nil {
			rightNeeds = sfn.Y.NeedsExploded(depth + 1)
		}
		return leftNeeds || rightNeeds
	}
}

func (sfn *SFN) NeedsSplit() bool {
	if sfn.Value > 9 && sfn.X == nil && sfn.Y == nil {
		return true
	} else {
		var leftNeeds, rightNeeds bool
		if sfn.X != nil {
			leftNeeds = sfn.X.NeedsSplit()
		}
		if sfn.Y != nil {
			rightNeeds = sfn.Y.NeedsSplit()
		}
		return leftNeeds || rightNeeds
	}
}

func Run() {
	input := Input("/Users/wbean/AoC2021/day18/input.txt")
	sfn := SFN(input[0])
	for i, line := range input[1:] {
		sfn2 := SFN(line)
		sfn = Add(sfn, sfn2)
		fmt.Printf("==> %d: %s\n", i, sfn.String())
	}
	fmt.Printf("part1: magnitude is: %d\n", sfn.Magnitude())
}

func Input(file string) []SFN {
	sfns := []SFN{}
	input := utils.ParseFileToStrings(file)
	for _, line := range input {
		sfns = append(sfns, toSFN(line, 0, nil))
	}
	return sfns
}

func toSFN(s string, depth int, parent *SFN) SFN {
	creating := &SFN{}
	if s[0] == byte('[') &&
		s[len(s)-1] == byte(']') {
		s = s[1 : len(s)-1] // remove first and last char
		// find 0 depth ','
		localDepth := 0
		var leftStr, rightStr string
		for i, char := range s {
			if char == rune('[') {
				localDepth++
			} else if char == rune(']') {
				localDepth--
			} else if char == rune(',') && localDepth == 0 {
				leftStr = s[:i]
				rightStr = s[i+1:]
			}
		}
		left := toSFN(leftStr, depth+1, creating)
		right := toSFN(rightStr, depth+1, creating)
		creating.X = &left
		creating.Y = &right
		creating.Parent = parent
		creating.Depth = depth
		return *creating
	} else {
		value, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return SFN{Value: value, Depth: depth, Parent: parent}
	}
}

func (sfn *SFN) Magnitude() int {
	var x, y int
	if sfn.X != nil {
		x = 3 * sfn.X.Magnitude()
	}
	if sfn.Y != nil {
		y = 2 * sfn.Y.Magnitude()
	}
	if sfn.X == nil && sfn.Y == nil {
		return sfn.Value
	}
	return x + y
}
