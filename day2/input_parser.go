package day2

import (
	"strconv"
	"strings"

	"adventofcode/utils"
)

func ParseInput(line string) utils.Input {
	compound := strings.Split(line, " ")
	command := compound[0]
	parameter, _ := strconv.ParseInt(compound[1], 10, 0)
	return utils.Input{S: command, I: int(parameter)}
}
