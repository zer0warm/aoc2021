package day2

import (
	"testing"

	"adventofcode/utils"
)

func TestFindEndOfPlan(t *testing.T) {
	instructions := []utils.Input{}
	utils.ReadFileByLines("testdata/input", &instructions, ParseInput)

	want := 150
	got := FindEndOfPlan(instructions)

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}
