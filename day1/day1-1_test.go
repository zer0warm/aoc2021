package day1

import (
	"adventofcode/utils"
	"testing"
)

func TestCountIncreasingDepths(t *testing.T) {
	depths := []utils.Input{}
	utils.ReadFileByLines("testdata/input", &depths, ParseInput)

	want := 7
	got := CountIncreasingDepths(depths)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
