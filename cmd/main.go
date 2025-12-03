package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	day1 "adv2025/aoc/day1"
	day2 "adv2025/aoc/day2"
	day3 "adv2025/aoc/day3"
)

type solver struct {
	day   int
	part  int
	solve func(string) (int, error)
}

var solvers = []solver{
	{1, 1, day1.Part1},
	{1, 2, day1.Part2},
	{2, 1, day2.Part1},
	{2, 2, day2.Part2},
	{3, 1, day3.Part1},
	{3, 2, day3.Part2},
}

func main() {
	day := flag.Int("day", 0, "Day to run (0 for all)")
	part := flag.Int("part", 0, "Part to run (0 for all parts of the day)")
	flag.Parse()

	toRun := filterSolvers(*day, *part)
	if len(toRun) == 0 {
		log.Fatalf("No solutions found for day %d part %d", *day, *part)
	}

	printHeader()
	totalStart := time.Now()

	for _, s := range toRun {
		runSolver(s)
	}

	fmt.Printf("\n⏱️  Total time: %v\n", time.Since(totalStart))
}

func filterSolvers(day, part int) []solver {
	if day == 0 {
		return solvers
	}

	var filtered []solver
	for _, s := range solvers {
		if s.day == day && (part == 0 || s.part == part) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func runSolver(s solver) {
	inputPath := filepath.Join("inputs", fmt.Sprintf("day%d_input.txt", s.day))

	if _, err := os.Stat(inputPath); err != nil {
		fmt.Printf("❌ Day %d Part %d: Input file not found\n", s.day, s.part)
		return
	}

	start := time.Now()
	result, err := s.solve(inputPath)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("❌ Day %d Part %d: %v\n", s.day, s.part, err)
	} else {
		fmt.Printf("✅ Day %d Part %d: %d (%v)\n", s.day, s.part, result, elapsed)
	}
}

func printHeader() {
	fmt.Println("🎄 Advent of Code 2025 Runner")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}
