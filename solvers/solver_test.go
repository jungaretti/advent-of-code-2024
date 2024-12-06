package solvers_test

import (
	"advent-of-code-2024/solvers"
	"testing"
)

func testExampleInput(t *testing.T, solver solvers.Day1Solver, part int, input []string, expected string) {
	switch part {
	case 1:
		result, err := solver.Part1(input)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if result != expected {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	case 2:
		result, err := solver.Part2(input)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if result != expected {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	default:
		t.Fatalf("invalid part: %v", part)
	}
}
