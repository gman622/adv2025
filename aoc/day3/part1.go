package day3

import "fmt"

// Part1 solves Day 3 Part 1: find the maximum joltage from each battery bank
// and return the total output joltage
func Part1(inputPath string) (int, error) {
	banks, err := FromFile(inputPath)
	if err != nil {
		return 0, fmt.Errorf("loading input: %w", err)
	}

	totalJoltage := 0
	for _, bank := range banks {
		maxJoltage := findMaxJoltage(bank)
		totalJoltage += maxJoltage
	}

	return totalJoltage, nil
}

// findMaxJoltage finds the maximum two-digit joltage from a battery bank
// by selecting any two batteries (maintaining their order)
func findMaxJoltage(bank string) int {
	if len(bank) < 2 {
		return 0
	}

	maxJoltage := 0

	// Try all pairs of batteries at positions i < j
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			// Form the two-digit number
			digit1 := int(bank[i] - '0')
			digit2 := int(bank[j] - '0')
			joltage := digit1*10 + digit2

			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
	}

	return maxJoltage
}
