package day03

import (
	"advent-of-code-2024/days"
	"testing"
)

func TestDay03Part1(t *testing.T) {
	days.TestDayPart(t, Day03Solver{}, 1, []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}, "161")
}

func TestDay03Part2(t *testing.T) {
	days.TestDayPart(t, Day03Solver{}, 2, []string{}, "")
}
