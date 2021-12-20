package day1

import (
	"adventofcode/utils"
	"testing"
)

func TestCountIncreasingTripletSumDepths(t *testing.T) {
	depths := []utils.Input{}
	utils.ReadFileByLines("testdata/input", &depths, ParseInput)

	want := 5
	got := CountIncreasingTripletSumDepths(depths)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
