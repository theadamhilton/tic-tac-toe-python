package ai

import (
	"tic-tac-toe/board" // Importing the board package for accessing GameState and its related functions
)

func isBoardFull(game *board.GameState) bool {
	for _, row := range game.Board {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}

// Check if the player or AI wins, or if it's a tie
func evaluate(game *board.GameState, aiPlayer, humanPlayer int) int {
	if board.GetWinner(game) == aiPlayer {
		return +10
	} else if board.GetWinner(game) == humanPlayer {
		return -10
	}
	return 0
}

// Minimax algorithm to find the best move
func minimax(game *board.GameState, depth int, isMaximizing bool, aiPlayer, humanPlayer int) int {
	score := evaluate(game, aiPlayer, humanPlayer)

	// Base cases: return score if game is over
	if score == 10 || score == -10 || isBoardFull(game) {
		return score
	}

	if isMaximizing {
		best := -1000
		for row := 0; row < len(game.Board); row++ {
			for col := 0; col < len(game.Board[row]); col++ {
				if game.Board[row][col] == 0 {
					game.Board[row][col] = aiPlayer
					best = max(best, minimax(game, depth+1, false, aiPlayer, humanPlayer))
					game.Board[row][col] = 0
				}
			}
		}
		return best
	} else {
		best := 1000
		for row := 0; row < len(game.Board); row++ {
			for col := 0; col < len(game.Board[row]); col++ {
				if game.Board[row][col] == 0 {
					game.Board[row][col] = humanPlayer
					best = min(best, minimax(game, depth+1, true, aiPlayer, humanPlayer))
					game.Board[row][col] = 0
				}
			}
		}
		return best
	}
}

// Find the best move for the AI
func FindBestMove(game *board.GameState, aiPlayer, humanPlayer int) (int, int) {
	bestVal := -1000
	bestRow, bestCol := -1, -1

	// Iterate through all possible moves
	for row := 0; row < len(game.Board); row++ {
		for col := 0; col < len(game.Board[row]); col++ {
			if game.Board[row][col] == 0 { // Check if the cell is empty
				game.Board[row][col] = aiPlayer
				moveVal := minimax(game, 0, false, aiPlayer, humanPlayer)
				game.Board[row][col] = 0 // Undo the move

				if moveVal > bestVal {
					bestRow, bestCol = row, col
					bestVal = moveVal
				}
			}
		}
	}

	// Prioritize the center if no winning or blocking moves are found
	centerRow, centerCol := len(game.Board)/2, len(game.Board[0])/2
	if bestVal <= 0 && game.Board[centerRow][centerCol] == 0 { // Fallback logic
		return centerRow, centerCol
	}

	return bestRow, bestCol
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
