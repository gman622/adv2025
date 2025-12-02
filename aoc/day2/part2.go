package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part2 solves Day 2 Part 2: sum all invalid product IDs with relaxed rules
// An ID is invalid if it's made of a pattern repeated at least twice
func Part2(inputPath string) (int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return 0, fmt.Errorf("opening input file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, fmt.Errorf("empty input file")
	}

	line := scanner.Text()
	ranges, err := parseRanges(line)
	if err != nil {
		return 0, fmt.Errorf("parsing ranges: %w", err)
	}

	sum := 0
	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalidIDPart2(id) {
				sum += id
			}
		}
	}

	return sum, scanner.Err()
}

// isInvalidIDPart2 checks if an ID is made of a pattern repeated at least twice
// Examples: 111 (1 three times), 123123123 (123 three times), 1212 (12 twice)
func isInvalidIDPart2(id int) bool {
	s := strconv.Itoa(id)

	// Check for leading zeros (numbers like 0101 are not valid IDs)
	if s[0] == '0' {
		return false
	}

	n := len(s)

	// Try all possible pattern lengths from 1 to n/2
	// The pattern must repeat at least twice
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		// The string length must be divisible by pattern length
		if n%patternLen != 0 {
			continue
		}

		// Check if the string is made by repeating the first patternLen characters
		pattern := s[:patternLen]
		isRepeating := true

		for i := patternLen; i < n; i += patternLen {
			if s[i:i+patternLen] != pattern {
				isRepeating = false
				break
			}
		}

		if isRepeating {
			return true
		}
	}

	return false
}
