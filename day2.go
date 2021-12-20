package main

import (
	"fmt"

	"adventofcode/day2"
	"adventofcode/utils"
)

func main() {
	instructions := []utils.Input{}
	utils.ReadFileByLines("day2/input", &instructions, day2.ParseInput)
	
	firstResult := day2.FindEndOfPlan(instructions)
	fmt.Print("1. final horizontal position x final depth = ")
	fmt.Println(firstResult)

	secondResult := day2.FindNewEndOfPlan(instructions)
	fmt.Print("2. (newer) final horizontal position x final depth = ")
	fmt.Println(secondResult)
}
