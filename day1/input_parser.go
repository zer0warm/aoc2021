package day1

import (
	"strconv"

	"adventofcode/utils"
)

func ParseInput(line string) utils.Input {
	number, _ := strconv.ParseInt(line, 10, 0)
	return utils.Input{I: int(number)}
}
