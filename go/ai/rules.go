package ai

import (
	"tic-tac-toe/board" // Importing the board package for accessing GameState and its related functions
)

// FindWinningOrBlockingMove determines if there are any winning moves for the AI
// or blocking moves to prevent the opponent from winning.
func FindWinningOrBlockingMove(gameState *board.GameState, aiPlayer, opponentPlayer int) (int, int, bool) {
	// Iterate through every cell on the board
	for y := 0; y < gameState.Height; y++ {
		for x := 0; x < gameState.Width; x++ {
			// Check if the cell is empty
			if gameState.Board[y][x] == 0 {
				// Simulate AI's move and check for a win
				gameState.Board[y][x] = aiPlayer
				if board.GetWinner(gameState) == aiPlayer {
					gameState.Board[y][x] = 0 // Undo the simulated move
					return y, x, true         // Return the winning move
				}
				gameState.Board[y][x] = 0 // Undo the simulated move

				// Simulate opponent's move and check for a block
				gameState.Board[y][x] = opponentPlayer
				if board.GetWinner(gameState) == opponentPlayer {
					gameState.Board[y][x] = 0 // Undo the simulated move
					return y, x, true         // Return the blocking move
				}
				gameState.Board[y][x] = 0 // Undo the simulated move
			}
		}
	}

	// Return false if no winning or blocking moves are found
	return -1, -1, false
}
