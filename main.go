package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wbean1/AoC/day1"
	"github.com/wbean1/AoC/day2"
	"github.com/wbean1/AoC/day3"
	"github.com/wbean1/AoC/day4"
	"github.com/wbean1/AoC/day5"
	"github.com/wbean1/AoC/day6"
)

func getDays() map[string]func() {
	days := make(map[string]func())
	days["day1"] = func() { day1.Run() }
	days["day2"] = func() { day2.Run() }
	days["day3"] = func() { day3.Run() }
	days["day4"] = func() { day4.Run() }
	days["day5"] = func() { day5.Run() }
	days["day6"] = func() { day6.Run() }
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
