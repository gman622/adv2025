package day1

// Part1 solves part 1: count how many times the dial ends at position 0
func Part1(inputPath string) (int, error) {
	return SolveWith(EndPositionCounter{}, inputPath)
}
