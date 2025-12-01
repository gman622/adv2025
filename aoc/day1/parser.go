package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// RotationParser reads and parses rotations from an io.Reader
type RotationParser struct {
	scanner *bufio.Scanner
}

// NewRotationParser creates a parser from an io.Reader
func NewRotationParser(r io.Reader) *RotationParser {
	return &RotationParser{
		scanner: bufio.NewScanner(r),
	}
}

// Parse reads all rotations and applies a function to each one
func (p *RotationParser) Parse(fn func(Rotation) error) error {
	lineNum := 0
	for p.scanner.Scan() {
		lineNum++
		line := p.scanner.Text()
		if line == "" {
			continue
		}

		rotation, err := ParseRotation(line)
		if err != nil {
			return fmt.Errorf("line %d: %w", lineNum, err)
		}

		if err := fn(rotation); err != nil {
			return fmt.Errorf("line %d: %w", lineNum, err)
		}
	}

	if err := p.scanner.Err(); err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	return nil
}

// ParseAll reads all rotations into a slice
func (p *RotationParser) ParseAll() ([]Rotation, error) {
	var rotations []Rotation
	err := p.Parse(func(r Rotation) error {
		rotations = append(rotations, r)
		return nil
	})
	return rotations, err
}

// FromFile creates a parser from a file path
func FromFile(path string) (*RotationParser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	// Note: caller should close the file, but for this use case we'll accept the defer
	return NewRotationParser(f), nil
}

// ProcessFile is a convenience function that opens a file, parses it, and processes each rotation
func ProcessFile(path string, fn func(Rotation) error) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	parser := NewRotationParser(f)
	return parser.Parse(fn)
}
