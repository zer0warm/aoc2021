package day2

import (
	"testing"

	"adventofcode/utils"
)

func TestFindNewEndOfPlan(t *testing.T) {
	instructions := []utils.Input{}
	utils.ReadFileByLines("testdata/input", &instructions, ParseInput)

	want := 900
	got := FindNewEndOfPlan(instructions)

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}
