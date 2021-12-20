package day1

import (
	"adventofcode/utils"
)

func CountIncreasingTripletSumDepths(depths []utils.Input) int {
	count := 0
	tripletSum := depths[0].I + depths[1].I + depths[2].I
	for i := 3; i < len(depths); i++ {
		nextTripletSum := tripletSum - depths[i-3].I + depths[i].I
		if nextTripletSum > tripletSum {
			count += 1
		}
	}
	return count
}
