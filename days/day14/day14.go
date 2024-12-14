package day14

import (
	"advent-of-code-2024/days"
	"fmt"
	"regexp"
	"strconv"
)

type Day14Solver struct{}

var _ days.Solver = Day14Solver{}

func (d Day14Solver) SolvePart1(input []string) (string, error) {
	return "", nil
}

func (d Day14Solver) SolvePart2(input []string) (string, error) {
	return "", nil
}

type position = struct {
	x, y   int
	dx, dy int
}

func ParsePositions(input []string) ([]position, error) {
	var positionRegex = regexp.MustCompile(`p=([-\d]+),([-\d]+)\s+v=([-\d]+),([-\d]+)`)

	positions := make([]position, 0, len(input))
	for _, line := range input {
		matches := positionRegex.FindStringSubmatch(line)
		if len(matches) != 5 {
			return nil, fmt.Errorf("invalid input: %s", line)
		}

		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		dx, _ := strconv.Atoi(matches[3])
		dy, _ := strconv.Atoi(matches[4])

		newPosition := position{x, y, dx, dy}
		positions = append(positions, newPosition)
	}

	return positions, nil
}

func AdvancePositions(positions []position, height, width int) ([]position, error) {
	newPositions := make([]position, 0, len(positions))
	for _, pos := range positions {
		newPosition := position{pos.x + pos.dx, pos.y + pos.dy, pos.dx, pos.dy}
		newPositions = append(newPositions, newPosition)
	}

	return newPositions, nil
}
