package day4

import "fmt"

// Part1 solves Day 4 Part 1: count rolls of paper accessible by forklifts
// A roll is accessible if it has fewer than 4 adjacent rolls (8 directions)
func Part1(inputPath string) (int, error) {
	grid, err := FromFile(inputPath)
	if err != nil {
		return 0, fmt.Errorf("loading input: %w", err)
	}

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' && isAccessible(grid, row, col) {
				count++
			}
		}
	}

	return count, nil
}

// isAccessible returns true if a roll at (row, col) has fewer than 4 adjacent rolls
func isAccessible(grid []string, row, col int) bool {
	adjacentCount := 0

	// Check all 8 adjacent positions
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // top row
		{0, -1}, {0, 1}, // left and right
		{1, -1}, {1, 0}, {1, 1}, // bottom row
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check bounds
		if newRow >= 0 && newRow < len(grid) &&
			newCol >= 0 && newCol < len(grid[newRow]) &&
			grid[newRow][newCol] == '@' {
			adjacentCount++
		}
	}

	return adjacentCount < 4
}
