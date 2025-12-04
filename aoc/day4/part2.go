package day4

import "fmt"

// Part2 solves Day 4 Part 2: iteratively remove accessible rolls
// Keep removing accessible rolls until no more can be removed
func Part2(inputPath string) (int, error) {
	lines, err := FromFile(inputPath)
	if err != nil {
		return 0, fmt.Errorf("loading input: %w", err)
	}

	// Convert to mutable grid
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	totalRemoved := 0

	// Keep removing accessible rolls until none remain
	for {
		accessible := findAccessibleRolls(grid)
		if len(accessible) == 0 {
			break
		}

		// Remove all accessible rolls
		for _, pos := range accessible {
			grid[pos.row][pos.col] = '.'
		}

		totalRemoved += len(accessible)
	}

	return totalRemoved, nil
}

type position struct {
	row, col int
}

// findAccessibleRolls returns positions of all accessible rolls in the grid
func findAccessibleRolls(grid [][]byte) []position {
	var accessible []position

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' && isAccessibleMutable(grid, row, col) {
				accessible = append(accessible, position{row, col})
			}
		}
	}

	return accessible
}

// isAccessibleMutable checks if a roll is accessible in a mutable grid
func isAccessibleMutable(grid [][]byte, row, col int) bool {
	adjacentCount := 0

	// Check all 8 adjacent positions
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(grid) &&
			newCol >= 0 && newCol < len(grid[newRow]) &&
			grid[newRow][newCol] == '@' {
			adjacentCount++
		}
	}

	return adjacentCount < 4
}
