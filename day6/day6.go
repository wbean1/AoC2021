package day6

import "fmt"

type Laternfish struct {
	timer int
}

func Input() map[int]uint64 {
	i := []int{1, 1, 3, 5, 3, 1, 1, 4, 1, 1, 5, 2, 4, 3, 1, 1, 3, 1, 1, 5, 5, 1, 3, 2, 5, 4, 1, 1, 5, 1, 4, 2, 1, 4, 2, 1, 4, 4, 1, 5, 1, 4, 4, 1, 1, 5, 1, 5, 1, 5, 1, 1, 1, 5, 1, 2, 5, 1, 1, 3, 2, 2, 2, 1, 4, 1, 1, 2, 4, 1, 3, 1, 2, 1, 3, 5, 2, 3, 5, 1, 1, 4, 3, 3, 5, 1, 5, 3, 1, 2, 3, 4, 1, 1, 5, 4, 1, 3, 4, 4, 1, 2, 4, 4, 1, 1, 3, 5, 3, 1, 2, 2, 5, 1, 4, 1, 3, 3, 3, 3, 1, 1, 2, 1, 5, 3, 4, 5, 1, 5, 2, 5, 3, 2, 1, 4, 2, 1, 1, 1, 4, 1, 2, 1, 2, 2, 4, 5, 5, 5, 4, 1, 4, 1, 4, 2, 3, 2, 3, 1, 1, 2, 3, 1, 1, 1, 5, 2, 2, 5, 3, 1, 4, 1, 2, 1, 1, 5, 3, 1, 4, 5, 1, 4, 2, 1, 1, 5, 1, 5, 4, 1, 5, 5, 2, 3, 1, 3, 5, 1, 1, 1, 1, 3, 1, 1, 4, 1, 5, 2, 1, 1, 3, 5, 1, 1, 4, 2, 1, 2, 5, 2, 5, 1, 1, 1, 2, 3, 5, 5, 1, 4, 3, 2, 2, 3, 2, 1, 1, 4, 1, 3, 5, 2, 3, 1, 1, 5, 1, 3, 5, 1, 1, 5, 5, 3, 1, 3, 3, 1, 2, 3, 1, 5, 1, 3, 2, 1, 3, 1, 1, 2, 3, 5, 3, 5, 5, 4, 3, 1, 5, 1, 1, 2, 3, 2, 2, 1, 1, 2, 1, 4, 1, 2, 3, 3, 3, 1, 3, 5}
	m := make(map[int]uint64)
	for _, timer := range i {
		m[timer]++
	}
	return m
}

func Run() {
	fish := Input()
	for i := 1; i <= 80; i++ {
		fish = Next(fish)
	}
	fmt.Printf("part1: there are now %d fish\n", Sum(fish))
	fish = Input()
	for i := 1; i <= 256; i++ {
		fish = Next(fish)
	}
	fmt.Printf("part2: there are now %d fish\n", Sum(fish))
}

func Next(m map[int]uint64) map[int]uint64 {
	nextFish := make(map[int]uint64)
	for key, value := range m {
		if key == 0 {
			nextFish[6] += value
			nextFish[8] += value
		} else {
			nextFish[key-1] += value
		}
	}
	return nextFish
}

func Sum(m map[int]uint64) uint64 {
	var sum uint64
	for _, value := range m {
		sum += value
	}
	return sum
}
