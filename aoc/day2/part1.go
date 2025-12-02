package day2

import "fmt"

// Part1 solves Day 2 Part 1: sum all invalid product IDs in the given ranges
// An ID is invalid if it's made of a sequence repeated exactly twice
func Part1(inputPath string) (int, error) {
	parser, err := FromFile(inputPath)
	if err != nil {
		return 0, fmt.Errorf("loading input: %w", err)
	}

	ranges, err := parser.ParseAll()
	if err != nil {
		return 0, fmt.Errorf("parsing ranges: %w", err)
	}

	validator := ExactlyTwiceValidator{}
	sum := 0

	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if validator.IsInvalid(id) {
				sum += id
			}
		}
	}

	return sum, nil
}
