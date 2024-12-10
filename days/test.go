package days

import (
	"testing"
)

func TestDayPart(t *testing.T, solver Solver, part int, input []string, expected string) {
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
		t.Fatalf("could not solve day %d part %d: %w", err)
	}
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
