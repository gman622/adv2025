# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an Advent of Code 2025 solution repository written in Go. The project uses a centralized runner pattern where each day's solutions are implemented as separate packages and registered with a main runner.

## Architecture

### Directory Structure

```
adv2025/
├── aoc/            # Solution packages organized by day
│   └── day{N}/     # Each day has its own package
│       ├── part1.go
│       └── part2.go
├── cmd/            # Main runner and problem descriptions
│   ├── main.go     # Centralized runner
│   └── day{N}.md   # Problem descriptions (copied from AOC website)
└── inputs/         # Puzzle inputs
    └── day{N}_input.txt
```

### Solution Pattern

Each day's solution follows this pattern:

1. **Package Structure**: Solutions live in `aoc/day{N}/` with separate `part1.go` and `part2.go` files
2. **Function Signature**: Each part exports a function: `func Part{N}(inputPath string) (int, error)`
3. **Registration**: Import and register solvers in `cmd/main.go`:
   ```go
   import day{N} "adv2025/aoc/day{N}"

   var solvers = []DayPart{
       {N, 1, day{N}.Part1},
       {N, 2, day{N}.Part2},
   }
   ```

### Module Configuration

- Module name: `adv2025` (defined in `go.mod`)
- Import paths use format: `adv2025/aoc/day{N}`

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

### Input File Expectations

- Input files must be named `day{N}_input.txt` in the `inputs/` directory
- The runner looks for inputs in: `./go/inputs`, `./inputs`, or `../go/inputs`

## Adding a New Day

1. Create directory: `aoc/day{N}/`
2. Implement `part1.go` with function: `func Part1(inputPath string) (int, error)`
3. Implement `part2.go` with function: `func Part2(inputPath string) (int, error)`
4. Add import to `cmd/main.go`: `day{N} "adv2025/aoc/day{N}"`
5. Register solvers in `cmd/main.go` by adding to the `solvers` slice
6. Place input in `inputs/day{N}_input.txt`
7. Optionally add problem description to `cmd/day{N}.md`

## Implementation Notes

- Solutions read input files directly (not stdin)
- All solutions return `(int, error)` - adapt for different return types as needed
- The runner provides timing information for each solution
- Empty lines in input should be handled gracefully (typically skipped)
