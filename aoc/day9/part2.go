package day9

import "fmt"

// Part2 solves Day 9 Part 2
func Part2(inputPath string) (int, error) {
	lines, err := FromFile(inputPath)
	if err != nil {
		return 0, fmt.Errorf("loading input: %w", err)
	}

	// TODO: Implement solution when problem is available
	_ = lines
	return 0, nil
}
