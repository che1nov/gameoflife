package adapters

import (
	"gameoflife/internal/core"
	"math/rand"
)

// Seed fills the universe with random live cells.
func Seed(u core.Universe, width, height int) {
	for i := 0; i < (width * height / 4); i++ {
		x, y := rand.Intn(width), rand.Intn(height)
		u.Set(x, y, true, width, height)
	}
}
