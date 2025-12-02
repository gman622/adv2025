package day2

import "fmt"

// Part2 solves Day 2 Part 2: sum all invalid product IDs with relaxed rules
// An ID is invalid if it's made of a pattern repeated at least twice
func Part2(inputPath string) (int, error) {
	parser, err := FromFile(inputPath)
	if err != nil {
		return 0, fmt.Errorf("loading input: %w", err)
	}

	ranges, err := parser.ParseAll()
	if err != nil {
		return 0, fmt.Errorf("parsing ranges: %w", err)
	}

	validator := AtLeastTwiceValidator{}
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
