package solvers_test

import (
	"advent-of-code-2024/solvers"
	"testing"
)

var day1Solver = solvers.Day1Solver{}

func TestDay1Part1(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	expected := "11"

	testExampleInput(t, day1Solver, 1, input, expected)
}

func TestDay1Part2(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	expected := "31"

	testExampleInput(t, day1Solver, 2, input, expected)
}
