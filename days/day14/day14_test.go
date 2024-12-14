package day14

import (
	"advent-of-code-2024/days"
	"testing"
)

func TestDay14Part1(t *testing.T) {
	days.TestDayPart(t, Day14Solver{}, 1, []string{}, "")
}

func TestDay14Part2(t *testing.T) {
	days.TestDayPart(t, Day14Solver{}, 2, []string{}, "")
}

func TestParsePositions(t *testing.T) {
	positions, err := ParsePositions([]string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(positions) != 12 {
		t.Fatalf("unexpected number of positions: %d", len(positions))
	}

	if positions[0] != (position{0, 4, 3, -3}) {
		t.Fatalf("unexpected position: %v", positions[0])
	}
	if positions[1] != (position{6, 3, -1, -3}) {
		t.Fatalf("unexpected position: %v", positions[1])
	}
	if positions[2] != (position{10, 3, -1, 2}) {
		t.Fatalf("unexpected position: %v", positions[2])
	}
}

func TestAdvancePositions(t *testing.T) {
	positions, err := ParsePositions([]string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = AdvancePositions(positions, 7, 11)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
