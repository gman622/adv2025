package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Parser reads and parses input for Day 4
type Parser struct {
	scanner *bufio.Scanner
}

// NewParser creates a parser from an io.Reader
func NewParser(r io.Reader) *Parser {
	return &Parser{
		scanner: bufio.NewScanner(r),
	}
}

// ParseAll reads all lines from the input
func (p *Parser) ParseAll() ([]string, error) {
	var lines []string
	lineNum := 0

	for p.scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(p.scanner.Text())
		if line == "" {
			continue
		}

		// Validate that the line contains only '@' and '.' characters
		for _, ch := range line {
			if ch != '@' && ch != '.' {
				return nil, fmt.Errorf("line %d: invalid character %q, expected '@' or '.'", lineNum, ch)
			}
		}

		lines = append(lines, line)
	}

	if err := p.scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}

	return lines, nil
}

// FromFile creates a parser from a file path and parses all lines immediately
func FromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	parser := NewParser(file)
	return parser.ParseAll()
}
