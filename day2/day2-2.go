package day2

import (
	"adventofcode/utils"
)

func FindNewEndOfPlan(instructions []utils.Input) int {
	h := 0
	d := 0
	a := 0
	for _, inst := range instructions {
		switch inst.S {
		case "forward":
			h += inst.I
			d += a * inst.I
		case "up":
			a -= inst.I
		case "down":
			a += inst.I
		}
	}
	return h * d
}
