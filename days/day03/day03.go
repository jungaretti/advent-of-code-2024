package day03

import (
	"advent-of-code-2024/days"
	"regexp"
	"strconv"
)

type Day03Solver struct{}

var _ days.Solver = Day03Solver{}

func (d Day03Solver) SolvePart1(input []string) (string, error) {
	instructionRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	instructions := []string{}
	for _, line := range input {
		matches := instructionRegex.FindAllString(line, -1)
		for _, match := range matches {
			instructions = append(instructions, match)
		}
	}

	result := 0
	for _, instruction := range instructions {
		matches := instructionRegex.FindStringSubmatch(instruction)
		left, err := strconv.Atoi(matches[1])
		if err != nil {
			return "", err
		}
		right, err := strconv.Atoi(matches[2])
		if err != nil {
			return "", err
		}

		result += (left * right)
	}

	return days.NumericResult(result, nil)
}

func (d Day03Solver) SolvePart2(input []string) (string, error) {
	tokenRegex := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(don\'t)|(do)`)
	instructionRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	instructions := []string{}
	instructionsEnabled := true
	for _, line := range input {
		tokens := tokenRegex.FindAllString(line, -1)
		for _, token := range tokens {
			switch token {
			case "don't":
				instructionsEnabled = false
			case "do":
				instructionsEnabled = true
			default:
				if instructionsEnabled {
					instructions = append(instructions, token)
				}
			}
		}
	}

	result := 0
	for _, instruction := range instructions {
		matches := instructionRegex.FindStringSubmatch(instruction)
		left, err := strconv.Atoi(matches[1])
		if err != nil {
			return "", err
		}
		right, err := strconv.Atoi(matches[2])
		if err != nil {
			return "", err
		}

		result += (left * right)
	}

	return days.NumericResult(result, nil)
}
