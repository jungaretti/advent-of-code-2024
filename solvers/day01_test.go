package solvers_test

import (
	"advent-of-code-2024/solvers"
	"testing"
)

func TestDay1Solver_Day2(t *testing.T) {
	solver := solvers.Day1Solver{}
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	expected := "31"
	result, err := solver.Part2(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
