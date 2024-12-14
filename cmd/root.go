package cmd

import (
	"advent-of-code-2024/days"
	"advent-of-code-2024/days/day01"
	"advent-of-code-2024/days/day02"
	"advent-of-code-2024/days/day03"
	"advent-of-code-2024/days/day04"
	"advent-of-code-2024/days/day14"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func Execute() {
	root := newRootCmd()

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:  "aoc2024 <day> <part>",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			day, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid day: %v", err)
			}
			if day < 1 || day > 25 {
				return fmt.Errorf("invalid day: %d", day)
			}

			part, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid part: %v", err)
			}
			if part < 1 || part > 2 {
				return fmt.Errorf("invalid part: %d", part)
			}

			result, err := root(day, part)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	root.AddCommand(newScaffoldCmd())

	return root
}

func root(day, part int) (string, error) {
	solvers := map[int]days.Solver{
		1:  day01.Day01Solver{},
		2:  day02.Day02Solver{},
		3:  day03.Day03Solver{},
		4:  day04.Day04Solver{},
		14: day14.Day14Solver{},
	}

	solver, exists := solvers[day]
	if !exists {
		return "", fmt.Errorf("no solver for day %d", day)
	}

	inputFile := fmt.Sprintf("inputs/day%02d.txt", day)
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %v", err)
	}
	lines := strings.Split(string(input), "\n")

	switch part {
	case 1:
		return solver.SolvePart1(lines)
	case 2:
		return solver.SolvePart2(lines)
	default:
		return "", fmt.Errorf("no solver for part %d", part)
	}
}
