package core

// Universe represents a two-dimensional grid of cells.
type Universe [][]bool

// NewUniverse creates an empty universe.
func NewUniverse(width, height int) Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Set sets the state of a specific cell.
func (u Universe) Set(x, y int, alive bool, width, height int) {
	x = (x + width) % width // Wrap around if out of bounds
	y = (y + height) % height
	u[y][x] = alive
}

// Alive checks if a cell is alive.
func (u Universe) Alive(x, y int, width, height int) bool {
	x = (x + width) % width // Wrap around if out of bounds
	y = (y + height) % height
	return u[y][x]
}

// Neighbors counts the number of live neighboring cells.
func (u Universe) Neighbors(x, y int, width, height int) int {
	count := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v, width, height) {
				count++
			}
		}
	}
	return count
}

// Next calculates the state of a cell for the next step.
func (u Universe) Next(x, y int, width, height int) bool {
	n := u.Neighbors(x, y, width, height)
	return n == 3 || (n == 2 && u.Alive(x, y, width, height))
}
