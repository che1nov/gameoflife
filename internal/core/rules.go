package core

// Step updates the next universe (b) from the current universe (a).
func Step(a, b Universe, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y, width, height), width, height)
		}
	}
}
