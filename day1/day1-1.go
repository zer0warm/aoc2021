package day1

import (
	"adventofcode/utils"
)

func CountIncreasingDepths(depths []utils.Input) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i].I > depths[i-1].I {
			count += 1
		}
	}
	return count
}
