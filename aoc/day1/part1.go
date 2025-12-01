package day1

func Part1(inputPath string) (int, error) {
	d := newDial()
	count := 0

	err := parseRotations(inputPath, func(direction byte, distance int) error {
		if d.rotate(direction, distance) == 0 {
			count++
		}
		return nil
	})

	return count, err
}
