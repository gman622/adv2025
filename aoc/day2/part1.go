package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Range represents a product ID range with start and end values
type Range struct {
	Start, End int
}

// Part1 solves Day 2 Part 1: sum all invalid product IDs in the given ranges
func Part1(inputPath string) (int, error) {
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
			if isInvalidID(id) {
				sum += id
			}
		}
	}

	return sum, scanner.Err()
}

// parseRanges parses comma-separated ranges like "11-22,95-115"
func parseRanges(line string) ([]Range, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, fmt.Errorf("empty line")
	}

	// Remove trailing comma if present
	line = strings.TrimSuffix(line, ",")

	parts := strings.Split(line, ",")
	ranges := make([]Range, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		nums := strings.Split(part, "-")
		if len(nums) != 2 {
			return nil, fmt.Errorf("invalid range format: %s", part)
		}

		start, err := strconv.Atoi(strings.TrimSpace(nums[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start number: %w", err)
		}

		end, err := strconv.Atoi(strings.TrimSpace(nums[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end number: %w", err)
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges, nil
}

// isInvalidID checks if an ID is made of a sequence repeated exactly twice
// Examples: 55 (5 twice), 6464 (64 twice), 123123 (123 twice)
func isInvalidID(id int) bool {
	s := strconv.Itoa(id)

	// Must have even length to be repeated twice
	if len(s)%2 != 0 {
		return false
	}

	// Check for leading zeros (numbers like 0101 are not valid IDs)
	if s[0] == '0' {
		return false
	}

	// Split in half and check if both halves are equal
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf == secondHalf
}
