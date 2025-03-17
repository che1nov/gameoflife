package main

import (
	"gameoflife/internal/adapters"
	"gameoflife/internal/core"
	"math/rand"
	"time"
)

const (
	width  = 100 // Width of the universe
	height = 40  // Height of the universe
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Create two universes: one for the current state and one for the next state
	a, b := core.NewUniverse(width, height), core.NewUniverse(width, height)

	// Seed the initial universe with random live cells
	adapters.Seed(a, width, height)

	// Run the simulation for 5000 steps
	for i := 0; i < 5000; i++ {
		core.Step(a, b, width, height)    // Update the universe
		adapters.Show(a, width, height)   // Display the universe
		time.Sleep(33 * time.Millisecond) // Wait for ~30 FPS
		a, b = b, a                       // Swap universes
	}
}
