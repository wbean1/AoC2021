package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wbean1/AoC/day1"
	"github.com/wbean1/AoC/day10"
	"github.com/wbean1/AoC/day11"
	"github.com/wbean1/AoC/day12"
	"github.com/wbean1/AoC/day13"
	"github.com/wbean1/AoC/day14"
	"github.com/wbean1/AoC/day15"
	"github.com/wbean1/AoC/day16"
	"github.com/wbean1/AoC/day17"
	"github.com/wbean1/AoC/day2"
	"github.com/wbean1/AoC/day3"
	"github.com/wbean1/AoC/day4"
	"github.com/wbean1/AoC/day5"
	"github.com/wbean1/AoC/day6"
	"github.com/wbean1/AoC/day7"
	"github.com/wbean1/AoC/day8"
	"github.com/wbean1/AoC/day9"
)

func getDays() map[string]func() {
	days := make(map[string]func())
	days["day1"] = func() { day1.Run() }
	days["day2"] = func() { day2.Run() }
	days["day3"] = func() { day3.Run() }
	days["day4"] = func() { day4.Run() }
	days["day5"] = func() { day5.Run() }
	days["day6"] = func() { day6.Run() }
	days["day7"] = func() { day7.Run() }
	days["day8"] = func() { day8.Run() }
	days["day9"] = func() { day9.Run() }
	days["day10"] = func() { day10.Run() }
	days["day11"] = func() { day11.Run() }
	days["day12"] = func() { day12.Run() }
	days["day13"] = func() { day13.Run() }
	days["day14"] = func() { day14.Run() }
	days["day15"] = func() { day15.Run() }
	days["day16"] = func() { day16.Run() }
	days["day17"] = func() { day17.Run() }
	return days
}

func main() {
	dayMap := getDays()
	fmt.Println("Available days:")
	for k := range dayMap {
		fmt.Println(k)
	}
	for {
		fmt.Println("Which day to run?:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if myFunc, ok := dayMap[text]; ok {
			myFunc()
			os.Exit(0)
		}
	}
}
