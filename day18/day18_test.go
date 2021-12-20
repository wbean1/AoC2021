package day18

import (
	"fmt"
	"testing"
)

var NeedsExplodingTestCases = map[string]bool{
	"[[[[[9,8],1],2],3],4]":                                         true,
	"[[6,[5,[4,[3,2]]]],1]":                                         true,
	"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":                             true,
	"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]":                             false,
	"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]":                     false,
	"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]": false,
}

var NeedsSplitTestCases = map[string]bool{
	"[[[[0,7],4],[15,[0,13]]],[1,1]]":                               true,
	"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]":                            true,
	"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]":                         false,
	"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]":                             false,
	"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]":                     false,
	"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]": false,
}

var ExplodeTestCases = map[string]string{
	"[[[[[9,8],1],2],3],4]":                                   "[[[[0,9],2],3],4]",
	"[7,[6,[5,[4,[3,2]]]]]":                                   "[7,[6,[5,[7,0]]]]",
	"[[6,[5,[4,[3,2]]]],1]":                                   "[[6,[5,[7,0]]],3]",
	"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]":                   "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
	"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":                       "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	"[[[[0,8],[[9,9],0]],[[10,0],[4,8]]],[[[0,24],[0,7]],8]]": "[[[[0,17],[0,9]],[[10,0],[4,8]]],[[[0,24],[0,7]],8]]",
	"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]": "[[[[4,0],[5,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]",
}

func TestNeedsExploded(t *testing.T) {
	for input, expected := range NeedsExplodingTestCases {
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		got := inputSFN.NeedsExploded(0)
		if got != expected {
			t.Errorf("wrong value of NeedsExploded for %s.  got: %t, expected: %t", input, got, expected)
		} else {
			fmt.Printf("yay! NeedsExploded for %s.  got: %t, expected: %t\n", input, got, expected)
		}
	}
}

func TestNeedsSplit(t *testing.T) {
	for input, expected := range NeedsSplitTestCases {
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		got := inputSFN.NeedsSplit()
		if got != expected {
			t.Errorf("wrong value of NeedsSplit for %s.  got: %t, expected: %t", input, got, expected)
		} else {
			fmt.Printf("yay! NeedsSplit for %s.  got: %t, expected: %t\n", input, got, expected)
		}
	}
}

func TestString(t *testing.T) {
	for input, _ := range NeedsSplitTestCases {
		expected := input
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		got := inputSFN.String()
		if got != expected {
			t.Errorf("wrong value of String for %s.  got: %s, expected: %s", input, got, expected)
		} else {
			fmt.Printf("yay! String for %s.  got: %s, expected: %s\n", input, got, expected)
		}
	}
}

func TestExplode(t *testing.T) {
	for input, expected := range ExplodeTestCases {
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		inputSFN.Explode()
		got := inputSFN.String()
		if got != expected {
			t.Errorf("wrong tree after Explode for %s.  got: %s, expected: %s", input, got, expected)
		} else {
			fmt.Printf("yay! Explode for %s.  got: %s, expected: %s\n", input, got, expected)
		}
	}
}

var AddToLeftTestCases = map[string]string{
	"[[5,5],1]": "[[5,5],1]",
	"[1,[5,5]]": "[5,[5,5]]",
}

func TestAddToLeft(t *testing.T) {
	for input, expected := range AddToLeftTestCases {
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		where, _ := inputSFN.Find(1)
		where.AddToLeft(4)
		got := inputSFN.String()
		if got != expected {
			t.Errorf("wrong tree after AddToLeft for %s.  got: %s, expected: %s", input, got, expected)
		} else {
			fmt.Printf("yay! AddToLeft for %s.  got: %s, expected: %s\n", input, got, expected)
		}
	}

}

var AddToRightTestCases = map[string]string{
	"[[5,5],1]": "[[5,5],5]",
	"[1,[5,5]]": "[1,[5,5]]",
}

func TestAddToRight(t *testing.T) {
	for input, expected := range AddToRightTestCases {
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		where, _ := inputSFN.Find(1)
		where.AddToRight(4)
		got := inputSFN.String()
		if got != expected {
			t.Errorf("wrong tree after AddToRight for %s.  got: %s, expected: %s", input, got, expected)
		} else {
			fmt.Printf("yay! AddToRight for %s.  got: %s, expected: %s\n", input, got, expected)
		}
	}
}

var SplitTestCases = map[string]string{
	"[[[[0,7],4],[15,[0,13]]],[1,1]]":                         "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
	"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]":                      "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
	"[[[[0,0],[7,7]],[[0,7],[0,7]]],[[[0,8],[0,0]],[14,15]]]": "[[[[0,0],[7,7]],[[0,7],[0,7]]],[[[0,8],[0,0]],[[7,7],15]]]",
}

func TestSplit(t *testing.T) {
	for input, expected := range SplitTestCases {
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		inputSFN.Split()
		got := inputSFN.String()
		if got != expected {
			t.Errorf("wrong value after Split for %s.  got: %s, expected: %s", input, got, expected)
		} else {
			fmt.Printf("yay! Split for %s.  got: %s, expected: %s\n", input, got, expected)
		}
	}
}

var MagnitudeTestCases = map[string]int{
	"[[1,2],[[3,4],5]]":                                             143,
	"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                             1384,
	"[[[[1,1],[2,2]],[3,3]],[4,4]]":                                 445,
	"[[[[3,0],[5,3]],[4,4]],[5,5]]":                                 791,
	"[[[[5,0],[7,4]],[5,5]],[6,6]]":                                 1137,
	"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]":         3488,
	"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]": 4140,
}

func TestMagnitude(t *testing.T) {
	for input, expected := range MagnitudeTestCases {
		fmt.Printf("parsing string: %s\n", input)
		inputSFN := toSFN(input, 0, nil)
		got := inputSFN.Magnitude()
		if got != expected {
			t.Errorf("wrong Magnitude for %s.  got: %d, expected: %d", input, got, expected)
		} else {
			fmt.Printf("yay! Magnitude for %s.  got: %d, expected: %d\n", input, got, expected)
		}
	}
}

// var AddTestCases = map[string]string{
// 	"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]+[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]": "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
// }

// func TestAdd(t *testing.T) {
// 	for input, expected := range AddTestCases {
// 		sfns := strings.Split(input, "+")
// 		sfn1 := toSFN(sfns[0], 0, nil)
// 		sfn2 := toSFN(sfns[1], 0, nil)
// 		sfn := Add(sfn1, sfn2)
// 		got := sfn.String()
// 		if got != expected {
// 			t.Errorf("wrong Sum for %s.  got: %s, expected: %s", input, got, expected)
// 		} else {
// 			fmt.Printf("yay! Sum for %s.  got: %s, expected: %s\n", input, got, expected)
// 		}
// 	}
// }
