package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	// --- FIX: ADD REQUIRED LOCAL IMPORTS ---
	// The import path is constructed from:
	// <ModulePrefix>/<DirectoryPath to the package>
	day1 "adv2025/aoc/day1"
	day2 "adv2025/aoc/day2"
	// ----------------------------------------pw
)

//nolint:staticcheck
type DayPart struct {
	day   int
	part  int
	solve func(inputPath string) (int, error)
}


var solvers = []DayPart{
	{1, 1, day1.Part1},
	{1, 2, day1.Part2},
	{2, 1, day2.Part1},
	{2, 2, day2.Part2},
}

func main() {
	dayFlag := flag.Int("day", 0, "Day to run (0 for all)")
	partFlag := flag.Int("part", 0, "Part to run (0 for all parts of the day)")
	flag.Parse()

	// Find inputs directory - try multiple paths
	var inputDir string
	possiblePaths := []string{
		filepath.Join(".", "go", "inputs"),
		filepath.Join(".", "inputs"),
		filepath.Join("..", "go", "inputs"),
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			inputDir = path
			break
		}
	}

	if inputDir == "" {
		log.Fatalf("Inputs directory not found. Tried: %v", possiblePaths)
	}

	// Filter solvers based on flags
	var toRun []DayPart
	if *dayFlag == 0 {
		// Run all
		toRun = solvers
	} else {
		// Filter by day and optional part
		for _, s := range solvers {
			if s.day == *dayFlag {
				if *partFlag == 0 || s.part == *partFlag {
					toRun = append(toRun, s)
				}
			}
		}
		if len(toRun) == 0 {
			log.Fatalf("No solutions found for day %d part %d", *dayFlag, *partFlag)
		}
	}

	// Run solvers
	fmt.Println("🎄 Advent of Code 2025 Runner")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()

	totalStart := time.Now()
	for _, solver := range toRun {
		inputPath := filepath.Join(inputDir, fmt.Sprintf("day%d_input.txt", solver.day))

		if _, err := os.Stat(inputPath); err != nil {
			fmt.Printf("❌ Day %d Part %d: Input file not found at %s\n", solver.day, solver.part, inputPath)
			continue
		}

		start := time.Now()
		result, err := solver.solve(inputPath)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("❌ Day %d Part %d: %v\n", solver.day, solver.part, err)
		} else {
			fmt.Printf("✅ Day %d Part %d: %d (%v)\n", solver.day, solver.part, result, elapsed)
		}
	}

	totalElapsed := time.Since(totalStart)
	fmt.Println()
	fmt.Printf("⏱️  Total time: %v\n", totalElapsed)
}
