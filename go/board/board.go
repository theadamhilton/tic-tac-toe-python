package board

import (
	"fmt"
)

// GameState holds board dimensions and current game board.
type GameState struct {
	Board  [][]int // Dynamic board
	Width  int
	Height int
}

// NewBoard initializes a board with given dimensions.
func NewBoard(width, height int) *GameState {
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}
	return &GameState{Board: board, Width: width, Height: height}
}

// RenderBoard displays the current board in a visually friendly way.
func (g *GameState) RenderBoard() {
	// fmt.Println("  ", "0 1 2") // Adjust dynamically based on width
	// Generate dynamic column headers
	header := "   "
	for i := 0; i < g.Width; i++ {
		header += fmt.Sprintf("%d ", i)
	}
	fmt.Println(header)

	// fmt.Println("  ", "------")
	// Generate dynamic separator line
	separator := "   "
	for i := 0; i < g.Width; i++ {
		separator += "--"
	}
	fmt.Println(separator)

	// Render the board
	for y := 0; y < g.Height; y++ {
		row := ""
		for x := 0; x < g.Width; x++ {
			if g.Board[y][x] == 0 {
				row += " ." // Empty space
			} else if g.Board[y][x] == 1 {
				row += " X" // Player 1
			} else if g.Board[y][x] == 2 {
				row += " O" // Player 2
			}
		}
		fmt.Printf("%d |%s |\n", y, row)
	}

	// fmt.Println("  ", "------")
	fmt.Println(separator) // Footer separator
}
