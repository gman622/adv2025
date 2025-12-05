# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an Advent of Code 2025 solution repository written in Go, showcasing expert-level Go patterns and best practices. The project uses a centralized runner with a plugin-style architecture where each day's solutions are implemented as separate packages following idiomatic Go design patterns.

## Architecture Philosophy

This codebase follows **idiomatic Go best practices**, embodying the language's core philosophy:

### Go Principles Applied
- **Clear is better than clever** - Well-structured code over complex abstractions
- **Interfaces define behavior** - Small, focused interfaces enable polymorphism
- **Composition over inheritance** - Strategy pattern and struct embedding for code reuse
- **Strong typing prevents errors** - Custom types when they add safety and meaning
- **Errors are values** - Explicit error handling with context wrapping
- **Accept interfaces, return concrete types** - Flexible inputs, concrete outputs

### Pattern Application Strategy
**Use Day 1 as architectural reference, but apply patterns thoughtfully:**
- Always use `io.Reader` for parsers (testability)
- Always separate parsing from solving logic
- Use custom types when they prevent mixing incompatible values
- Use interfaces when parts differ in behavior (not just parameters)
- Use functional patterns for complex collection transformations
- **Don't create files or abstractions just for symmetry**

### Solution Quality Over Speed
**Prioritize elegant, idiomatic Go solutions over brute force:**
- **Elegance first**: Favor clean algorithms and data structures over "just make it work"
- **Idiomatic Go**: Use channels, goroutines, and standard library patterns where appropriate
- **Algorithmic thinking**: Consider time/space complexity, look for O(n) when brute force is O(n²)
- **Data structures matter**: Hash maps for lookups, heaps for priority, tries for prefixes
- **Avoid nested loops when alternatives exist**: Maps, sets, or mathematical solutions often eliminate them
- **Think before coding**: A few minutes of algorithm design saves hours of optimization later

**However, balance is key:**
- If the elegant solution is significantly more complex, brute force is acceptable for small inputs
- Document why brute force was chosen and what alternatives exist
- This is a learning repository - demonstrate good practices even when unnecessary

The goal is production-quality Go code that showcases professional problem-solving, not just correct answers.

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
│       └── example.go      # Usage examples (optional)
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

### Architecture Pattern - Apply Thoughtfully

**Core principle: Use patterns when they add value, not for ceremony.**

#### Always Include:

1. **parser.go** - Input handling with `io.Reader` pattern
   - Accept `io.Reader` for testability (not `*os.File`)
   - Provide `ProcessFile()` and `FromFile()` convenience functions
   - Parse with `bufio.Scanner`
   - Wrap errors with context (`fmt.Errorf` with `%w`)
   - **Why**: Makes code testable with `strings.NewReader()` and provides clean abstraction

2. **part1.go / part2.go** - Clean, focused solution functions
   - Export required signatures: `func Part1(inputPath string) (int, error)`
   - Delegate to solver/strategy when appropriate
   - Keep implementation logic out of these files
   - **Why**: Maintains consistent API and separates concerns

#### Include When Valuable:

3. **types.go** - Domain modeling (when types prevent errors or add meaning)
   - Create custom types when they prevent mixing incompatible values
   - Example: `type Position int` vs `type Direction int` (both ints, different meanings)
   - Add methods for domain operations
   - Implement `fmt.Stringer` for debugging
   - **Skip if**: Primitives like `int`, `string` are already clear (no confusion possible)
   - **Good**: `Range struct { Start, End int }` (clear without wrapping)
   - **Questionable**: `type ProductID int` with no methods (just ceremony)

4. **validator.go / counter.go / etc.** - Strategy interfaces (when parts differ in behavior)
   - Define interfaces for behavioral variation between Part 1 and Part 2
   - Implement strategy pattern for varying logic
   - Use composition over inheritance
   - **Skip if**: Parts differ only in parameters, not behavior
   - **Good**: Different validation algorithms (Part 1: exact twice, Part 2: at least twice)
   - **Questionable**: Same algorithm with different threshold values

5. **solution.go** - Solver and functional patterns (when reusable abstractions exist)
   - Solver struct composing strategies
   - Functional patterns: Map, Filter, Reduce (when working with collections)
   - Reusable pipeline abstractions
   - **Skip if**: Problem is a simple one-pass algorithm
   - **Good**: Multiple transformations, filtering, aggregation
   - **Questionable**: Single loop that's clearer inline

### Registration in cmd/main.go

Each day package exports a `Parts` slice:
```go
// In aoc/day{N}/day{N}.go
package day{N}

var Parts = []func(string) (int, error){Part1, Part2}
```

Register in main using variadic arguments:
```go
import day{N} "adv2025/aoc/day{N}"

func init() {
    register(1, day1.Parts...)
    register(2, day2.Parts...)
    register(N, day{N}.Parts...)
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

**IMPORTANT: Days 5-12 already have scaffolding created!**
- Check for existing files first: `parser.go`, `part1.go`, `part2.go`, `day{N}.go` are already present
- Read existing files to see what's there before trying to create new ones
- Update/append to existing files rather than creating from scratch
- This saves time and avoids failed Write operations

**Follow idiomatic Go patterns, using Day 1 as reference. Apply each pattern thoughtfully based on problem needs.**

### Decision Framework

Ask yourself these questions to decide which patterns to apply:

1. **Do I have incompatible types that could be mixed?** → Custom types in `types.go`
   - Yes: Create domain types (e.g., `Position` vs `Direction`)
   - No: Use primitives or simple structs (e.g., `Range struct { Start, End int }`)

2. **Do Part 1 and Part 2 differ in behavior/algorithm?** → Strategy pattern with interfaces
   - Yes: Create interface + implementations in `validator.go` / `counter.go`
   - No: Share logic, parameterize differences

3. **Is there complex processing with multiple transformations?** → Solver with functional patterns
   - Yes: Create `solution.go` with Map/Filter/Reduce pipelines
   - No: Implement directly in part files

### Implementation Steps

**Always:**

1. **Create `parser.go`** - Input handling with `io.Reader` pattern
   - Accept `io.Reader` for testability
   - Provide `FromFile()` convenience function
   - Parse with `bufio.Scanner`
   - Wrap errors with context
   - Add validation for expected input format

2. **Create `part1.go` and `part2.go`** - Solution entry points
   - Export required signatures: `func Part1(inputPath string) (int, error)`
   - Keep clean and focused
   - Delegate to parsers, solvers, or strategies

3. **Create `day{N}.go`** - Package exports
   - Export `Parts` slice: `var Parts = []func(string) (int, error){Part1, Part2}`
   - Can add Part3+ if needed (rare but supported)

4. **Register in `cmd/main.go`**
   - Import the new day package
   - Add to `init()`: `register(N, dayN.Parts...)`

**When valuable:**

5. **Create `types.go`** (if domain types prevent errors or add meaning)
   - Define custom types for distinct concepts
   - Add domain methods
   - Implement `fmt.Stringer`

6. **Create strategy file** (if parts have different behaviors)
   - Define small interface for variation point
   - Implement concrete strategies
   - Use composition in solver

7. **Create `solution.go`** (if complex transformations exist)
   - Build solver with composed strategies
   - Add functional patterns for collections
   - Provide `SolveWith()` helper

**Example - Simple Day:**
```
day4/
├── day4.go      # Parts slice export
├── parser.go    # io.Reader-based input parsing with validation
├── part1.go     # Direct implementation
└── part2.go     # Direct implementation
```

**Example - Complex Day (like Day 1):**
```
day1/
├── day1.go      # Parts slice export
├── types.go     # Position, Direction, Rotation types
├── counter.go   # Counter interface + strategies
├── parser.go    # io.Reader-based parsing
├── solution.go  # Solver with functional patterns
├── part1.go     # One-liner: SolveWith(Strategy1{}, path)
└── part2.go     # One-liner: SolveWith(Strategy2{}, path)
```

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
