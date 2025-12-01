package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// Position represents a position on the dial (0-99)
type Position int

// Normalize ensures position is in valid range [0, 100)
func (p Position) Normalize() Position {
	pos := int(p) % 100
	if pos < 0 {
		pos += 100
	}
	return Position(pos)
}

// IsZero checks if position is at 0
func (p Position) IsZero() bool {
	return p == 0
}

// Direction represents the rotation direction
type Direction rune

const (
	Left  Direction = 'L'
	Right Direction = 'R'
)

// String implements fmt.Stringer
func (d Direction) String() string {
	return string(d)
}

// Rotation represents a single dial rotation instruction
type Rotation struct {
	Direction Direction
	Distance  int
}

// ParseRotation parses a rotation string like "L68" or "R48"
func ParseRotation(s string) (Rotation, error) {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return Rotation{}, fmt.Errorf("invalid rotation: too short")
	}

	dir := Direction(s[0])
	if dir != Left && dir != Right {
		return Rotation{}, fmt.Errorf("invalid direction: %c", s[0])
	}

	distance, err := strconv.Atoi(s[1:])
	if err != nil {
		return Rotation{}, fmt.Errorf("invalid distance in %q: %w", s, err)
	}

	return Rotation{Direction: dir, Distance: distance}, nil
}

// String implements fmt.Stringer
func (r Rotation) String() string {
	return fmt.Sprintf("%s%d", r.Direction, r.Distance)
}

// Apply applies the rotation to a position and returns the new position
func (r Rotation) Apply(p Position) Position {
	switch r.Direction {
	case Left:
		return (p - Position(r.Distance)).Normalize()
	case Right:
		return (p + Position(r.Distance)).Normalize()
	default:
		return p
	}
}

// CountZeroCrossings counts how many times this rotation crosses position 0
func (r Rotation) CountZeroCrossings(from Position) int {
	switch r.Direction {
	case Left:
		if from.IsZero() {
			// Starting at 0, count complete wraps
			return r.Distance / 100
		}
		// Going left from position p, we hit 0 after p steps
		if r.Distance >= int(from) {
			return 1 + (r.Distance-int(from))/100
		}
		return 0
	case Right:
		// Going right, we cross 0 every 100 steps starting from (100 - position)
		return (int(from) + r.Distance) / 100
	default:
		return 0
	}
}
