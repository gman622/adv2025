# 🎄 Advent of Code 2025 - Expert Go Solutions

> Production-quality Go solutions showcasing idiomatic patterns, clean architecture, and functional programming techniques.

[![Go Version](https://img.shields.io/badge/Go-1.24.5-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## 🌟 Project Philosophy

This repository **isn't about writing the shortest or fastest solutions**. Instead, it demonstrates:

- 🏗️ **Production-quality architecture** - How to structure Go code for maintainability
- 🎯 **Idiomatic Go patterns** - Interfaces, custom types, composition over inheritance
- 🔧 **Clean code principles** - Single responsibility, dependency injection, testability
- 🚀 **Advanced features** - Generics, functional patterns (Map/Filter/Reduce), method chaining
- 📚 **Self-documenting code** - Clear types, meaningful names, strategic comments

**Day 1** serves as the architectural reference implementation, showcasing 10+ Go best practices in a single solution.

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/gman622/adv2025.git
cd adv2025

# Run all solutions
go run cmd/main.go

# Run a specific day
go run cmd/main.go -day 1

# Run a specific part
go run cmd/main.go -day 1 -part 1
```

## 📁 Project Structure

```
adv2025/
├── aoc/                    # Daily solutions as packages
│   └── day1/               # Each day is self-contained
│       ├── types.go        # Domain-specific types
│       ├── counter.go      # Strategy pattern implementations
│       ├── parser.go       # Input parsing with io.Reader
│       ├── solution.go     # Core solver + functional patterns
│       ├── part1.go        # Part 1 solution (one-liner)
│       ├── part2.go        # Part 2 solution (one-liner)
│       ├── example.go      # Usage examples
│       └── README.md       # Day-specific documentation
├── cmd/
│   ├── main.go             # Centralized runner with timing
│   └── day1.md             # Problem descriptions
├── inputs/
│   └── day1_input.txt      # Puzzle inputs
├── CLAUDE.md               # Development guidelines
└── README.md               # This file
```

## 🎯 Day 1: Showcase Implementation

Day 1 demonstrates expert-level Go through a "safe dial" puzzle. The implementation showcases:

### 1. Strong Type System
```go
type Position int       // Not just "int" - semantic meaning
type Direction rune     // Type-safe directions
type Rotation struct {  // Composite types for domain modeling
    Direction Direction
    Distance  int
}
```

### 2. Strategy Pattern
```go
type Counter interface {
    Count(rotation Rotation, position Position) int
}

// Part 1 uses EndPositionCounter
// Part 2 uses ZeroCrossingCounter
```

### 3. Functional Programming
```go
pipeline.
    Filter(func(r Rotation) bool { return r.Direction == Left }).
    Map(func(r Rotation) Rotation { return transform(r) }).
    Reduce(initialPos, counter)
```

### 4. Method Chaining
```go
NewDial(counter).
    Rotate(r1).
    Rotate(r2).
    Rotate(r3).
    Count()
```

### 5. Clean Solutions
```go
// Part 1: Count end positions at zero
func Part1(inputPath string) (int, error) {
    return SolveWith(EndPositionCounter{}, inputPath)
}

// Part 2: Count all zero crossings
func Part2(inputPath string) (int, error) {
    return SolveWith(ZeroCrossingCounter{}, inputPath)
}
```

**See [aoc/day1/README.md](aoc/day1/README.md) for detailed explanation of all patterns used.**

## 🏆 Progress

| Day | Part 1 | Part 2 | Patterns Showcased |
|-----|--------|--------|-------------------|
| 1   | ⭐     | ⭐     | Strategy, Functional, Types, io.Reader, Stringer |
| 2   | ⭐     | ⭐     | Range validation, Custom types, Strategy pattern |
| 3   | ⭐     | ⭐     | Grid algorithms, Parser validation, Direct implementation |
| 4   | ⭐     | ⭐     | Iterative algorithms, Mutable grids, Adjacency checking |

## 🛠️ Architecture Highlights

### Plugin-Style Solutions
Each day is a Go package that exports `Part1` and `Part2` functions:
```go
func Part1(inputPath string) (int, error)
func Part2(inputPath string) (int, error)
```

### Centralized Runner
The main runner automatically discovers and executes solutions:
```go
var solvers = []DayPart{
    {1, 1, day1.Part1},
    {1, 2, day1.Part2},
}
```

### Type-Safe Domain Modeling
Custom types prevent mixing incompatible values:
```go
position := Position(50)  // Can't accidentally pass a Distance
```

### Composable Abstractions
Small, focused interfaces enable flexible composition:
```go
type Counter interface {
    Count(rotation Rotation, position Position) int
}
```

## 🎓 Go Features Demonstrated

This repository showcases:

- ✅ Custom types with methods
- ✅ Interface-based polymorphism
- ✅ Strategy pattern
- ✅ Functional programming (Map/Filter/Reduce)
- ✅ Method chaining / Fluent API
- ✅ io.Reader for testable I/O
- ✅ Error wrapping with `%w`
- ✅ fmt.Stringer interface
- ✅ Constructor pattern (`New*` functions)
- ✅ Pure functions (no side effects)
- ✅ Composition over inheritance
- ✅ Encapsulation (unexported fields)

## 📊 Performance

Solutions prioritize clarity but remain efficient:

```
🎄 Advent of Code 2025 Runner
==================================================

✅ Day 1 Part 1: 1147 (581µs)
✅ Day 1 Part 2: 6789 (212µs)
✅ Day 2 Part 1: 56660955519 (56.6ms)
✅ Day 2 Part 2: 79183223243 (56.7ms)
✅ Day 3 Part 1: 17405 (1.9ms)
✅ Day 3 Part 2: 171990312704598 (105µs)
✅ Day 4 Part 1: 1409 (755µs)
✅ Day 4 Part 2: 8366 (6.5ms)

⏱️  Total time: 123ms
```

## 🔧 Development

```bash
# Format code
go fmt ./...

# Run linter (if installed)
golangci-lint run

# Build standalone binary
go build -o aoc-runner cmd/main.go

# Run binary
./aoc-runner -day 1
```

## 📚 Learning Resources

Each day includes:
- ✅ Problem description (`cmd/day{N}.md`)
- ✅ Implementation with inline comments
- ✅ README explaining patterns used
- ✅ Example usage code

**See [CLAUDE.md](CLAUDE.md) for comprehensive development guidelines.**

## 🤝 Contributing

This is a personal learning repository, but suggestions for improving Go idioms are welcome!

## 📝 License

MIT License - feel free to use this code for learning.

## 🙏 Acknowledgments

- [Advent of Code](https://adventofcode.com/) by Eric Wastl
- The Go community for excellent documentation and idioms
- Built with assistance from [Claude Code](https://claude.com/claude-code)

---

**Note**: This repository favors **clarity and learning** over brevity. If you're looking for minimal solutions, there are many excellent AoC repositories. This one is about writing Go code you'd be proud to deploy to production.
