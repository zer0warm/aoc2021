package main

import (
	"adventofcode/day1"
	"adventofcode/utils"
	"fmt"
)

func main() {
	depths := []utils.Input{}
	utils.ReadFileByLines("day1/input", &depths, day1.ParseInput)

	firstResult := day1.CountIncreasingDepths(depths)
	fmt.Print("1. Number of measurements larger than previous measurement: ")
	fmt.Println(firstResult)

	secondResult := day1.CountIncreasingTripletSumDepths(depths)
	fmt.Print("2. Number of sums larger than the previous sum: ")
	fmt.Println(secondResult)
}
