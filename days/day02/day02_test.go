package day02

import (
	"advent-of-code-2024/days"
	"testing"
)

func TestDay02Part1(t *testing.T) {
	days.TestDayPart(t, Day02Solver{}, 1, []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}, "2")
}

func TestDay02Part2(t *testing.T) {
	days.TestDayPart(t, Day02Solver{}, 2, []string{}, "")
}
