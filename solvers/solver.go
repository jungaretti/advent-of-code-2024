package solvers

type Solver interface {
	Day1(input []string) (string, error)
	Day2(input []string) (string, error)
}
