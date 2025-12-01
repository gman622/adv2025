# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an Advent of Code 2025 solution repository written in Go, showcasing expert-level Go patterns and best practices. The project uses a centralized runner with a plugin-style architecture where each day's solutions are implemented as separate packages following idiomatic Go design patterns.

## Architecture Philosophy

This codebase demonstrates **production-quality Go** rather than minimal solutions:
- Strong typing with domain-specific types
- Interface-based design for polymorphism
- Functional programming patterns (Map/Filter/Reduce)
- Comprehensive error handling with context
- Composable abstractions and clean separation of concerns

**Day 1** serves as the architectural reference - it's intentionally over-engineered to showcase Go's capabilities.

## Directory Structure

```
adv2025/
├── aoc/                    # Solution packages organized by day
│   └── day{N}/             # Each day is a self-contained package
│       ├── types.go        # Domain types and custom type definitions
│       ├── counter.go      # Interfaces and strategy implementations
│       ├── parser.go       # Input parsing with io.Reader
│       ├── solution.go     # Core solver logic and functional patterns
│       ├── part1.go        # Part 1 solution (typically a one-liner)
│       ├── part2.go        # Part 2 solution (typically a one-liner)
│       ├── example.go      # Usage examples (optional)
│       └── README.md       # Day-specific documentation (optional)
├── cmd/                    # Main runner and problem descriptions
│   ├── main.go             # Centralized runner with timing
│   └── day{N}.md           # Problem descriptions from AOC
├── inputs/                 # Puzzle inputs
│   └── day{N}_input.txt    # Input data for each day
├── CLAUDE.md               # This file
├── README.md               # Project README
└── go.mod                  # Module definition
```

## Solution Pattern

### Core Requirements

Each day's solution must export:
```go
func Part1(inputPath string) (int, error)
func Part2(inputPath string) (int, error)
```

### Recommended Architecture (following Day 1 pattern)

1. **types.go** - Domain modeling
   - Custom types for problem domain (e.g., `Position`, `Direction`)
   - Methods on types for domain operations
   - Implement `fmt.Stringer` for debugging

2. **counter.go** / strategy implementation
   - Define interfaces for behavioral variation
   - Implement strategy pattern when parts differ in behavior
   - Use composition over inheritance

3. **parser.go** - Input handling
   - Accept `io.Reader` for testability
   - Provide `ProcessFile()` convenience function
   - Parse line-by-line with `bufio.Scanner`
   - Wrap errors with context (`fmt.Errorf` with `%w`)

4. **solution.go** - Core logic
   - Solver struct composing strategies
   - Functional patterns: Map, Filter, Reduce
   - Reusable pipeline abstractions

5. **part1.go / part2.go** - Clean one-liners
   ```go
   func Part1(inputPath string) (int, error) {
       return SolveWith(Strategy1{}, inputPath)
   }
   ```

### Registration in cmd/main.go

```go
import day{N} "adv2025/aoc/day{N}"

var solvers = []DayPart{
    {N, 1, day{N}.Part1},
    {N, 2, day{N}.Part2},
}
```

## Module Configuration

- Module name: `adv2025` (defined in `go.mod`)
- Import paths: `adv2025/aoc/day{N}`
- Go version: 1.24.5

## Commands

### Running Solutions

```bash
# Run all implemented solutions
go run cmd/main.go

# Run specific day
go run cmd/main.go -day 1

# Run specific part of a day
go run cmd/main.go -day 1 -part 1
```

### Development

```bash
# Format code
go fmt ./...

# Run linter (if installed)
golangci-lint run

# Build binary
go build -o aoc-runner cmd/main.go
```

## Adding a New Day

### Quick Start (Simple Approach)

1. Create directory: `mkdir -p aoc/day{N}`
2. Implement `part1.go` and `part2.go` with required signatures
3. Add import and registration in `cmd/main.go`
4. Place input in `inputs/day{N}_input.txt`

### Expert Approach (Following Day 1 Pattern)

1. **Design domain types** - Create `types.go` with custom types
2. **Identify behavioral variations** - Create interfaces in `counter.go` or similar
3. **Build parser** - Create `parser.go` using `io.Reader`
4. **Implement solver** - Create `solution.go` with functional patterns
5. **Wire it up** - Make `part1.go` and `part2.go` one-liners
6. **Document** - Add `README.md` explaining patterns used
7. **Register** - Update `cmd/main.go`

### File Template (Simple Version)

```go
package day{N}

import (
    "bufio"
    "os"
)

func Part1(inputPath string) (int, error) {
    file, err := os.Open(inputPath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Implementation here

    return result, scanner.Err()
}
```

## Go Best Practices in This Codebase

### Type Safety
- Use custom types instead of primitives (`type Position int` vs `int`)
- This prevents mixing incompatible values

### Interfaces
- Define small, focused interfaces
- Accept interfaces, return concrete types
- Use for strategy/behavior variation

### Error Handling
- Always check errors
- Wrap errors with context: `fmt.Errorf("context: %w", err)`
- Return errors, don't panic

### io.Reader Pattern
- Accept `io.Reader` for parsers (testable without files)
- Provide `FromFile()` convenience functions
- Use `bufio.Scanner` for line-by-line processing

### Method Receivers
- Use pointer receivers for methods that modify state
- Use value receivers for methods that don't mutate
- Be consistent within a type

### Functional Patterns
- Higher-order functions for flexible processing
- Map/Filter/Reduce when working with collections
- Callbacks for custom iteration logic

## Testing Approach

While tests aren't required for AoC, the architecture supports easy testing:

```go
// Parser is testable with strings.NewReader
parser := NewRotationParser(strings.NewReader("L10\nR20"))

// Strategies are testable independently
counter := EndPositionCounter{}
count := counter.Count(rotation, position)
```

## Performance Notes

- Solutions prioritize clarity over micro-optimization
- The runner reports timing for each solution
- Typical solutions run in microseconds
- Consider concurrency for days with independent work items

## Style Guidelines

- Follow standard Go conventions (`gofmt`, `golint`)
- Export only what needs to be public
- Use descriptive names over comments
- Implement `fmt.Stringer` for custom types
- Group related functionality in the same file

## Input File Expectations

- Input files must be named `day{N}_input.txt` in `inputs/`
- The runner searches: `./go/inputs`, `./inputs`, `../go/inputs`
- Parsers should handle empty lines gracefully
- Always close files with `defer file.Close()`
