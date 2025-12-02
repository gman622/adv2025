# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an Advent of Code 2025 solution repository written in Go, showcasing expert-level Go patterns and best practices. The project uses a centralized runner with a plugin-style architecture where each day's solutions are implemented as separate packages following idiomatic Go design patterns.

## Architecture Philosophy

This codebase follows **idiomatic Go best practices** for all solutions, embodying the language's core philosophy:
- **Clear is better than clever** - Well-structured abstractions over ad-hoc solutions
- **Interfaces define behavior** - Small, focused interfaces enable polymorphism and testability
- **Composition over inheritance** - Strategy pattern and struct embedding for code reuse
- **Strong typing prevents errors** - Domain-specific types (not primitive obsession)
- **Errors are values** - Explicit error handling with context wrapping
- **Accept interfaces, return concrete types** - Flexible inputs, concrete outputs

**IMPORTANT: Day 1's architecture is the DEFAULT pattern for all days.** Every solution should follow this structure unless explicitly instructed otherwise. This approach aligns with Go's philosophy of writing maintainable, testable, and extensible code.

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

### Required Architecture (Day 1 pattern - use for ALL days)

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

**ALWAYS follow the Day 1 pattern unless explicitly told to use a simpler approach.**

### Step-by-Step Process

1. **Design domain types** - Create `types.go` with custom types
   - Model the problem domain (e.g., `Position`, `Instruction`, `Range`)
   - Add methods for domain operations
   - Implement `fmt.Stringer` for debugging
   - Use strong typing to prevent errors

2. **Identify behavioral variations** - Create interfaces (e.g., `counter.go`, `validator.go`)
   - Define small, focused interfaces for strategy pattern
   - Implement concrete strategies for Part 1 and Part 2 variations
   - This embodies "accept interfaces, return concrete types"

3. **Build parser** - Create `parser.go` using `io.Reader`
   - Accept `io.Reader` for maximum testability
   - Provide `ProcessFile()` and `FromFile()` convenience functions
   - Parse line-by-line with `bufio.Scanner`
   - Wrap all errors with context (`fmt.Errorf` with `%w`)

4. **Implement solver** - Create `solution.go` with functional patterns
   - Define `Solver` struct that composes strategies
   - Add functional patterns: `Map`, `Filter`, `Reduce` if applicable
   - Create `SolveWith()` helper function
   - Build reusable pipeline abstractions

5. **Wire it up** - Make `part1.go` and `part2.go` clean one-liners
   ```go
   func Part1(inputPath string) (int, error) {
       return SolveWith(Part1Strategy{}, inputPath)
   }
   ```

6. **Register** - Update `cmd/main.go`
   - Import the new package
   - Add entries to `solvers` slice

7. **Document** (optional) - Add `README.md` or problem description
   - Save problem description as `cmd/day{N}.md`
   - Optionally document patterns used in package README

## Go Best Practices (Follow for ALL Solutions)

### Type Safety - Fight Primitive Obsession
- **Always** use custom types instead of primitives (`type Position int` vs `int`)
- This prevents mixing incompatible values and adds domain meaning
- Add methods to custom types for domain operations
- Example: `type ProductID int` with `func (id ProductID) IsValid() bool`

### Interfaces - Define Behavior, Not Data
- Define small, focused interfaces (single-method is ideal)
- **Accept interfaces, return concrete types** - This is Go's golden rule
- Use interfaces for strategy/behavior variation between parts
- Example: `type Validator interface { IsValid(int) bool }`

### Error Handling - Errors Are Values
- **Always** check errors, never ignore them
- Wrap errors with context: `fmt.Errorf("parsing input: %w", err)`
- Return errors, don't panic (panics are for programmer errors only)
- Provide context at each layer of the call stack

### io.Reader Pattern - Depend on Abstractions
- **Always** accept `io.Reader` for parsers (not `*os.File`)
- This makes code testable with `strings.NewReader()`
- Provide `FromFile()` convenience functions that wrap `os.Open()`
- Use `bufio.Scanner` for line-by-line processing

### Composition - Build Complex Behavior from Simple Parts
- Use struct embedding and composition over inheritance
- Example: `Solver` struct composes `Dial` and strategies
- Strategy pattern for varying behavior (Part 1 vs Part 2)
- This embodies "composition over inheritance"

### Method Receivers - Pointer vs Value
- Use pointer receivers `(s *Solver)` for methods that modify state
- Use value receivers `(p Position)` for methods that don't mutate
- Be consistent within a type - don't mix pointer and value receivers

### Functional Patterns - First-Class Functions
- Higher-order functions for flexible processing
- Map/Filter/Reduce when working with collections
- Callbacks for custom iteration logic (`ProcessFile` in Day 1)
- Example: `func ProcessFile(path string, fn func(Rotation) error)`

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
