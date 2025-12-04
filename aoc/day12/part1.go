package day12

import "fmt"

// Part1 solves Day 12 Part 1
func Part1(inputPath string) (int, error) {
	lines, err := FromFile(inputPath)
	if err != nil {
		return 0, fmt.Errorf("loading input: %w", err)
	}

	// TODO: Implement solution when problem is available
	_ = lines
	return 0, nil
}
