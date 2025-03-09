package ai

import (
	"testing"
	"tic-tac-toe/board" // Import the board package for GameState and related functions
)

func TestFindWinningOrBlockingMove(t *testing.T) {
	// Test Case 1: AI has a winning move
	game := board.NewBoard(3, 3)
	game.Board[0][0], game.Board[0][1] = 2, 2 // AI player
	game.Board[0][2] = 0                      // Winning move for AI
	y, x, found := FindWinningOrBlockingMove(game, 2, 1)
	if !found || y != 0 || x != 2 {
		t.Errorf("Expected winning move at (0, 2), but got (%d, %d, %t)", y, x, found)
	}

	// Test Case 2: Opponent has a winning move, AI blocks it
	game = board.NewBoard(3, 3)
	game.Board[1][0], game.Board[1][1] = 1, 1 // Opponent player
	game.Board[1][2] = 0                      // Blocking move needed
	y, x, found = FindWinningOrBlockingMove(game, 2, 1)
	if !found || y != 1 || x != 2 {
		t.Errorf("Expected blocking move at (1, 2), but got (%d, %d, %t)", y, x, found)
	}

	// Test Case 3: No winning or blocking move
	game = board.NewBoard(3, 3)
	game.Board[0][0], game.Board[0][1], game.Board[0][2] = 1, 2, 1
	game.Board[1][0], game.Board[1][1], game.Board[1][2] = 2, 1, 2
	game.Board[2][0], game.Board[2][1], game.Board[2][2] = 2, 1, 2 // Full board, no winning moves
	y, x, found = FindWinningOrBlockingMove(game, 2, 1)
	if found {
		t.Errorf("Expected no moves, but got (%d, %d, %t)", y, x, found)
	}
}
