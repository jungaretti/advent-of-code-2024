package day02

import (
	"advent-of-code-2024/days"
	"strconv"
	"strings"
)

type Day02Solver struct{}

var _ days.Solver = Day02Solver{}

func (d Day02Solver) SolvePart1(input []string) (string, error) {
	reports, err := parseReports(input)
	if err != nil {
		return "", err
	}

	validReports := 0
	for _, report := range reports {
		if (allIncreasing(report) || allDecreasing(report)) && differsWithinBounds(report, 1, 3) {
			validReports++
		}
	}

	return strconv.Itoa(validReports), nil
}

func (d Day02Solver) SolvePart2(input []string) (string, error) {
	return "", nil
}

func parseReports(input []string) ([][]int, error) {
	var reports [][]int
	for _, line := range input {
		var currentReport []int
		for _, level := range strings.Fields(line) {
			value, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			currentReport = append(currentReport, value)
		}
		reports = append(reports, currentReport)
	}

	return reports, nil
}

func allIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] < report[i-1] {
			return false
		}
	}
	return true
}

func allDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			return false
		}
	}
	return true
}

func differsWithinBounds(report []int, atLeast int, atMost int) bool {
	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]
		if difference < 0 {
			difference = -difference
		}

		if difference < atLeast || difference > atMost {
			return false
		}
	}

	return true
}
