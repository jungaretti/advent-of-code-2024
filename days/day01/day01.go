package day01

import (
	"advent-of-code-2024/days"
	"sort"
	"strconv"
	"strings"
)

type Day01Solver struct{}

var _ days.Solver = Day01Solver{}

func (d Day01Solver) SolvePart1(input []string) (string, error) {
	leftValues := []int{}
	rightValues := []int{}

	for _, line := range input {
		parts := strings.Fields(line)
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", err
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}

		leftValues = append(leftValues, left)
		rightValues = append(rightValues, right)
	}

	sort.Ints(leftValues)
	sort.Ints(rightValues)

	result := 0
	for i := 0; i < len(leftValues); i++ {
		left := leftValues[i]
		right := rightValues[i]

		if left < right {
			result += (right - left)
		} else {
			result += (left - right)
		}
	}

	return strconv.Itoa(result), nil
}

func (d Day01Solver) SolvePart2(input []string) (string, error) {
	rightCount := map[int]int{}

	for _, line := range input {
		parts := strings.Fields(line)
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}

		rightCount[right]++
	}

	result := 0
	for _, line := range input {
		parts := strings.Fields(line)
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", err
		}

		rightCount := rightCount[left]
		result += (left * rightCount)
	}

	return strconv.Itoa(result), nil
}
