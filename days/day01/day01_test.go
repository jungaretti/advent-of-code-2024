package day01

import (
	"advent-of-code-2024/days"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	days.TestDayPart(t, Day01Solver{}, 1, []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}, "11")
}

func TestDay1Part2(t *testing.T) {
	days.TestDayPart(t, Day01Solver{}, 2, []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}, "31")
}
