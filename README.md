## Conway's Game of Life in Go
This is an implementation of John Conway's classic "Game of Life" in the Go programming language. The game is a cellular automaton where cells can be "alive" or "dead", and they interact with their neighbors according to specific rules.

### Game Description
The Game of Life consists of a grid of cells, each of which can be either alive or dead. At each time step, the state of all cells is updated based on the state of their neighbors.

### Game Rules
- Birth: If a dead cell has exactly 3 living neighbors, it becomes alive.
- Survival: If a living cell has 2 or 3 living neighbors, it remains alive.
- Death: If a living cell has fewer than 2 or more than 3 living neighbors, it dies.
- The game grid is a two-dimensional array with finite size, and cells that go beyond the grid boundaries "reappear" on the opposite side (periodic boundary).

### Main Components
- Universe: A structure that represents a two-dimensional array of cells, where each cell can be either alive (true) or dead (false).
- NewUniverse: A function that creates a new empty universe with given dimensions.
- Seed: A function that fills the universe with random living cells.
- Set: A function to set the state of a cell in a specific position (alive/dead).
- Alive: A function to check if a cell is alive.
- Neighbors: A function to count the number of living neighbors around a cell.
- Next: A function to compute the next state of a cell based on its neighbors.
- Step: A function that updates the state of the universe for the next time step.
- Show: A function to display the current state of the universe on the screen.

### How to Run
- Make sure you have Go installed. If Go is not installed, download and install it from the official website: https://golang.org/dl/

- Download or clone this repository to your machine.

- Navigate to the project directory and run the command:

`go run main.go`

The game will start, and you will see an animation of the cells' states changing. Each update will occur at an interval of 1/30 of a second.

### Project Structure
- main.go: The main program file that contains all the game logic and cell display.
- README.md: This file with instructions and project description.

### Additional Settings
- You can change the size of the universe by adjusting the values of the width and height constants at the beginning of the file.
- You can also change the game speed by modifying the delay in `time.Sleep(time.Second / 30)` in the `main` function.

### Sample Output
When you run the program, the screen will display living cells (denoted by the * symbol) and dead cells (denoted by spaces). The game will continue running until manually stopped.

#### License
This project is distributed under the MIT license. See LICENSE for more information.