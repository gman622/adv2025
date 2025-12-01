package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// mod returns a positive modulo result (Go's % can return negative values)
func mod(n, m int) int {
	result := n % m
	if result < 0 {
		result += m
	}
	return result
}

// dial represents the safe's dial position
type dial struct {
	position int
}

// newDial creates a dial starting at position 50
func newDial() *dial {
	return &dial{position: 50}
}

// rotate performs a rotation and returns the new position
func (d *dial) rotate(direction byte, distance int) int {
	switch direction {
	case 'L':
		d.position = mod(d.position-distance, 100)
	case 'R':
		d.position = mod(d.position+distance, 100)
	}
	return d.position
}

// countPassesThroughZero counts how many times the dial passes through 0 during a rotation
func (d *dial) countPassesThroughZero(direction byte, distance int) int {
	count := 0
	switch direction {
	case 'L':
		// Going left from position p, we hit 0 after p clicks (unless p is 0)
		// Then every 100 clicks after that
		if d.position == 0 {
			count = distance / 100
		} else if distance >= d.position {
			count = 1 + (distance-d.position)/100
		}
	case 'R':
		// Going right from position p, we hit 0 after (100-p) clicks
		// Which means every time (p + distance) crosses a multiple of 100
		count = (d.position + distance) / 100
	}
	return count
}

// parseRotations reads rotations from a file and calls a handler for each one
func parseRotations(inputPath string, handler func(direction byte, distance int) error) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return fmt.Errorf("invalid rotation: %s", line)
		}

		if direction != 'L' && direction != 'R' {
			return fmt.Errorf("invalid direction: %c", direction)
		}

		if err := handler(direction, distance); err != nil {
			return err
		}
	}

	return scanner.Err()
}
