package day1

func Part2(inputPath string) (int, error) {
	d := newDial()
	count := 0

	err := parseRotations(inputPath, func(direction byte, distance int) error {
		count += d.countPassesThroughZero(direction, distance)
		d.rotate(direction, distance)
		return nil
	})

	return count, err
}
