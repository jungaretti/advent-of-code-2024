package day04

import (
	"advent-of-code-2024/days"
	"testing"
)

func TestDay04Part1(t *testing.T) {
	days.TestDayPart(t, Day04Solver{}, 1, []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}, "18")
}

func TestDay04Part2(t *testing.T) {
	days.TestDayPart(t, Day04Solver{}, 2, []string{}, "")
}
