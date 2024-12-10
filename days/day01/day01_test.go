package day01_test

import (
	"advent-of-code-2024/days/day01"
	"testing"
)

var day1Solver = day01.Day01Solver{}

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

func testExampleInput(t *testing.T, solver day01.Day01Solver, part int, input []string, expected string) {
	var (
		result string
		err    error
	)

	switch part {
	case 1:
		result, err = solver.SolvePart1(input)
	case 2:
		result, err = solver.SolvePart2(input)
	default:
		t.Fatalf("invalid part: %v", part)
	}

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
