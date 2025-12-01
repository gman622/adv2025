package day1

import "fmt"

// Solver represents a generic dial rotation problem solver
type Solver struct {
	dial   *Dial
	parser *RotationParser
}

// NewSolver creates a solver with a specific counter strategy
func NewSolver(counter Counter) *Solver {
	return &Solver{
		dial: NewDial(counter),
	}
}

// Solve processes all rotations and returns the final count
func (s *Solver) Solve(inputPath string) (int, error) {
	err := ProcessFile(inputPath, func(r Rotation) error {
		s.dial.Rotate(r)
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("solving: %w", err)
	}
	return s.dial.Count(), nil
}

// SolveWith is a functional approach - takes a counter and input path
func SolveWith(counter Counter, inputPath string) (int, error) {
	return NewSolver(counter).Solve(inputPath)
}

// Pipeline creates a processing pipeline for rotations
type Pipeline struct {
	rotations []Rotation
}

// NewPipeline creates a pipeline from a file
func NewPipeline(inputPath string) (*Pipeline, error) {
	f, err := FromFile(inputPath)
	if err != nil {
		return nil, err
	}

	rotations, err := f.ParseAll()
	if err != nil {
		return nil, err
	}

	return &Pipeline{rotations: rotations}, nil
}

// Reduce applies a reduction function over all rotations
func (p *Pipeline) Reduce(initial Position, counter Counter) int {
	position := initial
	count := 0

	for _, rotation := range p.rotations {
		count += counter.Count(rotation, position)
		position = rotation.Apply(position)
	}

	return count
}

// Map applies a transformation to each rotation
func (p *Pipeline) Map(fn func(Rotation) Rotation) *Pipeline {
	mapped := make([]Rotation, len(p.rotations))
	for i, r := range p.rotations {
		mapped[i] = fn(r)
	}
	return &Pipeline{rotations: mapped}
}

// Filter keeps only rotations that match the predicate
func (p *Pipeline) Filter(predicate func(Rotation) bool) *Pipeline {
	var filtered []Rotation
	for _, r := range p.rotations {
		if predicate(r) {
			filtered = append(filtered, r)
		}
	}
	return &Pipeline{rotations: filtered}
}

// Count returns the number of rotations
func (p *Pipeline) Count() int {
	return len(p.rotations)
}
