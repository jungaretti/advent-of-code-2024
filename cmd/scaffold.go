package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func newScaffoldCmd() *cobra.Command {
	root := &cobra.Command{
		Use:  "scaffold <day>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			day, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid day: %v", err)
			}
			if day < 1 || day > 25 {
				return fmt.Errorf("invalid day: %d", day)
			}

			err = scaffold(day)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return root
}

func scaffold(day int) error {
	dayPadded := fmt.Sprintf("%02d", day)

	dirName := fmt.Sprintf("./days/day%v", dayPadded)
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	solverFileName := fmt.Sprintf("%v/day%v.go", dirName, dayPadded)
	solverFileContent := fmt.Sprintf(`package day%v

import (
	"advent-of-code-2024/days"
)

type Day%vSolver struct{}

var _ days.Solver = Day%vSolver{}

func (d Day%vSolver) SolvePart1(input []string) (string, error) {
	return "", nil
}

func (d Day%vSolver) SolvePart2(input []string) (string, error) {
	return "", nil
}`, dayPadded, dayPadded, dayPadded, dayPadded, dayPadded)
	err = os.WriteFile(solverFileName, []byte(solverFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write solver file: %w", err)
	}

	solverTestFileName := fmt.Sprintf("%v/day%v_test.go", dirName, dayPadded)
	solverTestFileContent := fmt.Sprintf(`package day%v

import (
	"advent-of-code-2024/days"
	"testing"
)

func TestDay%vPart1(t *testing.T) {
	days.TestDayPart(t, Day%vSolver{}, 1, []string{
	}, "")
}

func TestDay%vPart2(t *testing.T) {
	days.TestDayPart(t, Day%vSolver{}, 2, []string{
	}, "")
}`, dayPadded, dayPadded, dayPadded, dayPadded, dayPadded)
	err = os.WriteFile(solverTestFileName, []byte(solverTestFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write solver test file: %w", err)
	}

	return nil
}
