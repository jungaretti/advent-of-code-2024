package solvers

import (
	"sort"
	"strconv"
	"strings"
)

type Day1Solver struct{}

func (d Day1Solver) Day1(input []string) (string, error) {
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

func (d Day1Solver) Day2(input []string) (string, error) {
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
