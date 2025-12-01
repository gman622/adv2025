package day1

// Counter defines a strategy for counting during dial rotations
type Counter interface {
	// Count processes a rotation and returns the count contribution
	Count(rotation Rotation, position Position) int
}

// EndPositionCounter counts only when the dial ends at position 0
type EndPositionCounter struct{}

func (EndPositionCounter) Count(rotation Rotation, position Position) int {
	if rotation.Apply(position).IsZero() {
		return 1
	}
	return 0
}

// ZeroCrossingCounter counts every time the dial passes through 0
type ZeroCrossingCounter struct{}

func (ZeroCrossingCounter) Count(rotation Rotation, position Position) int {
	return rotation.CountZeroCrossings(position)
}

// Dial represents the safe's dial with a current position
type Dial struct {
	position Position
	counter  Counter
	count    int
}

// NewDial creates a dial starting at position 50 with the given counter strategy
func NewDial(counter Counter) *Dial {
	return &Dial{
		position: 50,
		counter:  counter,
	}
}

// Rotate applies a rotation, updates the count, and returns the dial for chaining
func (d *Dial) Rotate(r Rotation) *Dial {
	d.count += d.counter.Count(r, d.position)
	d.position = r.Apply(d.position)
	return d
}

// Position returns the current position
func (d *Dial) Position() Position {
	return d.position
}

// Count returns the accumulated count
func (d *Dial) Count() int {
	return d.count
}

// Reset resets the dial to position 50 with zero count
func (d *Dial) Reset() *Dial {
	d.position = 50
	d.count = 0
	return d
}
