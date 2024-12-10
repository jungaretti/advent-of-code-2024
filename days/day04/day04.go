package day04

import (
	"advent-of-code-2024/days"
)

type Day04Solver struct{}

var _ days.Solver = Day04Solver{}

func (d Day04Solver) SolvePart1(input []string) (string, error) {
	var directions = []struct{ dx, dy int }{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // diagonal up-left
		{-1, 1},  // diagonal up-right
		{1, -1},  // diagonal down-left
		{1, 1},   // diagonal down-right
	}

	result := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			for _, dir := range directions {
				if matchWord(input, "XMAS", i, j, dir.dx, dir.dy) {
					result++
				}
			}
		}
	}

	return days.NumericResult(result, nil)
}

func (d Day04Solver) SolvePart2(input []string) (string, error) {
	return "", nil
}

func matchWord(input []string, word string, x, y, dx, dy int) bool {
	for i := 0; i < len(word); i++ {
		if x < 0 || x >= len(input) || y < 0 || y >= len(input[0]) {
			return false
		}
		if input[x][y] != word[i] {
			return false
		}

		x += dx
		y += dy
	}
	return true
}
