package day1

// Part2 solves part 2: count how many times the dial passes through position 0
func Part2(inputPath string) (int, error) {
	return SolveWith(ZeroCrossingCounter{}, inputPath)
}
