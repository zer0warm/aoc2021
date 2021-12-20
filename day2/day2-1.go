package day2

import (
	"adventofcode/utils"
)

func FindEndOfPlan(instructions []utils.Input) int {
	h := 0
	d := 0
	for _, inst := range instructions {
		switch inst.S {
		case "forward":
			h += inst.I
		case "up":
			d -= inst.I
		case "down":
			d += inst.I
		}
	}
	return h * d
}
