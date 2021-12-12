package day12

import (
	"fmt"
	"strings"

	"github.com/wbean1/AoC/utils"
)

type Node struct {
	Value string
	Paths []*Node
}

func Run() {
	start := Input()
	paths := start.FindPaths(start.Value)
	fmt.Printf("part1: there are %d paths found\n", len(paths))
	paths = start.FindPathsPartTwo(start.Value, 0)
	fmt.Printf("part2: there are %d paths found\n", len(paths))
}

func (n *Node) FindPaths(pathSoFar string) []string {
	paths := []string{}
	for _, node := range n.Paths {
		if node.Value == "end" {
			// we've reached the end of the path, no recurse, add to return
			paths = append(paths, pathSoFar+node.Value)
		} else if node.Value == strings.ToLower(node.Value) && !strings.Contains(pathSoFar, node.Value) {
			// is new lowercase... so add to paths and recurse
			paths = append(paths, node.FindPaths(pathSoFar+node.Value)...)
		} else if node.Value == strings.ToUpper(node.Value) {
			// is uppercase... so add to paths and recurse
			paths = append(paths, node.FindPaths(pathSoFar+node.Value)...)
		}
	}
	return paths
}

func (n *Node) FindPathsPartTwo(pathSoFar string, lcSoFar int) []string {
	paths := []string{}
	for _, node := range n.Paths {
		if node.Value == "start" {
			// can't go back here
			continue
		} else if node.Value == "end" {
			// we've reached the end of the path, no recurse, add to return
			paths = append(paths, pathSoFar+node.Value)
		} else if node.Value == strings.ToLower(node.Value) && !strings.Contains(pathSoFar, node.Value) {
			// is new lowercase... so add to paths and recurse
			paths = append(paths, node.FindPathsPartTwo(pathSoFar+node.Value, lcSoFar)...)
		} else if node.Value == strings.ToLower(node.Value) && strings.Contains(pathSoFar, node.Value) && lcSoFar == 0 {
			// is NOT-NEW lowercase, but we can do this once... so add to paths and recurse
			paths = append(paths, node.FindPathsPartTwo(pathSoFar+node.Value, lcSoFar+1)...)
		} else if node.Value == strings.ToUpper(node.Value) {
			// is uppercase... so add to paths and recurse
			paths = append(paths, node.FindPathsPartTwo(pathSoFar+node.Value, lcSoFar)...)
		}
	}
	return paths
}

func Input() *Node {
	input := utils.ParseFileToStrings("/Users/wbean/AoC2021/day12/input.txt")
	startNode := &Node{Value: "start"}
	toTryAgain := input
	for len(toTryAgain) > 0 {
		failures := []string{}
		for _, str := range toTryAgain {
			parts := strings.Split(str, "-")
			ok := startNode.AddNodePath(parts[0], parts[1])
			if !ok {
				failures = append(failures, str)
			}
		}
		toTryAgain = failures
	}
	return startNode
}

func (n *Node) AddNodePath(part1, part2 string) bool {
	node1, found1 := n.FindNode(part1, "")
	node2, found2 := n.FindNode(part2, "")
	if found1 && found2 {
		node1.AddPath(node2)
		node2.AddPath(node1)
		return true
	} else if found1 && !found2 {
		node2 := &Node{Value: part2}
		node1.AddPath(node2)
		node2.AddPath(node1)
		return true
	} else if !found1 && found2 {
		node1 := &Node{Value: part1}
		node2.AddPath(node1)
		node1.AddPath(node2)
		return true
	}
	return false
}

func (n *Node) FindNode(str string, alreadySeen string) (*Node, bool) {
	if n.Value == str {
		return n, true
	} else {
		for _, i := range n.Paths {
			if !strings.Contains(alreadySeen, i.Value) {
				k, found := i.FindNode(str, alreadySeen+n.Value)
				if found {
					return k, found
				}
			}
		}
	}
	return nil, false
}

func (n *Node) AddPath(another *Node) {
	n.Paths = append(n.Paths, another)
}
