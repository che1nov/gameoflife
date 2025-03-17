package adapters

import (
	"fmt"
	"gameoflife/internal/core"
)

// Show displays the universe in the terminal.
func Show(u core.Universe, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u.Alive(x, y, width, height) {
				fmt.Print("*") // Live cell
			} else {
				fmt.Print(" ") // Dead cell
			}
		}
		fmt.Println()
	}
	fmt.Println("\x0c") // Clear the screen
}
